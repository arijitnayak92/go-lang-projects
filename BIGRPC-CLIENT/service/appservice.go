package service

type Service struct {
	NotificationService NotificationService
}

func NewService(notificationService NotificationService) *Service {
	return &Service{NotificationService: notificationService}
}
