package usecase

import (
	"BlogApp/config"
	"BlogApp/domain"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
	"BlogApp/utils"
	"context"
	"errors"
	"time"
)

type AuthUseCase struct {
	context        context.Context
	environment    config.Environment
	UserRepository domain.UserRepository
}

// AdminRegister implements usecase.AuthUseCase.
func (a *AuthUseCase) AdminRegister(currUser *model.AuthenticatedUser, userCreate *model.UserCreate, param any) (*model.UserInfo, string, error) {
		//TODO : Authorization Handling
	//TODO : Validation Handling

	password, err := utils.EncryptPassword(userCreate.Password)
	if err != nil {
		return nil, "User Creation Unseccssfull", err
	}
	admin := &domain.User{
		Username: userCreate.Username,
		Email:    userCreate.Email,
		Password: password,
		Role:     "ADMIN",
		Bio: 	userCreate.Bio,
		Name: userCreate.Name,
		CreatedAt: time.Now(),
	}
	createdAdmin, err := a.UserRepository.Create(a.context, admin)
	if err != nil {
		return nil, "User Creation Unseccssfull", err
	}

	return &model.UserInfo{
		UserId: createdAdmin.UserID,
		Username : createdAdmin.Username,
		Name      : createdAdmin.Name,
		Bio          : createdAdmin.Bio,
	},"Account Created Successfull", nil
}

// Login implements usecase.AuthUseCase.
func (a *AuthUseCase) Login(currUser *model.AuthenticatedUser, userLogin *model.UserLogin, parma any) (*model.Token, string, error) {
	//TODO : Validation Handling
	user, err := a.UserRepository.GetByUsername(a.context, userLogin.Username)
	if err != nil {
		return nil, "Login Unseccessfull",err
	}
	if utils.ComparePasswords(user.Password, userLogin.Password) != nil {
		return nil,"Login Unseccessfull", errors.New("passwords do not match")
	}
	token, err := utils.TokenGenerate(&model.AuthenticatedUser{
		Username: user.Username,
		Role:     user.Role,
		Email:   user.Email,
		UserID:   user.UserID,
	}, a.environment.JwtSecret)
	if err != nil {
		return nil, "Login Unseccessfull", err
	}
	return &model.Token{Token: token}, "Login Successfull", nil
}

// Register implements usecase.AuthUseCase.
func (a *AuthUseCase) Register(currUser *model.AuthenticatedUser, userCreate *model.UserCreate, param any) (*model.UserInfo, string, error) {
	//TODO : Validation Handling

	password, err := utils.EncryptPassword(userCreate.Password)
	if err != nil {
		return nil, "User Creation Unseccssfull", err
	}
	admin := &domain.User{
		Username: userCreate.Username,
		Email:    userCreate.Email,
		Password: password,
		Role:     "USER",
		Bio: 	userCreate.Bio,
		Name: userCreate.Name,
		CreatedAt: time.Now(),
	}
	createdAdmin, err := a.UserRepository.Create(a.context, admin)
	if err != nil {
		return nil, "User Creation Unseccssfull", err
	}

	return &model.UserInfo{
		UserId: createdAdmin.UserID,
		Username : createdAdmin.Username,
		Name      : createdAdmin.Name,
		Bio          : createdAdmin.Bio,
	},"Account Created Successfull", nil
}

func NewAuthUseCase(context *context.Context, environment *config.Environment, userRepository *domain.UserRepository) usecase.AuthUseCase {
	return &AuthUseCase{
		context:        *context,
		environment:    *environment,
		UserRepository: *userRepository,
	}
}