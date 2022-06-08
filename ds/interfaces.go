package ds

import (
	"atms/models"
)

type Account interface {
	GetUserAccounts() ([]models.Account, error)
	GetUserWithIDAccount() (models.Account, error)
	CreateUserAccount() error
	UpdateUserAccount() error
	DeleteUserAccount() error
}
