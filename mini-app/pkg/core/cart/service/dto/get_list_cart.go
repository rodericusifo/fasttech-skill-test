package dto_service_cart

type GetListCartDTO struct {
	ProductName string
	ProductCode string
	Quantity    int64
}

type GetListCartPayloadDTO struct {
	ProductName string
	Quantity    int64
}
