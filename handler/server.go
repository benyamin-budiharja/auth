package handler

import user_repository "github.com/SawitProRecruitment/UserService/repository/user"

type Server struct {
	Repository user_repository.UserRepositoryInterface
}

type NewServerOptions struct {
	Repository user_repository.UserRepositoryInterface
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{}
}
