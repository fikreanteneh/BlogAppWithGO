package usecase

import "BlogApp/domain/model"

type AuthUseCase interface {
	Register(userCreate *model.UserCreate) (*model.UserInfo, error)
	Login(userLogin *model.UserLogin) (string, error)
	AdminRegister(userCreate *model.UserCreate, currUser *model.AuthenticatedUser) (*model.UserInfo, error)
}