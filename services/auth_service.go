package services

type AuthService struct {}

func NewAuthService() *AuthService {
    return &AuthService{}
}

func (s *AuthService) Authenticate(email, password string) bool {
    return false
}
