package active_user_mapper_fixture

import (
	"getfund-api-v2/internal/domain/auth/core/domain_service/activate_user_mapper"
	"getfund-api-v2/internal/domain/auth/core/entity"
)

func NewSut() activate_user_mapper.Mapper {
	return activate_user_mapper.New()
}

func GetUserEntity() *entity.User {
	return entity.NewUser(
		"fake-first-name",
		"fake-last-name",
		"fake-username",
		"fakeStrongPassword99",
	)
}
