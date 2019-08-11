package handler

import (
	"encoding/json"
	"net/http"

	"github.com/george-kj/krsc/domain"
	"github.com/george-kj/krsc/domain/users"
)

//User is the user handler struct.
type User struct{}

// Login check user credentials and give back an authentication token.
func (u *User) Login(w http.ResponseWriter, r *http.Request) (interface{}, *domain.ResponseError) {

	// Read request body.
	usr := users.User{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		return nil, domain.NewBadRequestError(err, "Invalid Request Body")
	}

	// Validate user credentials.
	userID, errRes := usr.Login(appctx.Backend().UserRepo)
	if errRes != nil {
		return nil, errRes
	}

	return struct {
		UserID int64  `json:"status"`
		Token  string `json:"token"`
	}{userID, ""}, nil
}
