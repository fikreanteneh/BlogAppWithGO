package controller

import (
	"BlogApp/domain/model"
	"BlogApp/middleware"
	"BlogApp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)





func PostHandler[T any, U any, V any](
	c *gin.Context, 
	handler func(currUser *model.AuthenticatedUser, dto T, param U) (V, string, error),
	dto T,
	param U) {
    err := c.BindJSON(dto)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }
	currUser, _ := utils.CheckUser(c)
    result, message, err := handler(currUser, dto, param)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    middleware.SuccessResponseHandler(c, http.StatusCreated, message, result)
}



// PUT
func PutHandler[T any, U any, V any](
    c *gin.Context, 
    handler func(currUser *model.AuthenticatedUser, dto T, param U) (V, string, error),
    dto T,
    param U) {
    err := c.BindJSON(dto)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }
    currUser, _ := utils.CheckUser(c)
    result, message, err := handler(currUser, dto, param)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    middleware.SuccessResponseHandler(c, http.StatusOK, message, result)
}

func GetHandler[T any, U any, V any](
    c *gin.Context, 
    handler func(currUser *model.AuthenticatedUser, dto T, param U) (V, string, error),
    dto T,
    param U) {
    currUser, _ := utils.CheckUser(c)
    result, message, err := handler(currUser, dto, param)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    middleware.SuccessResponseHandler(c, http.StatusOK, message, result)
}

func DeleteHandler[T any, U any, V any](
    c *gin.Context, 
    handler func(currUser *model.AuthenticatedUser, dto T, param U) (V, string, error),
    dto T,
    param U) {
    currUser, _ := utils.CheckUser(c)
    result, message, err := handler(currUser, dto, param)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    middleware.SuccessResponseHandler(c, http.StatusOK, message, result)
}

