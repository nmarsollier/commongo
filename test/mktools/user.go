package mktools

import (
	"github.com/nmarsollier/commongo/security"
	uuid "github.com/satori/go.uuid"
)

func TestUser() *security.User {
	return &security.User{
		ID:          uuid.NewV4().String(),
		Login:       "Login",
		Name:        "Name",
		Permissions: []string{"user"},
	}
}

func TestAdminUser() *security.User {
	return &security.User{
		ID:          uuid.NewV4().String(),
		Login:       "Login",
		Name:        "Name",
		Permissions: []string{"user", "admin"},
	}
}
