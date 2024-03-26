package usecase

import "BlogApp/domain/model"

type AuthUseCase interface {
	Register(userCreate *model.UserCreate) (*model.UserList, error)
	Login(userLogin *model.UserLogin) (string, error)
	AdminRegister(userCreate *model.UserCreate, currUser *model.AuthenticatedUser) (*model.UserList, error)
}