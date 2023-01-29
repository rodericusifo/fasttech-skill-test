package request_controller_cart

type DeleteProductFromCartRequestBody struct {
	ProductCode string `json:"product_code" validate:"required,uppercase,alphanum"`
}

func (r *DeleteProductFromCartRequestBody) CustomValidate() error {
	return nil
}
