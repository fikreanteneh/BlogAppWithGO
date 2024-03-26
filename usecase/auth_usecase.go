package usecase

import (
	"BlogApp/config"
	"BlogApp/domain"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
	"BlogApp/utils"
)

type AuthUseCase struct {
	environment    *config.Environment
	UserRepository *domain.UserRepository
}

func NewAuthUseCase(environment *config.Environment, userRepository *domain.UserRepository) usecase.AuthUseCase {
	return &AuthUseCase{
		environment:    environment,
		UserRepository: userRepository,
	}
}

// AdminRegister implements usecase.AuthUseCase.
func (a *AuthUseCase) AdminRegister(userCreate *model.UserCreate, currUser *model.AuthenticatedUser) (*model.UserList, error) {
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
		ProfilePicture: userCreate.ProfilePicture,
		Bio: 		  userCreate.Bio,
		
	}
	return t.TaskRepository.Create(c, task)
}

// Login implements usecase.AuthUseCase.
func (a *AuthUseCase) Login(userLogin *model.UserLogin) (string, error) {
	//TODO : Validation Handling
	
}

// Register implements usecase.AuthUseCase.
func (a *AuthUseCase) Register(userCreate *model.UserCreate) (*model.UserList, error) {
	//TODO : Validation Handling
}


