package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

type authService struct {
	userRepository *userRepository
}

func NewAuthService(userRepository *userRepository) *authService {
	return &authService{userRepository: userRepository}
}

func (service *authService) Login(ctx context.Context, login, password string) (token string, userID int64, err error) {
	user, err := service.userRepository.GetByLogin(ctx, login)
	if err != nil {
		return "", 0, fmt.Errorf("could not GetByLogin: %w", err)
	}
	providedPassword := service.getHashedPassword(password, user.Salt)
	if user.Password != providedPassword {
		return "", 0, fmt.Errorf("incorrect login or password")
	}
	token, err = service.emitNewToken(user, time.Minute*15)
	return token, user.ID, err
}

func (service *authService) IsTokenValid(ctx context.Context, rawToken string) (valid bool, userID int64, err error) {
	var claims Claims
	token, err := jwt.ParseWithClaims(rawToken, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return PublicKey, nil
	})
	if err != nil {
		if strings.Contains(err.Error(), "Token is expired") {
			return false, 0, nil
		}
		return false, 0, fmt.Errorf("jwt.Parse: %w", err)
	}
	return token.Valid, claims.UserID, nil
}

func (service *authService) getHashedPassword(password, salt string) string {
	hashedArray := sha256.Sum256([]byte(password + salt))
	return hex.EncodeToString(hashedArray[:])
}

func (service *authService) emitNewToken(user User, ttl time.Duration) (token string, err error) {
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "auth-service",
		},
		UserID: user.ID,
		Role:   "user",
	}
	newToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return newToken.SignedString(PrivateKey)
}
