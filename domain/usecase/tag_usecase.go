package usecase

import (
    "BlogApp/domain"
)

type TagUseCase interface {
    GetTags() (*[]*domain.Tag, error)
    GetTagByID(id string) (*domain.Tag, error)
    CreateTag(tag *domain.Tag) error
    UpdateTagByID(id string, tag *domain.Tag) error
    DeleteTagByID(id string) error
}