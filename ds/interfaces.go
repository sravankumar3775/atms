package ds

import (
	"atms/models"
)

type Account interface {
	GetUserAccounts() ([]models.Account, error)
	GetUserAccountByName() (models.Account, error)
	CreateUserAccount() error
	UpdateUserAccount() error
	DeleteUserAccount() error
}
