package controller

import (
	"BlogApp/config"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"

	"github.com/gin-gonic/gin"
)

type RatingController struct {
	environment   config.Environment
	ratingUseCase usecase.RatingUseCase
}

func NewRatingController(environment *config.Environment, ratingUseCase *usecase.RatingUseCase) *RatingController {
	return &RatingController{
		environment:   *environment,
		ratingUseCase: *ratingUseCase,
	}
}

func (rc *RatingController) GetRatingsByBlogID(c *gin.Context) {
	GetHandler(c, rc.ratingUseCase.GetRatingsByBlogID, nil, &model.IdParam{ID: c.Param("blog_id")})
}

func (rc *RatingController) RateBlog(c *gin.Context) {
	PostHandler(c, rc.ratingUseCase.RateBlogByID, &model.RatingCreate{}, &model.IdParam{ID: c.Param("blog_id")})
}

func (rc *RatingController) UpdateRating(c *gin.Context) {
	PutHandler(c, rc.ratingUseCase.UpdateRatingByID, &model.RatingCreate{}, &model.IdParam{ID: c.Param("rating_id")})
}

func (rc *RatingController) DeleteRating(c *gin.Context) {
	DeleteHandler(c, rc.ratingUseCase.DeleteRatingByID, nil, &model.IdParam{ID: c.Param("rating_id")})
}