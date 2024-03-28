package controller

import (
	"BlogApp/config"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
	"github.com/gin-gonic/gin"
)

type CommentController struct {
    environment    config.Environment
    commentUseCase usecase.CommentUseCase
}

func NewCommentController(environment *config.Environment, commentUseCase *usecase.CommentUseCase) *CommentController {
    return &CommentController{
        environment: *environment,
        commentUseCase: *commentUseCase,
    }
}


func (cc *CommentController) GetCommentByBlogID(c *gin.Context) {
    GetHandler(c, cc.commentUseCase.GetCommentsByBlogID, nil, &model.IdParam{ID: c.Param("blog_id")})
}

func (cc *CommentController) CreateComment(c *gin.Context) {
    GetHandler(c, cc.commentUseCase.CreateCommentByBlogID, nil, &model.IdParam{ID: c.Param("blog_id")})
}

func (cc *CommentController) UpdateComment(c *gin.Context) {
    GetHandler(c, cc.commentUseCase.UpdateCommentByID, nil, &model.IdParam{ID: c.Param("comment_id")})
}

func (cc *CommentController) DeleteComment(c *gin.Context) {
    GetHandler(c, cc.commentUseCase.DeleteCommentByBlogID, nil, &model.IdParam{ID: c.Param("comment_id")})
}