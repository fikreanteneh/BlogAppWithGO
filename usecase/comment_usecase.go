package usecase

import (
	"BlogApp/config"
	"BlogApp/domain"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
	"context"
)

type CommentUseCase struct {
	context           context.Context
	environment       config.Environment
	commentRepository domain.CommentRepository
}

// CreateCommentByBlogID implements usecase.CommentUseCase.
func (c *CommentUseCase) CreateCommentByBlogID(currUser *model.AuthenticatedUser, dto *model.CommentCreate, param *model.IdParam) (*domain.Comment, string, error) {
	panic("unimplemented")
}

// DeleteCommentByBlogID implements usecase.CommentUseCase.
func (c *CommentUseCase) DeleteCommentByBlogID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Comment, string, error) {
	panic("unimplemented")
}

// GetCommentsByBlogID implements usecase.CommentUseCase.
func (c *CommentUseCase) GetCommentsByBlogID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*[]*domain.Comment, string, error) {
	panic("unimplemented")
}

// UpdateCommentByBlogID implements usecase.CommentUseCase.
func (c *CommentUseCase) UpdateCommentByBlogID(currUser *model.AuthenticatedUser, dto *model.CommentCreate, param *model.IdParam) (*domain.Comment, string, error) {
	panic("unimplemented")
}

func NewCommentUseCase(context *context.Context, environment *config.Environment, commentRepository *domain.CommentRepository) usecase.CommentUseCase {
	return &CommentUseCase{
		context:           *context,
		environment:       *environment,
		commentRepository: *commentRepository,
	}
}
