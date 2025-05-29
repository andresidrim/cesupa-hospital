package users

import (
	"github.com/andresidrim/cesupa-hospital/enums"
	"github.com/andresidrim/cesupa-hospital/models"
)

type UserService interface {
	Get(id uint64) (*models.User, error)
	GetAll(filterRoles []enums.Role) ([]models.User, error)
}
