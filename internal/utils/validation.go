package utils

import "strings"

func IsValidEmail(email string) bool {
	// Simple email validation - consider using regex or proper validation package
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}
