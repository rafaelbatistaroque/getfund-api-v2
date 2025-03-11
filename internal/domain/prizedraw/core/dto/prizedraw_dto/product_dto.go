package prizedraw_dto

import "getfund-api-v2/internal/domain/prizedraw/core/entity"

type ProductDto struct {
	Id       int  `json:"id"`
	IsActive bool `json:"is_active"`
}

func (c ProductDto) ToEntity() *entity.Product {

	return entity.FillProduct(
		c.Id,
		c.IsActive,
	)
}
