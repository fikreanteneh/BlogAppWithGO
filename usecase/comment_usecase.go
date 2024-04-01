package usecase

import (
	"BlogApp/config"
	"BlogApp/domain"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
	"context"
	"errors"
	"time"
)

type CommentUseCase struct {
	context           context.Context
	environment       config.Environment
	commentRepository domain.CommentRepository
}

// CreateCommentByBlogID implements usecase.CommentUseCase.
func (c *CommentUseCase) CreateCommentByBlogID(currUser *model.AuthenticatedUser, dto *model.CommentCreate, param *model.IdParam) (*domain.Comment, string, error) {
	createdComment, err := c.commentRepository.Create(c.context, &domain.Comment{
		UserID:  currUser.UserID,
		BlogID:  param.ID,
		Content: dto.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

	})
	if err != nil {
		return nil,"COmmaenting Faild", err
	}
	return createdComment,"Commenting Successful", nil
}

// DeleteCommentByBlogID implements usecase.CommentUseCase.
func (c *CommentUseCase) DeleteCommentByBlogID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Comment, string, error) {
	comment, err := c.commentRepository.GetByID(c.context, param.ID)
	if err != nil {
		return nil, "Comment Not Found", errors.New("Comment Not Found")
	}
	if currUser.Role != "ADMIN" && currUser.UserID != comment.UserID {
		return nil, "Unauthorized", errors.New("Unauthorized")
	}
	deletedComment, err := c.commentRepository.Delete(c.context, param.ID)
	if err != nil {
		return nil,"Comment Deletion Unseccssful", err
	}
	return deletedComment, "Comment deleted", nil
}

// GetCommentsByBlogID implements usecase.CommentUseCase.
func (c *CommentUseCase) GetCommentsByBlogID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*[]*domain.Comment, string, error) {
	fetchedComments, err := c.commentRepository.GetByBlogID(c.context, param.ID)
	if err != nil {
		return nil, "Comments Fetching Failed",err
	}
	return fetchedComments, "Comment Retrieved SUccessfully", nil
}

// UpdateCommentByBlogID implements usecase.CommentUseCase.
func (c *CommentUseCase) UpdateCommentByID(currUser *model.AuthenticatedUser, dto *model.CommentCreate, param *model.IdParam) (*domain.Comment, string, error) {
	comment, err := c.commentRepository.GetByID(c.context, param.ID)
	if err != nil {
		return nil, "Comment Not Found", errors.New("Comment Not Found")
	}
	if currUser.Role != "ADMIN" && currUser.UserID != comment.UserID {
		return nil, "Unauthorized", errors.New("Unauthorized")
	}
	updatedComment, err := c.commentRepository.Update(c.context, &domain.Comment{
		CommentID: param.ID,
		Content: dto.Content,
    })
    if err != nil {
        return nil, "Comment Updation Failed",err
    }
    return updatedComment,"Commment Updation Successfull", nil
}

func NewCommentUseCase(context *context.Context, environment *config.Environment, commentRepository *domain.CommentRepository) usecase.CommentUseCase {
	return &CommentUseCase{
		context:           *context,
		environment:       *environment,
		commentRepository: *commentRepository,
	}
}
