package product_contract

import "getfund-api-v2/internal/domain/product/core/dto/product_dto"

type Repository interface {
	GetProductById(productId int) (*product_dto.ProductDto, error)
}
