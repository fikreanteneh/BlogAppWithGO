package usecase

import (
	"BlogApp/config"
	"BlogApp/domain"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
	"BlogApp/utils"
	"context"
	"errors"
)

type AuthUseCase struct {
	context 	  context.Context
	environment    config.Environment
	UserRepository domain.UserRepository
}

func NewAuthUseCase(environment *config.Environment, userRepository *domain.UserRepository) usecase.AuthUseCase {
	return &AuthUseCase{
		environment:    *environment,
		UserRepository: *userRepository,
	}
}

// AdminRegister implements usecase.AuthUseCase.
func (a *AuthUseCase) AdminRegister(userCreate *model.UserCreate, currUser *model.AuthenticatedUser) (*model.UserInfo, error) {
	//TODO : Authorization Handling
	//TODO : Validation Handling

	password, err := utils.EncryptPassword(userCreate.Password)
	if err != nil {
		return nil, err
	}
	admin := &domain.User{
		Username: userCreate.Username,
		Email:    userCreate.Email,
		Password: password,
		Role:     "ADMIN",
		Bio: 	userCreate.Bio,
		Name: userCreate.Name,
	}
	createdAdmin, err := a.UserRepository.Create(a.context, admin)
	if err != nil {
		return nil, err
	}

	return &model.UserInfo{
		Username : createdAdmin.Username,
		Name      : createdAdmin.Name,
		Bio          : createdAdmin.Bio,
	}, nil
}

// Login implements usecase.AuthUseCase.
func (a *AuthUseCase) Login(userLogin *model.UserLogin) (string, error) {
	//TODO : Validation Handling
	user, err := a.UserRepository.GetByUsername(a.context, userLogin.Username)
	if err != nil {
		return "", err
	}
	if utils.ComparePasswords(user.Password, userLogin.Password) != nil {
		return "", errors.New("passwords do not match")
	}
	token, err := utils.TokenGenerate(&model.AuthenticatedUser{
		Username: user.Username,
		Role:     user.Role,
		Email:   user.Email,
		UserID:   user.UserID,
	}, a.environment.JwtSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}

// Register implements usecase.AuthUseCase.
func (a *AuthUseCase) Register(userCreate *model.UserCreate) (*model.UserInfo, error) {
	//TODO : Validation Handling
	password, err := utils.EncryptPassword(userCreate.Password)
	if err != nil {
		return nil, err
	}
	admin := &domain.User{
		Username: userCreate.Username,
		Email:    userCreate.Email,
		Password: password,
		Role:     "USER",
		Bio: 	userCreate.Bio,
		Name: userCreate.Name,
	}
	createdAdmin, err := a.UserRepository.Create(a.context, admin)
	if err != nil {
		return nil, err
	}

	return &model.UserInfo{
		Username : createdAdmin.Username,
		Name      : createdAdmin.Name,
		Bio          : createdAdmin.Bio,
	}, nil
}


