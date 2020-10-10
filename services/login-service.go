package services

type LoginService interface {
	Login(username, password string) bool
}

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

func NewLoginService() LoginService {
	return &loginService{
		authorizedUsername: "pragmatic",
		authorizedPassword: "reviews",
	}
}

func (service *loginService) Login(username, password string) bool {
	return service.authorizedUsername == username && service.authorizedPassword == password
}
