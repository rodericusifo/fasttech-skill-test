package cart_postgres_repository

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/libs/util"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/model"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/shared/constant"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/shared/helper"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
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

func TestUpdateCart(t *testing.T) {
	type (
		args struct {
			payload *model.Cart
		}
		result struct {
			err error
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
			desc: "[ERROR] Because Something Error Happens",
			input: args{
				payload: &model.Cart{
					ID:          mockUUID,
					ProductCode: "A123",
					ProductName: "Orange Fruit",
					Quantity:    4,
					CreatedAt:   mockDate,
					UpdatedAt:   mockDate,
					DeletedAt:   gorm.DeletedAt{},
				},
			},
			output: result{
				err: errors.New("error something"),
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
						arg1             = mockUUID
						arg2             = "A123"
						arg3             = "Orange Fruit"
						arg4             = 4
						arg5             = mockDate
						arg6             = mockDate.Local()
						arg7 interface{} = nil
						arg8             = mockUUID
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectExec(
						regexp.QuoteMeta(
							`
								UPDATE 
									"carts" 
								SET 
									"id"=$1,"product_code"=$2,"product_name"=$3,"quantity"=$4,"created_at"=$5,"updated_at"=$6,"deleted_at"=$7 
								WHERE 
									id = $8 
								AND "carts"."deleted_at" IS NULL
							`,
						),
					).
						WithArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8).
						WillReturnError(errors.New("error something"))
					mockQuery.ExpectRollback()
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
			desc: "[SUCCESS] Success Update Cart",
			input: args{
				payload: &model.Cart{
					ID:          mockUUID,
					ProductCode: "A123",
					ProductName: "Orange Fruit",
					Quantity:    4,
					CreatedAt:   mockDate,
					UpdatedAt:   mockDate,
					DeletedAt:   gorm.DeletedAt{},
				},
			},
			output: result{
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
						arg1             = mockUUID
						arg2             = "A123"
						arg3             = "Orange Fruit"
						arg4             = 4
						arg5             = mockDate
						arg6             = mockDate.Local()
						arg7 interface{} = nil
						arg8             = mockUUID
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectExec(
						regexp.QuoteMeta(
							`
								UPDATE 
									"carts" 
								SET 
									"id"=$1,"product_code"=$2,"product_name"=$3,"quantity"=$4,"created_at"=$5,"updated_at"=$6,"deleted_at"=$7 
								WHERE 
									id = $8 
								AND "carts"."deleted_at" IS NULL
							`,
						),
					).
						WithArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8).
						WillReturnResult(sqlmock.NewResult(0, 1))
					mockQuery.ExpectCommit()
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

			err := mockCartRepository.UpdateCart(tC.input.payload)

			assert.Equal(t, tC.output.err, err)

			tC.after()
		})
	}
}
