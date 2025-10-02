package model

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

func requiredWithFallback(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	
	targetFieldName := fl.Param()
	
	parent := fl.Parent()

	if parent.Kind() == reflect.Ptr {
		parent = parent.Elem()
	}
	
	targetField := parent.FieldByName(targetFieldName)

	if !targetField.IsValid() {
		return password != ""
	}
	
	encryptedPassword := targetField.String()
	
	if encryptedPassword != "" {
		return true
	}
	
	if password == "" {
		return false
	}
	
	return validatePasswordComplexity(password)
}

func validatePasswordComplexity(password string) bool {
	if len(password) < 8 {
		return false
	}
	
	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false
	
	for _, char := range password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case 'a' <= char && char <= 'z':
			hasLower = true
		case '0' <= char && char <= '9':
			hasDigit = true
		case (33 <= char && char <= 47) || (58 <= char && char <= 64) || 
		     (91 <= char && char <= 96) || (123 <= char && char <= 126):
			hasSpecial = true
		}
	}
	
	return hasUpper && hasLower && hasDigit && hasSpecial
}