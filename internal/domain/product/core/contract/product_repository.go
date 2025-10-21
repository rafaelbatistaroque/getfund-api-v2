package contract

import "getfund-api-v2/internal/domain/product/core/dto"

type Repository interface {
	GetProductById(productId int) (*dto.ProductDto, error)
}
