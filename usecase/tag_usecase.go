package usecase

import (
	"BlogApp/config"
	"BlogApp/domain"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
	"context"
)

type TagUseCase struct {
	context       context.Context
	environment   config.Environment
	tagRepository domain.TagRepository
}

// CreateTag implements usecase.TagUseCase.
func (t *TagUseCase) CreateTag(currUser *model.AuthenticatedUser, dto *model.TagCreate, param any) (*domain.Tag, string, error) {
	panic("unimplemented")
}

// DeleteTagByID implements usecase.TagUseCase.
func (t *TagUseCase) DeleteTagByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Tag, string, error) {
	panic("unimplemented")
}

// GetTagByID implements usecase.TagUseCase.
func (t *TagUseCase) GetTagByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Tag, string, error) {
	panic("unimplemented")
}

// GetTags implements usecase.TagUseCase.
func (t *TagUseCase) GetTags(currUser *model.AuthenticatedUser, dto any, param *model.SearchParam) (*[]*domain.Tag, string, error) {
	panic("unimplemented")
}

// UpdateTagByID implements usecase.TagUseCase.
func (t *TagUseCase) UpdateTagByID(currUser *model.AuthenticatedUser, dto *model.TagCreate, param *model.IdParam) (*domain.Tag, string, error) {
	panic("unimplemented")
}

func NewTagUseCase(context *context.Context, environment *config.Environment, tagRepository *domain.TagRepository) usecase.TagUseCase {
	return &TagUseCase{
		context:       *context,
		environment:   *environment,
		tagRepository: *tagRepository,
	}
}
