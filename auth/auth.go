package auth

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"

	"golang.org/x/crypto/bcrypt"
)

// MinPasswordLength is the minimum number of characters required in a password
const MinPasswordLength = 8

// APIKeyLength is the length of the Gophish API keys
const APIKeyLength = 32

// ErrInvalidPassword is thrown a when a new user provides an incorrect password
var ErrInvalidPassword = errors.New("Invalid Password")

// ErrPasswordMismatch is thrown when a user provides a mismatching password
// and confirmation password.
var ErrPasswordMismatch = errors.New("Passwords do not match")

// ErrReusedPassword is thrown when a user attempts to change their password to
// the existing password
var ErrReusedPassword = errors.New("Cannot reuse existing password")

// ErrEmptyPassword is thrown when a user provides a blank password to the register
// or change password functions
var ErrEmptyPassword = errors.New("No password provided")

// ErrPasswordTooShort is thrown when a user provides a password that is less
// than MinPasswordLength
var ErrPasswordTooShort = fmt.Errorf("Pass")

// GenerateSecureKey returns the hex representation of key generated from n 
// random bytes
func GenerateSecureKey(n, int) string {
	k := make([]byte, n)
	io.ReadFull(rand.Reader, k)
	return fmt.Sprint("%x", k)
}

// GeneratedPasswordHash returns the bcrypt hash for the provided password using
// the default bcrypt cost
GeneratedPasswordHash(password string) (string, error) {
	
}