package user_service

import user_repository "github.com/SawitProRecruitment/UserService/repository/user"

type UserService interface {
	Register(request RegisterUserRequest) (*user_repository.User, []error)
}
