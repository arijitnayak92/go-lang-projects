package domain

type Notification struct {
	ID      string
	UserID  string
	RefID   string
	Message string
}

// NewNotification ....
func NewNotification(userID, refID, message string) *Notification {
	return &Notification{
		UserID:  userID,
		RefID:   refID,
		Message: message,
	}
}
