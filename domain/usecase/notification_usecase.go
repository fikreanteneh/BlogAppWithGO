package usecase

import (
	"BlogApp/domain/model"
)

type NotificationUseCase interface {
    GetNotifications(currUser *model.AuthenticatedUser) (*[]*model.NotificationMessage, error)
    GetNotificationByID(id string, currUser *model.AuthenticatedUser) (*model.NotificationMessage, error)
    DeleteNotificationByID(id string, currUser *model.AuthenticatedUser) (*model.NotificationMessage, error)
}