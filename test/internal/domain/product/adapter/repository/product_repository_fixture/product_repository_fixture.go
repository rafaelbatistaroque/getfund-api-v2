package product_repository_fixture

import (
	product_repository "getfund-api-v2/internal/domain/product/adapter/repository"
	product_contract "getfund-api-v2/internal/domain/product/core/contract"
	"getfund-api-v2/internal/infra/db"
	"getfund-api-v2/internal/infra/db/schema"
	"getfund-api-v2/test/helper/db_fixture"
)

func NewSUT() (product_contract.Repository, *db.GetFund) {
	get_fund_db := db_fixture.NewMemoryDB(
		&schema.Product{},
	)

	return product_repository.New(get_fund_db), get_fund_db
}
