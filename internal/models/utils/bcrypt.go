package models

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a plain password
func HashPassword(password string) (string, error) {
	// Generate the hash with bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// Return the hashed password as a string
	return string(hash), nil
}

// ValidatePassword checks if the plain password matches the hashed password
func ValidatePassword(hashedPassword, plainPassword string) bool {
	// Compare the hashed password with the plain password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return false
	}
	return true
}
