package usecase

import (
	"BlogApp/domain"
	"BlogApp/domain/model"
)

type ShareUseCase interface {
	GetSharesByBlogID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*[]*model.BlogInfo, string, error)
	ShareBlogByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Share, string, error)
	UnshareBlogByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Share, string, error)
}