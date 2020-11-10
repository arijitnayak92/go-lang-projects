package services

import (
	"github.com/arijitnayak92/taskAfford/REST/domain"
)

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}
