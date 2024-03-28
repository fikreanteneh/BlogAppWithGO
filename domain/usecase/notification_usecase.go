package usecase

import (
	"BlogApp/domain/model"
)

type NotificationUseCase interface {
    GetNotifications(currUser *model.AuthenticatedUser, dto any, param any) (*[]*model.NotificationMessage, string, error)
    GetNotificationByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*model.NotificationMessage, string, error)
    DeleteNotificationByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*model.NotificationMessage, string, error)
}