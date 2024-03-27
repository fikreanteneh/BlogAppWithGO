package controller

import (
	"BlogApp/config"
	"BlogApp/domain"
	"BlogApp/domain/usecase"
	"BlogApp/middleware"
	"BlogApp/utils"
	"net/http"

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
	var search = c.Query("search")
	users, err := uc.userUseCase.GetUsers(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := uc.userUseCase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) GetFollowersByID(c *gin.Context) {
	id := c.Param("id")
	followers, err := uc.userUseCase.GetFollowersByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, followers)
}

func (uc *UserController) FollowUserByID(c *gin.Context) {

	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var follow *domain.Follow
	id := c.Param("id")
	follow, err = uc.userUseCase.FollowUserByID(id, currUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	middleware.SuccessResponseHandler(c, 200, "successfully followed", follow)
}

func (uc *UserController) UnfollowUserByID(c *gin.Context) {
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	var unfollow *domain.Follow

	id := c.Param("id")
	unfollow ,err = uc.userUseCase.UnfollowUserByID(id, currUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	
	middleware.SuccessResponseHandler(c, 200, "successfully unfollowed", unfollow)
	
   }
}

func (uc *UserController) GetFollowingsByID(c *gin.Context){

	id := c.Param("id")
	followings, err := uc.userUseCase.GetFollowingsByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	middleware.SuccessResponseHandler(c, 200, "successfully retrieved followings", followings)
}

func (uc *UserController) GetBlogsByID(c *gin.Context) {
	id := c.Param("id")
	blogs, err := uc.userUseCase.GetBlogsByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	middleware.SuccessResponseHandler(c, 200, "successfully retrieved blogs", blogs)
}

func (uc *UserController) GetSharesByID(c *gin.Context) {
	id := c.Param("id")
	shares, err := uc.userUseCase.GetSharesByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	middleware.SuccessResponseHandler(c, 200, "successfully retrieved shares", shares)
}

func (uc *UserController) GetLikesByID(c *gin.Context) {
	id := c.Param("id")
	likes, err := uc.userUseCase.GetLikesByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	middleware.SuccessResponseHandler(c, 200, "successfully retrieved likes", likes)
}
