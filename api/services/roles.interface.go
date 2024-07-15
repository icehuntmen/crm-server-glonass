package services

import (
	"crm-glonass/api/dto"
)

type RoleInterface interface {
	Create(role *dto.Role) error
	List() ([]dto.RoleList, error)
}
