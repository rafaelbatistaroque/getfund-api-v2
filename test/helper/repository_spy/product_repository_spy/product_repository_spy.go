package product_repository_spy

import (
	"errors"
	"getfund-api-v2/internal/domain/product/core/dto"
)

type ProductRepositorySpy struct {
	Params        map[string]any
	CallsCount    map[string]int
	ErrorResult   map[string]error
	SuccessResult map[string]any
}

func New() *ProductRepositorySpy {

	return &ProductRepositorySpy{
		Params:        make(map[string]any, 1),
		ErrorResult:   make(map[string]error),
		SuccessResult: make(map[string]any, 1),
		CallsCount:    make(map[string]int, 1)}
}

func (r *ProductRepositorySpy) GetProductById(productId int) (*dto.ProductDto, error) {
	r.Params["GetProductById:productId"] = productId

	r.CallsCount["GetProductById"]++

	sucess := r.SuccessResult["GetProductById"]
	if sucess != nil {
		return sucess.(*dto.ProductDto), nil
	}

	return nil, r.ErrorResult["GetProductById"]
}

func (r *ProductRepositorySpy) DefineGetProductByIdError() {
	r.ErrorResult["GetProductById"] = errors.New("fake-error")
}

func (r *ProductRepositorySpy) DefineGetProductByIdSuccess(product *dto.ProductDto) {
	if product != nil {
		r.SuccessResult["GetProductById"] = product
		return
	}

	r.SuccessResult["GetProductById"] = &dto.ProductDto{}
}
