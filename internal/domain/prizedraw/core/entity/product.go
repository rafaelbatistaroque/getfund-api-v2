package entity

import "errors"

type Product struct {
	id       int
	isActive bool
}

const (
	_COUPON_INVALID_FOR_PRODUCT = "coupon is not valid for this product"
	_INACTIVE_PRODUCT           = "inactive product"
)

func FillProduct(id int, isActive bool) *Product {
	return &Product{
		id:       id,
		isActive: isActive,
	}
}
func (p *Product) GetId() int { return p.id }

func (p *Product) Validate(selectedProductId int) error {
	if !p.isActive {
		return errors.New(_INACTIVE_PRODUCT)
	}

	if p.id != selectedProductId {
		return errors.New(_COUPON_INVALID_FOR_PRODUCT)
	}

	return nil
}
