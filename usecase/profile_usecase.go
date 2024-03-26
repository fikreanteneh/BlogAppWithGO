package usecase

import (
	"BlogApp/config"
	"BlogApp/domain"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
	"context"
)

type ProfileUseCase struct {
	context        context.Context
	environment    config.Environment
	UserRepository domain.UserRepository
}

// DeleteProfile implements usecase.ProfileUseCase.
func (p *ProfileUseCase) DeleteProfile(currUser *model.AuthenticatedUser) (*model.UserInfo, error) {
	panic("unimplemented")
}

// GetProfile implements usecase.ProfileUseCase.
func (p *ProfileUseCase) GetProfile(currUser *model.AuthenticatedUser) (*model.UserInfo, error) {
	panic("unimplemented")
}

// UpdateEmail implements usecase.ProfileUseCase.
func (p *ProfileUseCase) UpdateEmail(updated *model.UserUpdateEmail, currUser *model.AuthenticatedUser) (*model.UserInfo, error) {
	panic("unimplemented")
}

// UpdatePassword implements usecase.ProfileUseCase.
func (p *ProfileUseCase) UpdatePassword(updated *model.UserUpdatePassword, currUser *model.AuthenticatedUser) (*model.UserInfo, error) {
	panic("unimplemented")
}

// UpdateProfile implements usecase.ProfileUseCase.
func (p *ProfileUseCase) UpdateProfile(updated *model.UserUpdateProfile, currUser *model.AuthenticatedUser) (*model.UserInfo, error) {
	panic("unimplemented")
}

// UpdateUsername implements usecase.ProfileUseCase.
func (p *ProfileUseCase) UpdateUsername(updated *model.UserUpdateUsername, currUser *model.AuthenticatedUser) (*model.UserInfo, error) {
	panic("unimplemented")
}

func NewProfileUseCase(environment *config.Environment, userRepository *domain.UserRepository) usecase.ProfileUseCase {
	return &ProfileUseCase{
		environment:    *environment,
		UserRepository: *userRepository,
	}
}
