package controller

import (
	"BlogApp/config"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	environment *config.Environment
	AuthUseCase usecase.AuthUseCase
}


func NewAuthController(environment *config.Environment, authUseCase *usecase.AuthUseCase) *AuthController {
	return &AuthController{
		environment: environment,
		AuthUseCase: *authUseCase,
	}
}

func (a *AuthController) Register(c *gin.Context) {
	PostHandler(c, a.AuthUseCase.Register, &model.UserCreate{}, nil)
}

func (a *AuthController) Login(c *gin.Context) {
	PostHandler(c, a.AuthUseCase.Login, &model.UserLogin{}, nil)
}

func (a *AuthController) AdminRegister(c *gin.Context) {
	PostHandler(c, a.AuthUseCase.AdminRegister, &model.UserCreate{}, nil)
}