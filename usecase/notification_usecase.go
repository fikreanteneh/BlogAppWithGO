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
	notification, err := nuc.notificationRepository.GetByUserId(nuc.context, currUser.UserID)
	if err != nil {
		return nil, "Notification Fetch Failed", err
	}
	var notifications []*model.NotificationMessage
	for _, not := range *notification {
		notifications = append(notifications, &model.NotificationMessage{
			NotificationID: not.NotificationID,
			UserID: not.UserID,
			Content: not.Content,
			CreatedAt: not.CreatedAt,
		})
	}
	return &notifications, "Notification Fetched Successfully", nil

}


func (nuc *NotificationUseCase) DeleteNotificationByID(currUser *model.AuthenticatedUser, dto any, param *model.IdParam) (*model.NotificationMessage, string, error) {
	deletedNotification, err := nuc.notificationRepository.Delete(nuc.context, param.ID)
	if err != nil {
		return nil, "Notification Deletion Failed", err
	}
	return &model.NotificationMessage{
		NotificationID: deletedNotification.NotificationID,
		UserID: deletedNotification.UserID,
		Content: deletedNotification.Content,
		CreatedAt: deletedNotification.CreatedAt,
	}, "Notification Deletion Successfull", nil
}