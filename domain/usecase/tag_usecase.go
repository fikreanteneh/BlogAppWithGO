package usecase

import (
	"BlogApp/domain"
	"BlogApp/domain/model"
)

type TagUseCase interface {
    GetTags(currUser *model.AuthenticatedUser, dto any, param *model.SearchParam) (*[]*domain.Tag, string, error)
    GetTagByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Tag, string, error)
    CreateTag(currUser *model.AuthenticatedUser, dto *model.TagCreate, param any) (*domain.Tag, string, error)
    UpdateTagByID(currUser *model.AuthenticatedUser, dto *model.TagCreate, param *model.IdParam) (*domain.Tag, string, error)
    DeleteTagByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*domain.Tag, string, error)
}