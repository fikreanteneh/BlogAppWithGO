package controller

import (
	"BlogApp/config"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"

	"github.com/gin-gonic/gin"
)

type TagController struct{
	environment *config.Environment
	TagUseCase usecase.TagUseCase
}

func NewTagController(environment *config.Environment, tagUseCase *usecase.TagUseCase) *TagController {
	return &TagController{
		environment: environment,
		TagUseCase: *tagUseCase,
	}

}

func (t *TagController) GetTags(c *gin.Context){
	GetHandler(c, t.TagUseCase.GetTags, nil, &model.SearchParam{Search: c.Query("search")})
}

func (t *TagController) GetTagByID(c *gin.Context){
	GetHandler(c, t.TagUseCase.GetTagByID, nil, &model.IdParam{ID: c.Param("tag_id")})

}

func (t *TagController) CreateTag(c *gin.Context){
	PostHandler(c, t.TagUseCase.CreateTag, &model.TagCreate{}, nil)
}

func (t *TagController) UpdateTagByID(c *gin.Context){
	PutHandler(c, t.TagUseCase.CreateTag, &model.TagCreate{}, nil)
}

func (t *TagController) DeleteTagByID(c *gin.Context){
	DeleteHandler(c, t.TagUseCase.DeleteTagByID, nil, &model.IdParam{ID: c.Param("tag_id")})
}
