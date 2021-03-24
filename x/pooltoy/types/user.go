package types

import (
	"github.com/google/uuid"
)

var _ User = User{}

func NewUser(
	userAccount string,
	isAdmin bool,
	name string,
	email string,
) User {
	return User{
		Id:          uuid.New().String(),
		UserAccount: userAccount,
		IsAdmin:     isAdmin,
		Name:        name,
		Email:       email,
	}
}

//Marshal
//Unmarshal
//Validate
