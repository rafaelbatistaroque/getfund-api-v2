package entity

type Product struct {
	id       int
	isActive bool
}

func ProductFill(id int, isActive bool) *Product {
	return &Product{
		id:       id,
		isActive: isActive,
	}
}
func (p *Product) GetId() int { return p.id }

func (p *Product) IsActive() bool { return p.isActive }
