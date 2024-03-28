package usecase

import (
	"BlogApp/config"
	"BlogApp/domain"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
	"context"
	"time"
)

type RatingUseCase struct {
	context          context.Context
	environment      config.Environment
	ratingRepository domain.BlogRatingRepository
}

// DeleteRatingByBlogID implements usecase.RatingUseCase.
func (r *RatingUseCase) DeleteRatingByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.BlogRating, string, error) {
	//TODO : Authorization Handling
	//TODO : Validation Handling
	deletedRating, err := r.ratingRepository.DeleteRating(r.context, param.ID)
	if err != nil {
		return nil,"Rating Deletion Failed", err
	}
	return deletedRating, "Rating Deletion Successfull", nil
	
}

// GetRatingsByBlogID implements usecase.RatingUseCase.
func (r *RatingUseCase) GetRatingsByBlogID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*[]*domain.BlogRating, string, error) {
	    ratings, err := r.ratingRepository.GetRatingByBlogID(r.context, param.ID)
    	if err != nil {
        	return nil,"Rating Fetch Failed", err
    	}
    	return ratings, "Rating Fetched SUccessfully", nil
}

// RateBlogByID implements usecase.RatingUseCase.
func (r *RatingUseCase) RateBlogByID(currUser *model.AuthenticatedUser, dto *model.RatingCreate, param *model.IdParam) (*domain.BlogRating, string, error) {
	    ratedBlog, err := r.ratingRepository.InsertRating(r.context, &domain.BlogRating{
        UserID: currUser.UserID,
        BlogID: param.ID,
        Rating: dto.Rating,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
    })
    if err != nil {
        return nil, "Rating Failed", err
    }
    return ratedBlog, "Rating Successfull", nil
}

// UpdateRatingByBlogID implements usecase.RatingUseCase.
func (r *RatingUseCase) UpdateRatingByID(currUser *model.AuthenticatedUser, dto *model.RatingCreate, param *model.IdParam) (*domain.BlogRating, string, error) {
	updatedRating, err := r.ratingRepository.UpdateRating(r.context, &domain.BlogRating{
		RatingID: param.ID,
        Rating: dto.Rating,
    })
    if err != nil {
        return nil,"Rating Updation Failed", err
    }
    return updatedRating,"Rating Updated Successfully", nil
}

func NewRatingUseCase(context *context.Context, environment *config.Environment, ratingRepository *domain.BlogRatingRepository) usecase.RatingUseCase {
	return &RatingUseCase{
		context:          *context,
		environment:      *environment,
		ratingRepository: *ratingRepository,
	}
}
