package server

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// GeneratePIN generates a random 4-digit PIN as a string.
func GeneratePIN() (string, error) {
	// Generate a number between 0 and 9999
	max := big.NewInt(10000)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	// Format as 4 digits with leading zeros
	return fmt.Sprintf("%04d", n.Int64()), nil
}

// TODO: AuthMiddleware could be added here to protect specific routes
