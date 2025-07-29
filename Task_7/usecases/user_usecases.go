package usecases

import (
	"errors"
	domain "task-manager/Domain"
)

type UserUsecase struct {
	userRepo 		domain.UserRepository
	PasswordService PasswordService
	JWTService		JWTService
}

func NewUserUsecase(ur domain.UserRepository, ps PasswordService, js JWTService) *UserUsecase {
	return &UserUsecase{
		userRepo: 		 ur,
		PasswordService: ps,
		JWTService: 	 js,
	}
}

func (u *UserUsecase) Register(user *domain.User) error {
	existing, _ := u.userRepo.FetchByUsername(user.Username)
	if existing != nil {
		return errors.New("usrename already taken")
	}
	hashed, err := u.PasswordService.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed
	if user.Role == "" {
		user.Role = domain.Roleuser
	}
	return u.userRepo.Create(user)

}

func (u *UserUsecase) Login(username, password string) (*domain.User, string, error) {
	user, err := u.userRepo.FetchByUsername(username)
	if err != nil || !u.PasswordService.CheckPassword(user.Password, password) {
		return 	nil, "", errors.New("username already taken")
	}

	token, err := u.JWTService.GenerateToken(user.ID, user.Username, string(user.Role))
	if err != nil {
		return nil, "", err 
	}
	return user, token, nil 
}

