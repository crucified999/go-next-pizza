package model

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id int `json:"id"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required_with_fallback=EncryptedPassword"`
	EncryptedPassword string `json:"-"`
}

func (u *User) Validate() error {
	v := validator.New()

	v.RegisterValidation("required_with_fallback", requiredWithFallback)

	return v.Struct(u)
}
 
func (u *User) BeforeCreate() error {
	if len(u.Password) != 0 {
		enc, err := encryptString(u.Password)

		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}

	return nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}