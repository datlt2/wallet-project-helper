package jwtHelper

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

// Function to hash password
func HashPassword(password string) (string, error) {
	// Generate a unique salt for each password
	salt := GenerateSalt()

	// Combine the password and salt, then generate the bcrypt hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Concatenate the salt and hashed password to store both in the database
	fullHashedPassword := fmt.Sprintf("%s$%s", salt, hashedPassword)

	return fullHashedPassword, nil
}

// Function to check password is hashed password
func CheckPasswordHash(password, fullHashedPassword string) bool {
	// Split the full hashed password into salt and hashed password
	parts := SplitHashedPassword(fullHashedPassword)

	// Combine the provided password and salt, then check against the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(parts[1]), []byte(password+parts[0]))
	return err == nil
}

// Function to generate salt value
func GenerateSalt() string {
	/// Create a byte slice to store the random data
	randomBytes := make([]byte, 10)

	// Read random data from the crypto/rand package
	_, err := rand.Read(randomBytes)
	if err != nil {
		return ""
	}

	// Convert the random bytes to a hexadecimal string
	salt := hex.EncodeToString(randomBytes)

	// Remove the character "$" from the salt
	salt = strings.ReplaceAll(salt, "$", "")

	return salt
}

// Function to split hashed password into array of string
func SplitHashedPassword(fullHashedPassword string) []string {
	// Split the full hashed password into salt and hashed password parts
	parts := strings.SplitN(fullHashedPassword, "$", 2)
	if len(parts) != 2 {
		return []string{"", ""}
	}
	return parts
}
