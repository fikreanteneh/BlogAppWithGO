package usecase

import (
	"BlogApp/config"
	"BlogApp/domain"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
	"context"
	"errors"
)

type ProfileUseCase struct {
	context        context.Context
	environment    config.Environment
	UserRepository domain.UserRepository
}

func NewProfileUseCase(environment *config.Environment, userRepository *domain.UserRepository) usecase.ProfileUseCase {
	return &ProfileUseCase{
		environment:    *environment,
		UserRepository: *userRepository,
	}
}

// DeleteProfile implements usecase.ProfileUseCase.
func (p *ProfileUseCase) DeleteProfile(currUser *model.AuthenticatedUser) (*model.UserInfo, error) {
	//TODO: Authorization Handling

	//     // Authorization Handling ... It might be done in this way later by defining some service handler thingy
	// if !p.AuthorizationService.CanDeleteProfile(currUser) {
    //     return nil, errors.New("unauthorized")
    // }
// can also do something like this for the authorization
// if currUser.Role != model.Admin && currUser.UserID != profile.UserID {
// 	return "", errors.New("unauthorized: user is not permitted to delete this profile")
// }

	profile, err := p.UserRepository.GetById(p.context, currUser.UserID)
	if err != nil {
		return nil, errors.New("profile not found")
	}

	deletedProfile, err := p.UserRepository.Delete(p.context, profile)
	if err != nil {
		return nil, errors.New("failed to delete profile")
	}

	return &model.UserInfo{
		Username : deletedProfile.Username,
		Name      : deletedProfile.Name,
		Bio          : deletedProfile.Bio,
	}, nil
	
}

// GetProfile implements usecase.ProfileUseCase.
func (p *ProfileUseCase) GetProfile(currUser *model.AuthenticatedUser) (*model.UserInfo, error) {
	panic("unimplemented")
	//TODO: implement this usecase
}

// UpdateEmail implements usecase.ProfileUseCase.
func (p *ProfileUseCase) UpdateEmail(updated *model.UserUpdateEmail, currUser *model.AuthenticatedUser) (*model.UserInfo, error) {
	panic("unimplemented")
	//TODO: implement this usecase
}

// UpdatePassword implements usecase.ProfileUseCase.
func (p *ProfileUseCase) UpdatePassword(updated *model.UserUpdatePassword, currUser *model.AuthenticatedUser) (*model.UserInfo, error) {
	panic("unimplemented")
	//TODO: implement this usecase
}

// UpdateProfile implements usecase.ProfileUseCase.
func (p *ProfileUseCase) UpdateProfile(updated *model.UserUpdateProfile, currUser *model.AuthenticatedUser) (*model.UserInfo, error) {
	panic("unimplemented")
	//TODO: implement this usecase
}

// UpdateUsername implements usecase.ProfileUseCase.
func (p *ProfileUseCase) UpdateUsername(updated *model.UserUpdateUsername, currUser *model.AuthenticatedUser) (*model.UserInfo, error) {
	//TODO: implement this usecase
	panic("unimplemented")
}

// UpdateProfilePicture implements usecase.ProfileUseCase.

//TODO: add the profile picture update method
func (p *ProfileUseCase) UpdateProfilePicture(updated *model.UserUpdateProfilePicture, currUser *model.AuthenticatedUser) (*model.UserInfo, error) {
	panic("unimplemented")
	//TODO: implement this usecase
}

