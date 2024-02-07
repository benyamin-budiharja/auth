package user_service

import (
	user_repository "github.com/SawitProRecruitment/UserService/repository/user"
	"github.com/SawitProRecruitment/UserService/service/validation"
)

type UserServiceImpl struct {
	UserRepository    user_repository.UserRepository
	NameValidator     validation.Validator
	EmailValidator    validation.Validator
	PhoneValidator    validation.Validator
	PasswordValidator validation.Validator
}

func NewUserService(repo user_repository.UserRepository) UserServiceImpl {
	return UserServiceImpl{
		UserRepository: repo,
	}
}
