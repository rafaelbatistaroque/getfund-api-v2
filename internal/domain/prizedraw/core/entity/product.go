package entity

type Product struct {
	id               int
	isActive         bool
	entranceQuantity int
}

func ProductFill(id int, isActive bool, entranceQuantity int) *Product {
	return &Product{
		id:               id,
		isActive:         isActive,
		entranceQuantity: entranceQuantity,
	}
}
func (p *Product) GetId() int { return p.id }

func (p *Product) GetEntranceQuantity() int { return p.entranceQuantity }

func (p *Product) IsActive() bool {
	return p.isActive
}
