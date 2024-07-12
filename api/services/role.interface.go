package services

import (
	"crm-glonass/api/dto"
)

type RoleInterface interface {
	CreateRole(role *dto.Role) error
	ListRoles() ([]dto.RoleList, error)
}
