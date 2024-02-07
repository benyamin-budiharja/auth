package user_service

type RegisterUserRequest struct {
	FullName string
	Email    string
	Phone    string
	Password string
}
