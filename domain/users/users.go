package users

import (
	"fmt"
	"strings"

	"github.com/george-kj/krsc/domain"
)

type repo interface {
	Login(user, pass string) (int64, *domain.ResponseError)
}

// User represent new user.
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login check user credentials and return an authentication token.
func (u *User) Login(r repo) (int64, *domain.ResponseError) {
	err := u.validate()
	if err != nil {
		return 0, domain.NewBadRequestError(err)
	}
	return -1, nil
}

func (u *User) validate() error {

	if strings.TrimSpace(u.Username) == "" {
		return fmt.Errorf("Empty username")
	}
	if strings.TrimSpace(u.Password) == "" {
		return fmt.Errorf("Empty password")
	}
	return nil
}
