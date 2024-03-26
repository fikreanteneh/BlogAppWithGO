package controller

import (
	"BlogApp/config"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
	"BlogApp/middleware"
	"BlogApp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileController struct{
	environment *config.Environment
	ProfileUseCase usecase.ProfileUseCase
}


func NewProfileController(environment *config.Environment, profileUseCase *usecase.ProfileUseCase) *ProfileController {
	return &ProfileController{
		environment: environment,
		ProfileUseCase: *profileUseCase,
	}

}

func (p *ProfileController) GetProfile(c *gin.Context) {
	c.Query("search")
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	profile, err := p.ProfileUseCase.GetProfile(currUser)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	middleware.SuccessResponseHandler(c, http.StatusOK, "Profile Retrieved Successfully", profile)

}

func (p *ProfileController) UpdateProfile(c *gin.Context) {
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var updated model.UserUpdateProfile
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	profile, err := p.ProfileUseCase.UpdateProfile(&updated, currUser)
	middleware.SuccessResponseHandler(c, http.StatusOK, "Profile Updated Successfully", profile)


}

func (p *ProfileController) DeleteProfile(c *gin.Context) {
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	profile, err := p.ProfileUseCase.DeleteProfile(currUser)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	middleware.SuccessResponseHandler(c, http.StatusOK, "Profile Deleted Successfully", profile)

}

func (p *ProfileController) UpdatePassword(c *gin.Context) {
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	var updated model.UserUpdatePassword
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	profile, err := p.ProfileUseCase.UpdatePassword(&updated, currUser)
	middleware.SuccessResponseHandler(c, http.StatusOK, "Profile Updated Successfully", profile)

}

func (p *ProfileController) UpdateEmail(c *gin.Context) {
		currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	var updated model.UserUpdateEmail
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	profile, err := p.ProfileUseCase.UpdateEmail(&updated, currUser)
	middleware.SuccessResponseHandler(c, http.StatusOK, "Profile Updated Successfully", profile)

}

func (p *ProfileController) UpdateUsername(c *gin.Context) {
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	var updated model.UserUpdateUsername
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	profile, err := p.ProfileUseCase.UpdateUsername(&updated, currUser)
	middleware.SuccessResponseHandler(c, http.StatusOK, "Profile Updated Successfully", profile)


}

func (p *ProfileController) UpdateProfilePicture(c *gin.Context) {

}
