package user_service

import user_repository "github.com/SawitProRecruitment/UserService/repository/user"

func (s *UserServiceImpl) Register(request RegisterUserRequest) (*user_repository.User, []error) {
	var err []error

	err = append(err, s.NameValidator.Validate("Full Name", request.FullName)[:]...)
	err = append(err, s.EmailValidator.Validate("Email", request.Email)[:]...)
	err = append(err, s.PhoneValidator.Validate("Phone", request.Phone)[:]...)
	err = append(err, s.PasswordValidator.Validate("Password", request.Password)[:]...)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *UserServiceImpl) saveToDB(request RegisterUserRequest) (*user_repository.User, error) {
	var user user_repository.User

	// s.UserRepository.Db

	return &user, nil
}
