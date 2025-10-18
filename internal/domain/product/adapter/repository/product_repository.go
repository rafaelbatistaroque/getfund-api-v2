package product_repository

import (
	"errors"
	product_contract "getfund-api-v2/internal/domain/product/core/contract"
	"getfund-api-v2/internal/domain/product/core/dto/product_dto"
	"getfund-api-v2/internal/infra/db"
	"getfund-api-v2/internal/infra/db/schema"

	"gorm.io/gorm"
)

type productRepository struct {
	db *db.GetFund
}

func New(db *db.GetFund) product_contract.Repository {
	return &productRepository{db: db}
}

func (p *productRepository) GetProductById(productId int) (*product_dto.ProductDto, error) {
	var product = &schema.Product{}

	result := p.db.
		Select("id, is_active").
		Where("id=?", productId).
		First(product)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("product not found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return product_dto.MapFromSchema(product), nil
}
