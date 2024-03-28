package usecase

import (
	"BlogApp/config"
	"BlogApp/domain"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
	"context"
)

type ShareUseCase struct {
	context         context.Context
	environment     config.Environment
	shareRepository domain.ShareRepository
}

// GetSharesByBlogID implements usecase.ShareUseCase.
func (s *ShareUseCase) GetSharesByBlogID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*[]*model.BlogInfo, string, error) {
	panic("unimplemented")
}

// ShareBlogByID implements usecase.ShareUseCase.
func (s *ShareUseCase) ShareBlogByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Share, string, error) {
	panic("unimplemented")
}

// UnshareBlogByID implements usecase.ShareUseCase.
func (s *ShareUseCase) UnshareBlogByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Share, string, error) {
	panic("unimplemented")
}

func NewShareUseCase(context *context.Context, environment *config.Environment, shareRepository *domain.ShareRepository) usecase.ShareUseCase {
	return &ShareUseCase{
		context:         *context,
		environment:     *environment,
		shareRepository: *shareRepository,
	}
}
