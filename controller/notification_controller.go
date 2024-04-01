package controller

import (
	"BlogApp/config"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
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
	GetHandler(c, n.NotificationUseCase.GetNotifications, nil, nil)
}

func (n *NotificationController) DeleteNotificationByID(c *gin.Context){
	DeleteHandler(c, n.NotificationUseCase.DeleteNotificationByID, nil, &model.IdParam{ID: c.Param("notification_id")})
}
