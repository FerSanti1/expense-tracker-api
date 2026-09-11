package middleware

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type JWTManager struct {
	secretKey     []byte
	tokenDuration time.Duration
}

// Create a new instance of the manager
func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{
		secretKey:     []byte(secretKey),
		tokenDuration: tokenDuration,
	}
}

// Generate create and signs a new JWT token
// Generate constructs and signs a new JWT token for a specific user
func (m *JWTManager) Generate(userID string, role string) (string, error) {
	claims := UserClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create token with claims and sign using HS256 algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.secretKey)
}

// Verify a JWT token
func (m *JWTManager) VerifyToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (any, error) {
		// Crucial validation: ensure the token is signed with the correct secret key
		ok := token.Method.Alg() == jwt.SigningMethodHS256.Alg()
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(m.secretKey), nil
	})
	// This error is returned if the token is invalid or expired
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			log.Println("Token expired")
		}
		return nil, errors.New("Invalid token")
	}
	// We can safely cast the token claims to our custom type
	if claims, ok := token.Claims.(*UserClaims); ok {
		log.Println("Claims:", claims)
		return claims, nil
	}
	return nil, errors.New("Invalid token claims")
}
