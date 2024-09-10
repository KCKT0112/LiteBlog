package services

type UsersService interface {
}

type usersService struct{}

func NewUsersService() UsersService {
	return &usersService{}
}
