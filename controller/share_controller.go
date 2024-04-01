package controller

import (
	"BlogApp/config"
	"BlogApp/domain/usecase"
	"BlogApp/domain/model"

	"github.com/gin-gonic/gin"
)
type ShareController struct {
	environment  config.Environment
	ShareUseCase usecase.ShareUseCase
}


func NewShareController(environment *config.Environment, shareUseCase *usecase.ShareUseCase) *ShareController {
	return &ShareController{
		environment:  *environment,
		ShareUseCase: *shareUseCase,
	}
}


func (sc *ShareController) GetSharesByBlogID(c *gin.Context) {
	GetHandler(c, sc.ShareUseCase.GetSharesByBlogID, nil, &model.IdParam{ID: c.Param("blog_id")})
}

func (sc *ShareController) ShareBlog(c *gin.Context) {
	GetHandler(c, sc.ShareUseCase.ShareBlogByID, nil, &model.IdParam{ID: c.Param("blog_id")})
}

func (sc *ShareController) DeleteShare(c *gin.Context) {
	GetHandler(c, sc.ShareUseCase.UnshareBlogByID, nil, &model.IdParam{ID: c.Param("share_id")})
}