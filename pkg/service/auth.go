package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/max-sanch/AuthJWT/pkg/repository"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	signingKey = "njf&%df#24gfd6zk%#fds01bld4%df#24gfd"
	refreshTokenTTL = 12 * time.Hour
	accessTokenTTL = 15 * time.Minute
)

type tokenClaims struct {
	jwt.StandardClaims
	UserGUID string `json:"user_guid"`
}

type AuthService struct {
	repo repository.Authentication
}

func NewAuthService(repo repository.Authentication)	*AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) GenerateTokens(guid string) (map[string]string, error) {
	user, err := a.repo.GetUser(guid)
	if err != nil {
		return nil, err
	}

	timeNow := time.Now().Unix()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenTTL).Unix(),
			IssuedAt: timeNow,
		},
		user.GUID,
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(refreshTokenTTL).Unix(),
			IssuedAt: timeNow,
		},
		user.GUID,
	})

	accessTokenStr, err := accessToken.SignedString([]byte(signingKey))
	if err != nil {
		return nil, err
	}

	refreshTokenStr, err := refreshToken.SignedString([]byte(signingKey))
	if err != nil {
		return nil, err
	}

	refreshTokenHash, err := bcrypt.GenerateFromPassword([]byte(refreshTokenStr), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	err = a.repo.SetRefresh(guid, string(refreshTokenHash))
	if err != nil {
		return nil, err
	}

	tokens := map[string]string{
		"access_token": accessTokenStr,
		"refresh_token": refreshTokenStr,
	}

	return tokens, nil
}

func (a *AuthService) RefreshTokens(accessToken, refreshToken string) (map[string]string, error) {
	accessClaims, err := parseToken(accessToken)
	if err != nil {
		return nil, err
	}

	refreshClaims, err := parseToken(refreshToken)
	if err != nil {
		return nil, err
	}

	if accessClaims.UserGUID != refreshClaims.UserGUID || accessClaims.IssuedAt != refreshClaims.IssuedAt {
		return nil, errors.New("invalid tokens")
	}

	refresh, err := a.repo.GetRefresh(refreshClaims.UserGUID)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(refresh), []byte(refreshToken)); err != nil {
		return nil, errors.New("invalid tokens")
	}

	tokens, err := a.GenerateTokens(refreshClaims.UserGUID)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func parseToken(token string) (*tokenClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*tokenClaims)
	if !ok {
		return nil, errors.New("token claims are not of type *tokenClaims")
	}

	return claims, nil
}
