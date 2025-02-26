package value_object

type CouponType struct {
	id          uint
	code        uint
	description string
}

func NewCouponType(id uint, code uint, description string) CouponType {
	return CouponType{
		id:          id,
		code:        code,
		description: description,
	}
}
func (c CouponType) GetId() int {
	return int(c.id)
}

func (c CouponType) GetCode() int {
	return int(c.code)
}

func (c CouponType) GetDescription() string {
	return c.description
}
