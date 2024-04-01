package controller

import (
	"BlogApp/config"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"

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
	GetHandler(c, p.ProfileUseCase.GetProfile, nil, nil)
}

func (p *ProfileController) UpdateProfile(c *gin.Context) {
	PutHandler(c, p.ProfileUseCase.UpdateProfile, &model.UserUpdateProfile{}, nil)
}

func (p *ProfileController) DeleteProfile(c *gin.Context) {
	DeleteHandler(c, p.ProfileUseCase.DeleteProfile, nil, nil)
}

func (p *ProfileController) UpdatePassword(c *gin.Context) {
	PutHandler(c, p.ProfileUseCase.UpdatePassword, &model.UserUpdatePassword{}, nil)
}

func (p *ProfileController) UpdateEmail(c *gin.Context) {
	PutHandler(c, p.ProfileUseCase.UpdateEmail, &model.UserUpdateEmail{}, nil)
}

func (p *ProfileController) UpdateUsername(c *gin.Context) {
	PutHandler(c, p.ProfileUseCase.UpdateUsername, &model.UserUpdateUsername{}, nil)
}


//TODO: add the profile picture update method
func (p *ProfileController) UpdateProfilePicture(c *gin.Context) {

}
