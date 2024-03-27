package usecase

import(
    "BlogApp/domain"
)

type NotificationUseCase interface {
    GetNotifications() (*[]*domain.Notification, error)
    GetNotificationByID(id string) (*domain.Notification, error)
    DeleteNotificationByID(id string) error
}