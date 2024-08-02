package entity

// Admin is a struct that represents an admin user entity in the system.
// Admins are users created by the system that have the ability to manage the system
type Admin struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"` // salted hashed password
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

// AdminUsecase is an interface that defines all business logic interactions with the Admin entity
type AdminUsecase interface {
	GetAdminByEmail(email string) (*Admin, error)
	CreateAdmin(admin *Admin) (*Admin, error)
	SetPassword(adminId, password string) error
	ListAdmins() ([]*Admin, error)
	RemoveAdmin(adminId string) error
}
