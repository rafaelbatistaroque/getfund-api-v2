package value_object

type UserCouponApply struct {
	userId int
}

func NewUserCouponApply(userId int) UserCouponApply {
	return UserCouponApply{
		userId: userId,
	}
}

func (u UserCouponApply) GetUserId() int {
	return u.userId
}

func (u UserCouponApply) IsEqual(userId int) bool {
	return u.userId == userId
}
