package request_controller_cart

type AddProductToCartRequestBody struct {
	ProductCode string `json:"product_code" validate:"required,uppercase,alphanum"`
	ProductName string `json:"product_name" validate:"required"`
	Quantity    int64  `json:"quantity" validate:"required,min=1"`
}

func (r *AddProductToCartRequestBody) CustomValidate() error {
	return nil
}
