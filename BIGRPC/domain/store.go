package domain

import (
	"sync"

	"github.com/twinj/uuid"
)

type NotificationStore interface {
	AddNotification(userID, refID, message string) (string, error)
}
type AMSNotification struct {
	notifications []*Notification
	mutex         sync.RWMutex
}

func NewAMSNotification() *AMSNotification {
	return &AMSNotification{}
}

func (a *AMSNotification) AddNotification(userID, refID, message string) (string, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	notificationReq := NewNotification(userID, refID, message)
	notificationReq.ID = uuid.NewV4().String()
	a.notifications = append(a.notifications, notificationReq)

	return notificationReq.ID, nil
}
