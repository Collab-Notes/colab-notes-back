package dto

// @name ErrorResponse
type ErrorResponse struct {
	Error string `json:"error" example:"Ocurrio un error inesperado en el servidor"`
}

// @name UpdateVaultRoleRequest
type UpdateVaultRoleRequest struct {
	Username string `json:"username" example:"juanito123"`
	NewRole  string `json:"new_role" example:"viewer"`
}

// @name MessageResponse
type MessageResponse struct {
	Message string `json:"message" example:"Rol actualizado exitosamente"`
}
