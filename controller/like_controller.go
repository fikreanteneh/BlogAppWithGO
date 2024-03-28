package controller

import (
	"BlogApp/config"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"

	"github.com/gin-gonic/gin"
)


type LikeController struct {
	environment config.Environment
	likeUseCase usecase.LikeUseCase
}

func NewLikeController(environment *config.Environment, likeUseCase *usecase.LikeUseCase) *LikeController {
	return &LikeController{
		environment: *environment,
		likeUseCase: *likeUseCase,
	}
}


func (lc *LikeController) GetLikesByBlogID(c *gin.Context) {
	GetHandler(c, lc.likeUseCase.GetLikesByBlogID, nil, &model.IdParam{ID: c.Param("blog_id")})
}

func (lc *LikeController) LikeBlog(c *gin.Context) {
	GetHandler(c, lc.likeUseCase.LikeBlogByID, nil, &model.IdParam{ID: c.Param("blog_id")})
}

func (lc *LikeController) UnlikeBlog(c *gin.Context) {
	GetHandler(c, lc.likeUseCase.UnlikeBlogByID, nil, &model.IdParam{ID: c.Param("like_id")})
}