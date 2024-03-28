package usecase

import (
	"BlogApp/domain"
	"BlogApp/domain/model"
)

type RatingUseCase interface {
	GetRatingsByBlogID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*[]*domain.BlogRating, string, error)
	RateBlogByID(currUser *model.AuthenticatedUser, dto *model.RatingCreate, param *model.IdParam) (*domain.BlogRating, string, error)
	UpdateRatingByBlogID(currUser *model.AuthenticatedUser, dto *model.RatingCreate, param *model.IdParam) (*domain.BlogRating, string, error)
	DeleteRatingByBlogID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.BlogRating, string, error)
}