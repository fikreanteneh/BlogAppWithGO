package controller

import (
	"BlogApp/config"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
	"BlogApp/middleware"
	"net/http"

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
	var userCreate model.UserCreate

	if err := c.ShouldBindJSON(&userCreate); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	createdUser, err := a.AuthUseCase.Register(&userCreate)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	middleware.SuccessResponseHandler(c, http.StatusCreated, "User Created Successfully" ,createdUser)
}



func (a *AuthController) Login(c *gin.Context) {
	var userLogin model.UserLogin

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token, err := a.AuthUseCase.Login(&userLogin)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	middleware.SuccessResponseHandler(c, http.StatusOK, "User Logged In Successfully" ,token)
}

func (a *AuthController) AdminRegister(c *gin.Context) {
	var adminCreate model.UserCreate

	if err := c.ShouldBindJSON(&adminCreate); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	value, exists := c.Get("AuthenticatedUser")
	if !exists {
		c.JSON(400, gin.H{"error": "User not found in context"})
		return
	}
	currUser, ok := value.(*model.AuthenticatedUser)
	if !ok {
		c.JSON(400, gin.H{"error": "User not found in context"})
    	return
	}

	createdAdmin, err := a.AuthUseCase.AdminRegister(&adminCreate, currUser)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	middleware.SuccessResponseHandler(c, http.StatusCreated, "Admin Created Successfully" ,createdAdmin)
}