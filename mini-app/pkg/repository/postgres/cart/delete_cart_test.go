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

func TestDeleteCart(t *testing.T) {
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
					UpdatedAt:   mockDate.Local(),
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
						arg1 = mockUUID
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectExec(
						regexp.QuoteMeta(
							`
								DELETE FROM 
									"carts" 
								WHERE 
									"carts"."id" = $1
							`,
						),
					).
						WithArgs(arg1).
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
			desc: "[SUCCESS] Success Delete Cart",
			input: args{
				payload: &model.Cart{
					ID:          mockUUID,
					ProductCode: "A123",
					ProductName: "Orange Fruit",
					Quantity:    4,
					CreatedAt:   mockDate,
					UpdatedAt:   mockDate.Local(),
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
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectExec(
						regexp.QuoteMeta(
							`
								DELETE FROM 
									"carts" 
								WHERE 
									"carts"."id" = $1
							`,
						),
					).
						WithArgs(arg1).
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

			err := mockCartRepository.DeleteCart(tC.input.payload)

			assert.Equal(t, tC.output.err, err)

			tC.after()
		})
	}
}
