package controller

import (
	"BlogApp/config"
	"BlogApp/domain/usecase"
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
	users, err := uc.userUseCase.GetUsers()
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
	id := c.Param("id")
	err := uc.userUseCase.FollowUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "followed"})
}

func (uc *UserController) UnfollowUserByID(c *gin.Context) {
	id := c.Param("id")
	err := uc.userUseCase.UnfollowUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "unfollowed"})
}

func (uc *UserController) GetFollowingsByID(c *gin.Context) {
	id := c.Param("id")
	followings, err := uc.userUseCase.GetFollowingsByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, followings)
}

func (uc *UserController) GetBlogsByID(c *gin.Context) {
	id := c.Param("id")
	blogs, err := uc.userUseCase.GetBlogsByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blogs)
}

func (uc *UserController) GetSharesByID(c *gin.Context) {
	id := c.Param("id")
	shares, err := uc.userUseCase.GetSharesByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, shares)
}

func (uc *UserController) GetLikesByID(c *gin.Context) {
	id := c.Param("id")
	likes, err := uc.userUseCase.GetLikesByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, likes)
}
