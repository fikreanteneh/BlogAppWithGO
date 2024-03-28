package controller

import (
	"BlogApp/config"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	environment *config.Environment
	userUseCase usecase.UserUseCase
}


func NewUserController(environment *config.Environment, userUseCase *usecase.UserUseCase) *UserController {
	return &UserController{
		environment: environment,
		userUseCase: *userUseCase,
	}
}

func (uc *UserController) GetUsers(c *gin.Context) {
	GetHandler(c, uc.userUseCase.GetUsers, nil, &model.SearchParam{Search: c.Query("search")})
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	GetHandler(c, uc.userUseCase.GetUserByID, nil, &model.IdParam{ID: c.Param("user_id")})
}

func (uc *UserController) GetFollowersByID(c *gin.Context) {
	GetHandler(c, uc.userUseCase.GetFollowersByID, nil, &model.IdParam{ID: c.Param("user_id")})
}

func (uc *UserController) FollowUserByID(c *gin.Context) {
	GetHandler(c, uc.userUseCase.FollowUserByID, nil, &model.IdParam{ID: c.Param("user_id")})
}

func (uc *UserController) UnfollowUserByID(c *gin.Context) {
	GetHandler(c, uc.userUseCase.UnfollowUserByID, nil, &model.IdParam{ID: c.Param("user_id")})
}

func (uc *UserController) GetFollowingsByID(c *gin.Context){
	GetHandler(c, uc.userUseCase.GetFollowingsByID, nil, &model.IdParam{ID: c.Param("user_id")})
}

func (uc *UserController) GetBlogsByID(c *gin.Context) {
	GetHandler(c, uc.userUseCase.GetBlogsByID, nil, &model.IdParam{ID: c.Param("user_id")})
}

func (uc *UserController) GetSharesByID(c *gin.Context) {
	GetHandler(c, uc.userUseCase.GetSharesByID, nil, &model.IdParam{ID: c.Param("user_id")})
}

func (uc *UserController) GetLikesByID(c *gin.Context) {
	GetHandler(c, uc.userUseCase.GetLikesByID, nil, &model.IdParam{ID: c.Param("user_id")})
}
