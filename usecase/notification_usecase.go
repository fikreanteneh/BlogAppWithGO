package usecase

import (
	"BlogApp/config"
	"BlogApp/domain"
	"BlogApp/domain/model"
	"BlogApp/domain/usecase"
	"context"
)

type NotificationUseCase struct {
	context        context.Context
	environment    config.Environment
	notificationRepository domain.NotificationRepository
}


func NewNotificationUseCase(context *context.Context, environment *config.Environment, notificationRepository *domain.NotificationRepository) usecase.NotificationUseCase {
	return &NotificationUseCase{
		context:        *context,
		environment:    *environment,
		notificationRepository: *notificationRepository,
	}
}


func (nuc *NotificationUseCase) GetNotifications(currUser *model.AuthenticatedUser, dto any, param any) (*[]*model.NotificationMessage, string, error) {
    panic("unimplemented")
}

func (nuc *NotificationUseCase) GetNotificationByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*model.NotificationMessage, string, error) {
    panic("unimplemented")
}

func (nuc *NotificationUseCase) DeleteNotificationByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*model.NotificationMessage, string, error) {
    panic("unimplemented")
}