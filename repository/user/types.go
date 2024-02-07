// This file contains types that are used in the repository layer.
package user_repository

type GetUserByIdInput struct {
	Id int64
}

type User struct {
	Id       int64
	FullName string
	Phone    string
	Email    string
	Password string
	Salt     string
}
