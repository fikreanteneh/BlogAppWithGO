package usecase

import (
	"BlogApp/config"
	"BlogApp/domain"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
	"context"
)

type RatingUseCase struct {
	context          context.Context
	environment      config.Environment
	ratingRepository domain.BlogRatingRepository
}

// DeleteRatingByBlogID implements usecase.RatingUseCase.
func (r *RatingUseCase) DeleteRatingByBlogID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.BlogRating, string, error) {
	panic("unimplemented")
}

// GetRatingsByBlogID implements usecase.RatingUseCase.
func (r *RatingUseCase) GetRatingsByBlogID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*[]*domain.BlogRating, string, error) {
	panic("unimplemented")
}

// RateBlogByID implements usecase.RatingUseCase.
func (r *RatingUseCase) RateBlogByID(currUser *model.AuthenticatedUser, dto *model.RatingCreate, param *model.IdParam) (*domain.BlogRating, string, error) {
	panic("unimplemented")
}

// UpdateRatingByBlogID implements usecase.RatingUseCase.
func (r *RatingUseCase) UpdateRatingByBlogID(currUser *model.AuthenticatedUser, dto *model.RatingCreate, param *model.IdParam) (*domain.BlogRating, string, error) {
	panic("unimplemented")
}

func NewRatingUseCase(context *context.Context, environment *config.Environment, ratingRepository *domain.BlogRatingRepository) usecase.RatingUseCase {
	return &RatingUseCase{
		context:          *context,
		environment:      *environment,
		ratingRepository: *ratingRepository,
	}
}
