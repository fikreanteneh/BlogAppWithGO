package usecase

import (
	"BlogApp/config"
	"BlogApp/domain"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
	"context"
)

type LikeUseCase struct {
	context        context.Context
	environment    config.Environment
	likeRepository domain.LikeRepository
}

// GetLikesByBlogID implements usecase.LikeUseCase.
func (l *LikeUseCase) GetLikesByBlogID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*[]*model.UserInfo, string, error) {
	panic("unimplemented")
}

// LikeBlogByID implements usecase.LikeUseCase.
func (l *LikeUseCase) LikeBlogByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Like, string, error) {
	panic("unimplemented")
}

// UnlikeBlogByID implements usecase.LikeUseCase.
func (l *LikeUseCase) UnlikeBlogByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Like, string, error) {
	panic("unimplemented")
}

func NewLikeUseCase(context *context.Context, environment *config.Environment, likeRepository *domain.LikeRepository) usecase.LikeUseCase {
	return &LikeUseCase{
		context:        *context,
		environment:    *environment,
		likeRepository: *likeRepository,
	}
}
