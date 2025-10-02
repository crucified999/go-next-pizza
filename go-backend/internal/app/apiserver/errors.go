package apiserver

import "errors"

var(
	ErrIncorrectEmailOrPassword = errors.New("Incorrect email or password")
	ErrNotAuthenticated = errors.New("Not authenticated")
)