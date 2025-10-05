package service

import (
	"errors"
	"time"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/storage"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthService struct {
	userRepo         storage.UserRepository
	jwtSecret        []byte
	jwtTTLSeconds    int
	refreshTTLSeconds int
}

type AuthTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

func NewAuthService(
	userRepo storage.UserRepository,
	jwtSecret []byte,
	jwtTTLSeconds int,
	refreshTTLSeconds int,
) *AuthService {
	return &AuthService{
		userRepo:          userRepo,
		jwtSecret:         jwtSecret,
		jwtTTLSeconds:     jwtTTLSeconds,
		refreshTTLSeconds: refreshTTLSeconds,
	}
}

func (as *AuthService) Register(user *model.User) (*model.User, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	createdUser, err := as.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	createdUser.Sanitize()
	return createdUser, nil
}

func (as *AuthService) Login(email, password string) (*AuthTokens, error) {
	user, err := as.userRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !user.ComparePassword(password) {
		return nil, errors.New("invalid credentials")
	}

	tokens, err := as.generateTokens(user.Id)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (as *AuthService) generateTokens(userID int) (*AuthTokens, error) {
	now := time.Now()
	
	access := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": now.Add(time.Duration(as.jwtTTLSeconds) * time.Second).Unix(),
		"iat": now.Unix(),
		"jti": uuid.New().String(),
	})

	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": now.Add(time.Duration(as.refreshTTLSeconds) * time.Second).Unix(),
		"iat": now.Unix(),
		"typ": "refresh",
		"jti": uuid.New().String(),
	})

	accessToken, err := access.SignedString(as.jwtSecret)
	if err != nil {
		return nil, err
	}

	refreshToken, err := refresh.SignedString(as.jwtSecret)
	if err != nil {
		return nil, err
	}

	return &AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    as.jwtTTLSeconds,
	}, nil
}

func (as *AuthService) ValidateToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return as.jwtSecret, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userID, ok := claims["sub"].(float64); ok {
			return int(userID), nil
		}
	}

	return 0, errors.New("invalid token")
}