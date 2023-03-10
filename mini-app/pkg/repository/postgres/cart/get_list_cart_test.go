package cart_postgres_repository

import (
	"errors"
	"fmt"
	"regexp"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/libs/types"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/libs/util"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/model"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/shared/constant"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/shared/helper"
	"github.com/stretchr/testify/assert"
)

func init() {
	db, mock := helper.MockConnectionDatabase(constant.POSTGRES)

	mockCartRepository = InitCartRepository(db)
	mockQuery = mock

	layoutFormat := "2006-01-02 15:04:05"
	value := "2015-09-02 08:04:00"
	mockDate, _ = time.Parse(layoutFormat, value)

	mockUUID = "ac0d6ce3-ff02-4024-896b-ea0ceba32182"
}

func TestGetListCart(t *testing.T) {
	type (
		args struct {
			query   *types.Query
			payload *model.Cart
		}
		result struct {
			value []*model.Cart
			err   error
		}
	)

	testCases := []struct {
		desc   string
		input  args
		output result
		before func()
		after  func()
	}{
		{
			desc: "[ERROR]_because_something_error_happens",
			input: args{
				query: &types.Query{
					SelectColumns: []string{
						"id",
						"product_code",
						"product_name",
						"quantity",
					},
				},
				payload: &model.Cart{
					ProductName: "Orange Fruit",
					Quantity:    4,
				},
			},
			output: result{
				value: nil,
				err:   errors.New("something error"),
			},
			before: func() {
				{
					monkey.Patch(util.GenerateUUID, func() string {
						return mockUUID
					})
				}
				{
					monkey.Patch(time.Now, func() time.Time {
						return mockDate
					})
				}
				{
					var (
						arg1 = fmt.Sprintf("%%%s%%", "Orange Fruit")
						arg2 = 4
					)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							`
								SELECT
									"id","product_code","product_name","quantity","created_at","updated_at"
								FROM
									"carts"
								WHERE
									product_name LIKE $1
								AND
									quantity = $2
								AND
									"carts"."deleted_at" IS NULL
							`,
						),
					).
						WithArgs(arg1, arg2).
						WillReturnError(errors.New("something error"))
				}
			},
			after: func() {
				{
					monkey.Unpatch(util.GenerateUUID)
				}
				{
					monkey.Unpatch(time.Now)
				}
			},
		},
		{
			desc: "[SUCCESS]_success_get_list_cart",
			input: args{
				query: &types.Query{
					SelectColumns: []string{
						"id",
						"product_code",
						"product_name",
						"quantity",
					},
				},
				payload: &model.Cart{
					ProductName: "Orange Fruit",
					Quantity:    4,
				},
			},
			output: result{
				value: []*model.Cart{
					{
						ID:          mockUUID,
						ProductCode: "A123",
						ProductName: "Orange Fruit",
						Quantity:    4,
						CreatedAt:   mockDate,
						UpdatedAt:   mockDate,
					},
				},
				err: nil,
			},
			before: func() {
				{
					monkey.Patch(util.GenerateUUID, func() string {
						return mockUUID
					})
				}
				{
					monkey.Patch(time.Now, func() time.Time {
						return mockDate
					})
				}
				{
					var (
						arg1 = fmt.Sprintf("%%%s%%", "Orange Fruit")
						arg2 = 4

						rowsInstance = sqlmock.NewRows([]string{"id", "product_code", "product_name", "quantity", "created_at", "updated_at"})
					)

					rowsInstance.AddRow(mockUUID, "A123", "Orange Fruit", 4, mockDate, mockDate)

					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							`
								SELECT
									"id","product_code","product_name","quantity","created_at","updated_at"
								FROM
									"carts"
								WHERE
									product_name LIKE $1
								AND
									quantity = $2
								AND
									"carts"."deleted_at" IS NULL
							`,
						),
					).
						WithArgs(arg1, arg2).
						WillReturnRows(rowsInstance)
				}
			},
			after: func() {
				{
					monkey.Unpatch(util.GenerateUUID)
				}
				{
					monkey.Unpatch(time.Now)
				}
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			result, err := mockCartRepository.GetListCart(tC.input.query, tC.input.payload)

			assert.Equal(t, tC.output.err, err)
			assert.Equal(t, tC.output.value, result)

			tC.after()
		})
	}
}
