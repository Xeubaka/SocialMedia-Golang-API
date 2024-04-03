package security

import "golang.org/x/crypto/bcrypt"

// Hash changes a string to a encrypted Hash
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// ValidatePassword compare a password string to a encrypted hash of a password
func ValidatePassword(passwordHash, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordString))
}
