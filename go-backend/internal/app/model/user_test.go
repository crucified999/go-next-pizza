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
			name: "with encrypted password",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "encrypted_password"
				
				return u
			},
			isValid: true,
		},
		{
			name: "empty email",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""

				return u
			},
			isValid: false,
		},
		{
			name: "invalid email",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Email = "invalid_email"

				return u
			},
			isValid: false,
		},
		{
			name: "valid password",
			user: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "uncomplex password",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Password = "2007coolo"

				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""

				return u
			},
			isValid: false,
		},
		{
			name: "uncomplex password",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Password = "Coolo!_ww"

				return u
			},
			isValid: false,
		},
		{
			name: "uncomplex password",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Password = "2007Coolo"

				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Password = "2007_Co"

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

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)

	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.Password)
}