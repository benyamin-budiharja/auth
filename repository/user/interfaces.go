// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package user_repository

import "context"

type UserRepositoryInterface interface {
	GetUserById(ctx context.Context, input GetUserByIdInput) (output *User, err error)
}
