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

func TestDeleteProductFromCart(t *testing.T) {
	type (
		args struct {
			payload *dto_service_cart.DeleteProductFromCartDTO
		}
		result struct {
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
			desc: "[ERROR]_because_product_not_found",
			input: args{
				payload: &dto_service_cart.DeleteProductFromCartDTO{
					ProductCode: "A321",
				},
			},
			output: result{
				err: errors.New("product not found"),
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
							ProductCode: "A321",
						}
					)
					var (
						result *model.Cart = nil
						err    error       = errors.New("product not found")
					)
					mockCartResource.On("GetCart", arg1, arg2).Return(result, err).Once()
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
			desc: "[ERROR]_because_invalid_db",
			input: args{
				payload: &dto_service_cart.DeleteProductFromCartDTO{
					ProductCode: "A123",
				},
			},
			output: result{
				err: errors.New("invalid db"),
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
							ProductCode: "A123",
						}
					)
					var (
						result *model.Cart = &model.Cart{
							ID:          mockUUID,
							ProductCode: "A123",
							ProductName: "Orange Fruit",
							Quantity:    4,
							CreatedAt:   mockDate,
							UpdatedAt:   mockDate,
						}
						err error = nil
					)
					mockCartResource.On("GetCart", arg1, arg2).Return(result, err).Once()
				}
				{
					var (
						arg1 *model.Cart = &model.Cart{
							ID:          mockUUID,
							ProductCode: "A123",
							ProductName: "Orange Fruit",
							Quantity:    4,
							CreatedAt:   mockDate,
							UpdatedAt:   mockDate,
						}
					)
					var (
						err error = errors.New("invalid db")
					)
					mockCartResource.On("DeleteCart", arg1).Return(err).Once()
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
			desc: "[SUCCESS]_success_delete_product_from_cart",
			input: args{
				payload: &dto_service_cart.DeleteProductFromCartDTO{
					ProductCode: "A123",
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
						arg1 *types.Query = nil
						arg2 *model.Cart  = &model.Cart{
							ProductCode: "A123",
						}
					)
					var (
						result *model.Cart = &model.Cart{
							ID:          mockUUID,
							ProductCode: "A123",
							ProductName: "Orange Fruit",
							Quantity:    4,
							CreatedAt:   mockDate,
							UpdatedAt:   mockDate,
						}
						err error = nil
					)
					mockCartResource.On("GetCart", arg1, arg2).Return(result, err).Once()
				}
				{
					var (
						arg1 *model.Cart = &model.Cart{
							ID:          mockUUID,
							ProductCode: "A123",
							ProductName: "Orange Fruit",
							Quantity:    4,
							CreatedAt:   mockDate,
							UpdatedAt:   mockDate,
						}
					)
					var (
						err error = nil
					)
					mockCartResource.On("DeleteCart", arg1).Return(err).Once()
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

			err := mockCartService.DeleteProductFromCart(tC.input.payload)

			assert.Equal(t, tC.output.err, err)

			tC.after()
		})
	}
}
