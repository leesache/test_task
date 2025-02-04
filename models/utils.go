package models

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CustomClaims defines the structure of the JWT claims
type CustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

// Custom error types
var (
	ErrInvalidReferral = errors.New("invalid referral")
	ErrSelfReferral    = errors.New("cannot refer yourself")
)

var JwtSecret []byte

func init() {
	jwtSecretEnv := os.Getenv("JWT_SECRET")
	if jwtSecretEnv == "" {
		log.Fatal("JWT_SECRET is not set")
	}
	JwtSecret = []byte(jwtSecretEnv)
}

// GenerateJWT creates a JWT token for the given user ID
func GenerateJWT(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &CustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}

// HashPassword hashes a plaintext password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// VerifyPassword verifies a plaintext password against a hashed password
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// IsNotFoundError checks if the error indicates a "not found" condition
func IsNotFoundError(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

// IsConflictError checks if the error indicates a conflict (e.g., invalid referral)
func IsConflictError(err error) bool {
	return errors.Is(err, ErrInvalidReferral) || errors.Is(err, ErrSelfReferral)
}

var (
	ErrTaskAlreadyCompleted = errors.New("task already completed")
)
