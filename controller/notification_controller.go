package controller

import (
	"BlogApp/config"
	"BlogApp/domain/usecase"
	"BlogApp/middleware"
	"BlogApp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotificationController struct{
	environment *config.Environment
	NotificationUseCase usecase.NotificationUseCase
}


func NewNotificationController(environment *config.Environment, notificationUseCase *usecase.NotificationUseCase) *NotificationController {
	return &NotificationController{
		environment: environment,
		NotificationUseCase: *notificationUseCase,
	}

}

func (n *NotificationController) GetNotification(c *gin.Context){
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	notifications, err := n.NotificationUseCase.GetNotifications(currUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get notifications"})
		return
	}

	middleware.SuccessResponseHandler(c, 200, "Notifications retrieved successfully", notifications)

}

func (n *NotificationController) GetNotificationByID(c *gin.Context){
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	notificationID := c.Param("id")

	notification, err := n.NotificationUseCase.GetNotificationByID(notificationID, currUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get notification"})
		return
	}

	middleware.SuccessResponseHandler(c, 200, "Notification retrieved successfully", notification)
}

func (n *NotificationController) DeleteNotificationByID(c *gin.Context){
	currUser, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	notificationID := c.Param("id")

	notification, err := n.NotificationUseCase.DeleteNotificationByID(notificationID, currUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete notification"})
		return
	}

	middleware.SuccessResponseHandler(c, 200, "Notification deleted successfully", notification)
}
