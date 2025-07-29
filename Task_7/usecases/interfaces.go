package usecases

type PasswordService interface {
	HashPassword(password string) (string, error)
	CheckPassword(password, hash string) bool
}

type JWTService interface {
	GenerateToken(userID , username, role string) (string, error)
	ValidateToken(token string) (map[string]interface{}, error)
}