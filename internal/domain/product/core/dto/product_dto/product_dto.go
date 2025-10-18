package product_dto

import "getfund-api-v2/internal/infra/db/schema"

type ProductDto struct {
	Id       int  `json:"id"`
	IsActive bool `json:"is_active"`
}

func MapFromSchema(schema *schema.Product) *ProductDto {
	dto := new(ProductDto)
	dto.Id = int(schema.ID)
	dto.IsActive = schema.IsActive

	return dto
}
