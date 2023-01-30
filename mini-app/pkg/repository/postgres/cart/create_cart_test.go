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
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	mockCartRepository ICartRepository
	mockQuery          sqlmock.Sqlmock
)

var (
	mockDate time.Time
	mockUUID string
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

func TestCreateCart(t *testing.T) {
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
			desc: "[ERROR] Because Missing Product Name",
			input: args{
				payload: &model.Cart{
					ProductCode: "A123",
					Quantity:    2,
				},
			},
			output: result{
				err: errors.New("missing product name"),
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
						arg3             = ""
						arg4             = 2
						arg5             = mockDate
						arg6             = mockDate
						arg7 interface{} = nil
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectExec(
						regexp.QuoteMeta(
							`INSERT INTO "carts" ("id","product_code","product_name","quantity","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6,$7)`,
						),
					).
						WithArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7).
						WillReturnError(errors.New("missing product name"))
					mockQuery.ExpectRollback()
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

			err := mockCartRepository.CreateCart(tC.input.payload)

			logrus.Info(err.Error())

			assert.Equal(t, tC.output.err, err)

			tC.after()
		})
	}
}
