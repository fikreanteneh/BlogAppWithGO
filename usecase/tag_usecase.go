package usecase

import (
	"BlogApp/config"
	"BlogApp/domain"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
	"context"
	"errors"
)

type TagUseCase struct {
	context       context.Context
	environment   config.Environment
	tagRepository domain.TagRepository
}

// CreateTag implements usecase.TagUseCase.
func (t *TagUseCase) CreateTag(currUser *model.AuthenticatedUser, dto *model.TagCreate, param any) (*domain.Tag, string, error) {
	if currUser.Role != "ADMIN" {
		return nil, "Tag Creation Failed", errors.New("Unauthorized")
	}
	createdTag, err := t.tagRepository.Create(t.context, &domain.Tag{
		Name: dto.Name,
	})
	if err != nil {
		return nil, "Tag Creation Failed", err
	}
	return createdTag, "Tag Created Successfully", nil
}

// DeleteTagByID implements usecase.TagUseCase.
func (t *TagUseCase) DeleteTagByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Tag, string, error) {
	if currUser.Role != "ADMIN" {
		return nil, "Tag Creation Failed", errors.New("Unauthorized")
	}
	deletedTag, err := t.tagRepository.Delete(t.context, param.ID)
	if err != nil {
		return nil, "Tag Deletion Failed", err
	}
	return deletedTag, "Tag Deleted Successfully", nil
}

// GetTagByID implements usecase.TagUseCase.
func (t *TagUseCase) GetTagByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Tag, string, error) {
	fetchedTag, err := t.tagRepository.GetByID(t.context, param.ID)
	if err != nil {
		return nil, "Tag Not Found", err
	}
	return fetchedTag, "Tag Retrieved Successfully", nil
}

// GetTags implements usecase.TagUseCase.
func (t *TagUseCase) GetTags(currUser *model.AuthenticatedUser, dto any, param *model.SearchParam) (*[]*domain.Tag, string, error) {
	fetchedTags, err := t.tagRepository.GetAll(t.context, param.Search)
	if err != nil {
		return nil, "Tags Fetching Failed", err
	}
	return fetchedTags, "Tags Retrieved Successfully", nil
}

// UpdateTagByID implements usecase.TagUseCase.
func (t *TagUseCase) UpdateTagByID(currUser *model.AuthenticatedUser, dto *model.TagCreate, param *model.IdParam) (*domain.Tag, string, error) {
	if currUser.Role != "ADMIN" {
		return nil, "Tag Creation Failed", errors.New("Unauthorized")
	}
	updatedTag, err := t.tagRepository.Update(t.context, &domain.Tag{
		TagID: param.ID,
		Name:  dto.Name,
	})
	if err != nil {
		return nil, "Tag Updation Failed", err
	}
	return updatedTag, "Tag Updated Successfully", nil
}

func NewTagUseCase(context *context.Context, environment *config.Environment, tagRepository *domain.TagRepository) usecase.TagUseCase {
	return &TagUseCase{
		context:       *context,
		environment:   *environment,
		tagRepository: *tagRepository,
	}
}
