package usecase

import "BlogApp/domain/model"

type ProfileUseCase interface {
	GetProfile(currUser *model.AuthenticatedUser) (*model.UserInfo, error)
	UpdateProfile(updated *model.UserUpdateProfile, currUser *model.AuthenticatedUser) (*model.UserInfo, error)
	DeleteProfile(currUser *model.AuthenticatedUser) (*model.UserInfo, error)
	UpdateUsername(updated *model.UserUpdateUsername, currUser *model.AuthenticatedUser) (*model.UserInfo, error)
	UpdatePassword(updated *model.UserUpdatePassword, currUser *model.AuthenticatedUser) (*model.UserInfo, error)
	UpdateEmail(updated *model.UserUpdateEmail, currUser *model.AuthenticatedUser) (*model.UserInfo, error)
	//TODO: add the profile picture update method
}