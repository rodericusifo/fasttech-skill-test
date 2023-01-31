package service_cart

import (
	"errors"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/libs/types"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/libs/util"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/mocks"
	dto_service_cart "github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/core/cart/service/dto"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/pkg/model"
	"github.com/stretchr/testify/assert"
)

func init() {
	mockCartResource = new(mocks.ICartResource)

	mockCartService = InitCartService(mockCartResource)

	layoutFormat := "2006-01-02 15:04:05"
	value := "2015-09-02 08:04:00"
	mockDate, _ = time.Parse(layoutFormat, value)

	mockUUID = "ac0d6ce3-ff02-4024-896b-ea0ceba32182"
}

func TestGetListCart(t *testing.T) {
	type (
		args struct {
			payload *dto_service_cart.GetListCartPayloadDTO
		}
		result struct {
			value []*dto_service_cart.GetListCartDTO
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
				payload: &dto_service_cart.GetListCartPayloadDTO{
					ProductName: "Orange Fruit",
					Quantity:    4,
				},
			},
			output: result{
				value: nil,
				err:   errors.New("error something"),
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
						arg1 *types.Query = nil
						arg2 *model.Cart  = &model.Cart{
							ProductName: "Orange Fruit",
							Quantity:    4,
						}
					)
					var (
						result []*model.Cart = nil
						err    error         = errors.New("error something")
					)
					mockCartResource.On("GetListCart", arg1, arg2).Return(result, err).Once()
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
				payload: &dto_service_cart.GetListCartPayloadDTO{
					ProductName: "Orange Fruit",
					Quantity:    4,
				},
			},
			output: result{
				value: []*dto_service_cart.GetListCartDTO{
					{
						ProductName: "Orange Fruit",
						ProductCode: "A123",
						Quantity:    4,
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
						arg1 *types.Query = nil
						arg2 *model.Cart  = &model.Cart{
							ProductName: "Orange Fruit",
							Quantity:    4,
						}
					)
					var (
						result []*model.Cart = []*model.Cart{
							{
								ID:          mockUUID,
								ProductCode: "A123",
								ProductName: "Orange Fruit",
								Quantity:    4,
								CreatedAt:   mockDate,
								UpdatedAt:   mockDate,
							},
						}
						err error = nil
					)
					mockCartResource.On("GetListCart", arg1, arg2).Return(result, err).Once()
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

			result, err := mockCartService.GetListCart(tC.input.payload)

			assert.Equal(t, tC.output.err, err)
			assert.Equal(t, tC.output.value, result)

			tC.after()
		})
	}
}
