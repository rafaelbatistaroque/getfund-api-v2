package active_user_mapper_fixture

import (
	"getfund-api-v2/internal/domain/user/core/domain_service/activate_user_mapper"
	"getfund-api-v2/internal/domain/user/core/entity/user_entity"
)

func NewSut() activate_user_mapper.Mapper {
	return activate_user_mapper.New()
}

func GetUserEntity() *user_entity.User {
	return user_entity.New(
		"fake-first-name",
		"fake-last-name",
		"fake-email",
		"m",
		"fakeStrongPassword99",
		"@FakeSocial",
		"https://social.com",
		1,
		1)
}
