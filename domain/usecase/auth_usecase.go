package usecase

import "BlogApp/domain/model"

type AuthUseCase interface {
	Register( currUser *model.AuthenticatedUser, userCreate *model.UserCreate, param any) (*model.UserInfo, string, error)
	Login( currUser *model.AuthenticatedUser, userLogin *model.UserLogin, parma any) (*model.Token, string, error)
	AdminRegister( currUser *model.AuthenticatedUser, userCreate *model.UserCreate, param any) (*model.UserInfo, string, error)
}