package usecase

import (
	"BlogApp/domain"
	"BlogApp/domain/model"
)

type RatingUseCase interface {
	GetRatingsByBlogID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*[]*domain.BlogRating, string, error)
	RateBlogByID(currUser *model.AuthenticatedUser, dto *model.RatingCreate, param *model.IdParam) (*domain.BlogRating, string, error)
	UpdateRatingByID(currUser *model.AuthenticatedUser, dto *model.RatingCreate, param *model.IdParam) (*domain.BlogRating, string, error)
	DeleteRatingByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.BlogRating, string, error)
}