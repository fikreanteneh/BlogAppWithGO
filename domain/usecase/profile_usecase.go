package usecase

import "BlogApp/domain/model"

type ProfileUseCase interface {
	GetProfile(
		currUser *model.AuthenticatedUser, 
		dto any, 
		param any) (*model.UserInfo, string, error)
	UpdateProfile(
		currUser *model.AuthenticatedUser, 
		updated *model.UserUpdateProfile, 
		param any) (*model.UserInfo, string, error)
	DeleteProfile(
		currUser *model.AuthenticatedUser, 
		dto any, 
		param *model.IdParam) (*model.UserInfo, string, error)
	UpdateUsername(
		currUser *model.AuthenticatedUser, 
		updated *model.UserUpdateUsername,
		param any) (*model.UserInfo, string, error)
	UpdatePassword(
		currUser *model.AuthenticatedUser, 
		updated *model.UserUpdatePassword,
		param any) (*model.UserInfo, string, error)
	UpdateEmail(
		currUser *model.AuthenticatedUser, 
		updated *model.UserUpdateEmail,
		param any) (*model.UserInfo, string, error)
	//TODO: add the profile picture update method
}