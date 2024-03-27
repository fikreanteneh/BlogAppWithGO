package usecase

import (
	"BlogApp/domain"
	"BlogApp/domain/model"
)

type TagUseCase interface {
    GetTags(currUser *model.AuthenticatedUser) (*[]*domain.Tag, error)
    GetTagByID(id string, currUser *model.AuthenticatedUser) (*domain.Tag, error)
    CreateTag(tag *domain.Tag, currUser *model.AuthenticatedUser) (*domain.Tag, error)
    UpdateTagByID(id string, tag *domain.Tag, currUser *model.AuthenticatedUser) (*domain.Tag, error)
    DeleteTagByID(id string, currUser *model.AuthenticatedUser) (*domain.Tag, error)
}