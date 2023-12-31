package handlerRole

import (
	"sistem-pengelolaan-bank-sampah/dto"
	"sistem-pengelolaan-bank-sampah/models"
	"sistem-pengelolaan-bank-sampah/repositories"
)

type handlerRole struct {
	RoleRepository repositories.RoleRepository
}

func HandlerRole(roleRepository repositories.RoleRepository) *handlerRole {
	return &handlerRole{roleRepository}
}

func convertRoleResponse(role *models.MstRole) *dto.RoleResponse {
	return &dto.RoleResponse{
		ID:   role.ID,
		Role: role.Role,
	}
}

func convertMultipleRoleResponse(role *[]models.MstRole) *[]dto.RoleResponse {
	var roles []dto.RoleResponse

	for _, r := range *role {
		roles = append(roles, *convertRoleResponse(&r))
	}

	return &roles
}
