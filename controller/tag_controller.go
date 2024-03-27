package controller

import (
	"BlogApp/config"
	"BlogApp/domain"
	"BlogApp/domain/usecase"
	"BlogApp/utils"
	"net/http"

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
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	tags, err := t.TagUseCase.GetTags(currUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get tags"})
		return
	}

	c.JSON(http.StatusOK, tags)

}

func (t *TagController) GetTagByID(c *gin.Context){
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	tagID := c.Param("id")

	tag, err := t.TagUseCase.GetTagByID(tagID, currUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get tag"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": " tag recieved successfully",
		"tag":    tag,
	})

}

func (t *TagController) CreateTag(c *gin.Context){
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var tag domain.Tag
	if err := c.BindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	newTag, err := t.TagUseCase.CreateTag(&tag, currUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tag"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": " tag created successfully",
		"tag":    newTag,
	})

}

func (t *TagController) UpdateTagByID(c *gin.Context){
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	tagID := c.Param("id")

	var tag domain.Tag
	if err := c.BindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	updatedTag, err := t.TagUseCase.UpdateTagByID(tagID, &tag, currUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tag"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": " tag updated successfully",
		"tag":    updatedTag,
	})

}

func (t *TagController) DeleteTagByID(c *gin.Context){
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	tagID := c.Param("id")

	tag, err := t.TagUseCase.DeleteTagByID(tagID, currUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tag"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": " tag deleted successfully",
		"tag":    tag,
	})

}
