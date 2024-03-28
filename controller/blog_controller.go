package controller

import (
	"BlogApp/config"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"

	"github.com/gin-gonic/gin"
)

type BlogController struct{
	environment *config.Environment
	BlogUseCase usecase.BlogUseCase
}


func NewBlogController(environment *config.Environment, blogUseCase *usecase.BlogUseCase) *BlogController {
	return &BlogController{
		environment: environment,
		BlogUseCase: *blogUseCase,
	}

}

func (b *BlogController) CreateBlog(c *gin.Context){
    PostHandler(c, b.BlogUseCase.CreateBlog, &model.BlogCreate{}, nil)
}

func (b *BlogController)GetByBlogID(c *gin.Context){
    GetHandler(c, b.BlogUseCase.GetBlogByID, nil, &model.IdParam{ID: c.Param("blog_id")})
}

func (b *BlogController)GetAllBlogs(c *gin.Context){
    GetHandler(c, b.BlogUseCase.GetBlogs, nil, &model.SearchParam{Search: c.Query("search")})
}

func (b *BlogController)UpdateBlog(c *gin.Context){
    PutHandler(c, b.BlogUseCase.UpdateBlogByID, &model.BlogUpdate{}, nil)
}

func (b BlogController)DeleteBlog(c *gin.Context){
    DeleteHandler(c, b.BlogUseCase.DeleteBlogByID, nil, &model.IdParam{ID: c.Param("blog_id")})
}
