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
func (p *ProfileUseCase) DeleteProfile(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*model.UserInfo, string, error) {
	panic("unimplemented")
}

// GetProfile implements usecase.ProfileUseCase.
func (p *ProfileUseCase) GetProfile(currUser *model.AuthenticatedUser, dto any, param any) (*model.UserInfo, string, error) {
	panic("unimplemented")
}

// UpdateEmail implements usecase.ProfileUseCase.
func (p *ProfileUseCase) UpdateEmail(currUser *model.AuthenticatedUser, updated *model.UserUpdateEmail, param any) (*model.UserInfo, string, error) {
	panic("unimplemented")
}

// UpdatePassword implements usecase.ProfileUseCase.
func (p *ProfileUseCase) UpdatePassword(currUser *model.AuthenticatedUser, updated *model.UserUpdatePassword, param any) (*model.UserInfo, string, error) {
	panic("unimplemented")
}

// UpdateProfile implements usecase.ProfileUseCase.
func (p *ProfileUseCase) UpdateProfile(currUser *model.AuthenticatedUser, updated *model.UserUpdateProfile, param any) (*model.UserInfo, string, error) {
	panic("unimplemented")
}

// UpdateUsername implements usecase.ProfileUseCase.
func (p *ProfileUseCase) UpdateUsername(currUser *model.AuthenticatedUser, updated *model.UserUpdateUsername, param any) (*model.UserInfo, string, error) {
	panic("unimplemented")
}

func NewProfileUseCase(context *context.Context, environment *config.Environment, userRepository *domain.UserRepository) usecase.ProfileUseCase {
	return &ProfileUseCase{
		context:        *context,
		environment:    *environment,
		UserRepository: *userRepository,
	}
}
