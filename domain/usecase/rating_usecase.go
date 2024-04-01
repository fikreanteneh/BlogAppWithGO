package usecase

import (
	"BlogApp/domain"
	"BlogApp/domain/model"
)

type RatingUseCase interface {
	GetRatingsByBlogID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*[]*domain.BlogRating, string, error)
	RateBlogByID(currUser *model.AuthenticatedUser, dto *model.RatingCreate, param *model.IdParam) (*domain.BlogRating, string, error)
	DeleteRatingByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.BlogRating, string, error)
	
	// as a use case, we need to add a method to update a rating for a given post id
	UpdateRatingByID(currUser *model.AuthenticatedUser, dto *model.RatingCreate, param *model.IdParam) (*domain.BlogRating, string, error)


}