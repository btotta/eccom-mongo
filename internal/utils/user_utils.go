package utils

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func ValidatePassword(password string, username string) bool {
	if len(password) < 8 {
		return false
	}

	if strings.Contains(password, " ") {
		return false
	}

	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			return false
		}
	}

	if IsSimpleSequence(password) {
		return false
	}

	if strings.Contains(password, username) {
		return false
	}

	var hasUppercase, hasLowercase, hasDigit, hasSpecial bool

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUppercase = true
		case unicode.IsLower(char):
			hasLowercase = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasUppercase && hasLowercase && hasDigit && hasSpecial
}

func IsSimpleSequence(password string) bool {
	for i := 0; i < len(password)-2; i++ {
		if password[i]+1 == password[i+1] && password[i+1]+1 == password[i+2] {
			return true
		}
	}
	return false
}

func HashBcrypt(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
