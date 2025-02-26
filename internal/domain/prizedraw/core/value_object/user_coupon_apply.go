package value_object

type CouponUserApply struct {
	userId int
}

func NewUserCouponApply(userId int) CouponUserApply {
	return CouponUserApply{
		userId: userId,
	}
}

func (u CouponUserApply) GetUserId() int {
	return u.userId
}

func (u CouponUserApply) IsEqual(userId int) bool {
	return u.userId == userId
}
