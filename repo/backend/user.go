package backend

import (
	"github.com/george-kj/krsc/domain"
)

// User function provide persistence methods for user repository.
type User struct {
	*Infra
}

// Login validate user credentials.
func (u *User) Login(user, pass string) (int64, *domain.ResponseError) {
	return -1, nil
}
