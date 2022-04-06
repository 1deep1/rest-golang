package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@user.ru",
		Username: "user",
		Password: "password",
	}
}
