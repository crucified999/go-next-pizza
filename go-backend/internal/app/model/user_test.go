package model_test

import (
	"testing"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct{
		name string
		user func() *model.User
		isValid bool
	} {
		{
			name: "valid",
			user: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty email",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Email.Valid = false

				return u
			},
			isValid: false,
		},
		{
			name: "invalid email",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Email.String = "invalid_email"

				return u
			},
			isValid: false,
		},
		{
			name: "empty phone",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Phone = ""

				return u
			},
			isValid: false,
		},
		{
			name: "invalid phone",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Phone = "1234567890"

				return u
			},
			isValid: false,
		},
		{
			name: "valid phone",
			user: func() *model.User {
				u := model.TestUser(t)
				return u
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.user().Validate())
			} else {
				assert.Error(t, tc.user().Validate())
			}
		})
	}
}
