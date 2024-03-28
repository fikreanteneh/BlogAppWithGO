package usecase

import (
	"BlogApp/domain"
	"BlogApp/domain/model"
)

type CommentUseCase interface {
	GetCommentsByBlogID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*[]*domain.Comment, string, error)
	CreateCommentByBlogID(currUser *model.AuthenticatedUser, dto *model.CommentCreate, param *model.IdParam) (*domain.Comment, string, error)
	UpdateCommentByBlogID(currUser *model.AuthenticatedUser, dto *model.CommentCreate, param *model.IdParam) (*domain.Comment, string, error)
	DeleteCommentByBlogID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Comment, string, error)
}