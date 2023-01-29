package request_controller_cart

type GetListCartRequestQuery struct {
	ProductName string `query:"product_name" validate:"omitempty"`
	Quantity    int64  `query:"quantity" validate:"omitempty,min=1"`
}

func (r *GetListCartRequestQuery) CustomValidate() error {
	return nil
}
