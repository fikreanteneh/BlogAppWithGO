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
	user, err := p.UserRepository.GetById(p.context, currUser.UserID)
	if err != nil {	
		return nil, "", err
	}
    deleted, err := p.UserRepository.Delete(p.context, user)
    if err != nil {
        return nil, "", err
    }
    return &model.UserInfo{
		Username: deleted.Username,
		Name: deleted.Name,
		Bio: deleted.Bio}, "User deleted successfully", nil
}

// GetProfile implements usecase.ProfileUseCase.
func (p *ProfileUseCase) GetProfile(currUser *model.AuthenticatedUser, dto any, param any) (*domain.User, string, error) {
    user, err := p.UserRepository.GetById(p.context, currUser.UserID)
    if err != nil {
        return nil, "", err
    }
    return &domain.User{
		Username: user.Username,
		Name: user.Name,
		Bio: user.Bio,
		Email: user.Email,
		Role: user.Role,
		UserID: user.UserID,
		CreatedAt: user.CreatedAt,

	}, "User fetched successfully", nil
}

// UpdateEmail implements usecase.ProfileUseCase.
func (p *ProfileUseCase) UpdateEmail(currUser *model.AuthenticatedUser, updated *model.UserUpdateEmail, param any) (*model.UserInfo, string, error) {
	//TODO: Validation
    user, err := p.UserRepository.GetById(p.context, currUser.UserID)
	if err != nil {
		return nil, "", err
	}
	user.Email = updated.Email
	updatedUser, err := p.UserRepository.UpdateEmail(p.context, user)
	if err != nil {
		return nil, "", err
	}
	return &model.UserInfo{
		Username: updatedUser.Username,
		Name: updatedUser.Name,
		Bio: updatedUser.Bio,
	}, "Email updated successfully", nil
}

// UpdatePassword implements usecase.ProfileUseCase.
func (p *ProfileUseCase) UpdatePassword(currUser *model.AuthenticatedUser, updated *model.UserUpdatePassword, param any) (*model.UserInfo, string, error) {
	//TODO: Validation
	user, err := p.UserRepository.GetById(p.context, currUser.UserID)
	if err != nil {
		return nil, "", err
	}
	user.Password = updated.Password
	updatedUser, err := p.UserRepository.UpdatePassword(p.context, user)
	if err != nil {
		return nil, "", err
	}
	return &model.UserInfo{
		Username: updatedUser.Username,
		Name: updatedUser.Name,
		Bio: updatedUser.Bio,
	}, "Password updated successfully", nil
}

// UpdateProfile implements usecase.ProfileUseCase.
func (p *ProfileUseCase) UpdateProfile(currUser *model.AuthenticatedUser, updated *model.UserUpdateProfile, param any) (*model.UserInfo, string, error) {
    	//TODO: Validation
	user, err := p.UserRepository.GetById(p.context, currUser.UserID)
	if err != nil {
		return nil, "", err
	}
	updated.Bio = updated.Bio
	updated.Name = updated.Name
	updatedUser, err := p.UserRepository.UpdateProfile(p.context, user)
	if err != nil {
		return nil, "", err
	}
	return &model.UserInfo{
		Username: updatedUser.Username,
		Name: updatedUser.Name,
		Bio: updatedUser.Bio,
	}, "Password updated successfully", nil
}

// UpdateUsername implements usecase.ProfileUseCase.
func (p *ProfileUseCase) UpdateUsername(currUser *model.AuthenticatedUser, updated *model.UserUpdateUsername, param any) (*model.UserInfo, string, error) {
    	//TODO: Validation
	user, err := p.UserRepository.GetById(p.context, currUser.UserID)
	if err != nil {
		return nil, "", err
	}
	user.Username = updated.Username
	updatedUser, err := p.UserRepository.UpdateUsername(p.context, user)
	if err != nil {
		return nil, "", err
	}
	return &model.UserInfo{
		Username: updatedUser.Username,
		Name: updatedUser.Name,
		Bio: updatedUser.Bio,
	}, "Password updated successfully", nil

}

func NewProfileUseCase(context *context.Context, environment *config.Environment, userRepository *domain.UserRepository) usecase.ProfileUseCase {
	return &ProfileUseCase{
		context:        *context,
		environment:    *environment,
		UserRepository: *userRepository,
	}
}
