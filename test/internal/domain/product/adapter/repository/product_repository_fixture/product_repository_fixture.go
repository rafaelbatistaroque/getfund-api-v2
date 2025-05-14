package product_repository_fixture

import (
	product_repository "getfund-api-v2/internal/domain/product/adapter/repository"
	product_contract "getfund-api-v2/internal/domain/product/core/contract"
	"getfund-api-v2/pkg/db/schema"
	"getfund-api-v2/test/helper/db_fixture"

	"gorm.io/gorm"
)

func NewSUT() (product_contract.Repository, *gorm.DB) {
	db := db_fixture.NewMemoryDB(
		&schema.Product{},
	)

	return product_repository.New(db), db
}
