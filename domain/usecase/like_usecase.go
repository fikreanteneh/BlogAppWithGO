package usecase

import (
	"BlogApp/domain"
	"BlogApp/domain/model"
)

type LikeUseCase interface {
	GetLikesByBlogID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*[]*model.UserInfo, string, error)
	LikeBlogByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Like, string, error)
	UnlikeBlogByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Like, string, error)
}