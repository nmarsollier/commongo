package security

// User is the logged-in user
type User struct {
	ID          string   `json:"id"  validate:"required"`
	Name        string   `json:"name"  validate:"required"`
	Permissions []string `json:"permissions"`
	Login       string   `json:"login"  validate:"required"`
}

func (u *User) HasPermission(permission string) bool {
	for _, p := range u.Permissions {
		if p == permission {
			return true
		}
	}
	return false
}
