package usecase

import "github.com/Arsaev/BasiliskConnect/internal/entity"

// AdminUsecase implements the entity.AdminUsecase interface
type AdminUsecase struct {
}

// NewAdminUsecase creates a new instance of the AdminUsecase
func NewAdminUsecase() *AdminUsecase {
	return &AdminUsecase{}
}

// GetAdminByEmail retrieves an admin by email
func (a *AdminUsecase) GetAdminByEmail(email string) (*entity.Admin, error) {
	return nil, nil
}

// CreateAdmin creates a new admin
func (a *AdminUsecase) CreateAdmin(admin *entity.Admin) (*entity.Admin, error) {
	return nil, nil
}

// SetPassword sets a new password for an admin
func (a *AdminUsecase) SetPassword(adminId, password string) error {
	return nil
}

// ListAdmins retrieves all admins in the system
func (a *AdminUsecase) ListAdmins() ([]*entity.Admin, error) {
	return nil, nil
}

// RemoveAdmin removes an admin from the system
func (a *AdminUsecase) RemoveAdmin(adminId string) error {
	return nil
}
