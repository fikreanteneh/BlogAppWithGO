package controller

import (
	"BlogApp/config"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
	"BlogApp/utils"
	"net/http"

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

	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var blogData *model.BlogCreate

    if err := c.BindJSON(&blogData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	// TODO: "Blog validation"

	blog, err := b.BlogUseCase.CreateBlog(blogData, currUser)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create blog"})
        return
    }

	c.JSON(http.StatusCreated, gin.H{
        "message": "Blog created successfully",
        "blog":    blog,
    })
}

func (b *BlogController)GetByBlogID(c *gin.Context){
	blogID := c.Param("id")
    
    blog, err := b.BlogUseCase.GetBlogByID(blogID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get blog by ID"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "message": "Blog retrieved successfully",
        "blog":    blog,
    })
}

func (b *BlogController)GetAllBlogs(c *gin.Context){
	blogs, err := b.BlogUseCase.GetAllBlogs()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all blogs"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "message": "All blogs retrieved successfully",
        "blogs":    blogs,
    })
}

func (b *BlogController)UpdateBlog(c *gin.Context){
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	blogID := c.Param("id")
    var updatedData model.BlogUpdateData
    if err := c.BindJSON(&updatedData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	updatedBlog, err := b.BlogUseCase.UpdateBlog(blogID, &updatedData, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "message": "Blog updated successfully",
        "blog":    updatedBlog,
    })
}

func (b BlogController)DeleteBlog(c *gin.Context){
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	blogID := c.Param("id")

	err = b.BlogUseCase.DeleteBlog(blogID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete blog"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}

func (b *BlogController)GetBlogLikes(c *gin.Context){
	blogID := c.Param("id")
    
    likes, err := b.BlogUseCase.GetBlogLikes(blogID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve blog likes"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "message": "Blog likes retrieved successfully",
        "likes":   likes,
    })
}

func (b *BlogController) LikeBlog(c *gin.Context){
	blogID := c.Param("id")
    
	// TODO : Does the user needs to be authenticated to like/unlike blogs
    err := b.BlogUseCase.LikeBlog(blogID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like blog"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Blog liked successfully"})
}

func (b *BlogController) UnlikeBlog(c *gin.Context)  {
	blogID := c.Param("id")
    
    err := b.BlogUseCase.UnlikeBlog(blogID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike blog"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Blog unliked successfully"})
}

func (b *BlogController) GetBlogComments(c *gin.Context) {
	blogID := c.Param("id")
    
    comments, err := b.BlogUseCase.GetBlogComments(blogID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve blog comments"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "message":  "Blog comments retrieved successfully",
        "comments": comments,
    })
}

func (b *BlogController) UpdateBlogComments(c *gin.Context){

	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	blogID := c.Param("id")
    
    var updatedComments model.Comment
    if err := c.BindJSON(&updatedComments); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    updatedComments, err = b.BlogUseCase.UpdateBlogComment(blogID, updatedComments, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog comments"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Blog comments updated successfully"})
}

func (b *BlogController) DeleteBlogComments(c *gin.Context){
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	blogID := c.Param("id")
    commentID := c.Param("comment_id")
    
    err := b.BlogUseCase.DeleteBlogComment(blogID, commentID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete blog comment"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Blog comment deleted successfully"})
}

func (b *BlogController) CommentOnBlog(c *gin.Context){
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	blogID := c.Param("id")
    
    var comment model.Comment
    if err := c.BindJSON(&comment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    addedComment, err := b.BlogUseCase.AddCommentToBlog(blogID, comment)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add comment to blog"})
        return
    }
    
    c.JSON(http.StatusCreated, gin.H{
        "message": "Comment added to blog successfully",
        "comment": addedComment,
    })
}

func (b *BlogController) GetBlogShares(c *gin.Context)  {
	blogID := c.Param("id")
    
    shares, err := b.BlogUseCase.GetBlogShares(blogID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve blog shares"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "message": "Blog shares retrieved successfully",
        "shares": shares,
    })
}

func (b *BlogController) ShareBlog(c *gin.Context){
	blogID := c.Param("id")
    currUser, err := utils.CheckUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

    err = b.BlogUseCase.ShareBlog(blogID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to share blog"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Blog shared successfully"})
}

func (b *BlogController) UnshareBlog(c *gin.Context){
	blogID := c.Param("id")
    currUser, err := utils.CheckUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

    err = b.BlogUseCase.UnshareBlog(blogID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unshare blog"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Blog unshared successfully"})
}

func (b *BlogController) GetBlogRatings(c *gin.Context){
	blogID := c.Param("id")
    
    ratings, err := b.BlogUseCase.GetBlogRatings(blogID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve blog ratings"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "message": "Blog ratings retrieved successfully",
        "ratings": ratings,
    })
}

func (b *BlogController) RateBlog(c *gin.Context){
	blogID := c.Param("id")
    currUser, err := utils.CheckUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

    err = b.BlogUseCase.RateBlog(blogID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to rate blog"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Blog rated successfully"})
}

func (b *BlogController) UpdateBlogRating(c *gin.Context){
	blogID := c.Param("id")
    currUser, err := utils.CheckUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

    err = b.BlogUseCase.UpdateBlogRating(blogID, currUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog rating"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Blog rating updated successfully"})
}