package prizedraw_dto

type ValidationData struct {
	Product *ProductData `json:"product"`
}

type ProductData struct {
	Id               int  `json:"id"`
	IsActive         bool `json:"is_active"`
	EntranceQuantity int  `json:"entrance_qty"`
}
