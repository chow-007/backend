package auth

import (
	"backend/configs"
	"backend/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTClaims struct {
	jwt.StandardClaims
	UserID   uint        `json:"user_id"`
	Username string      `json:"username"`
	Role     models.Role `json:"role"`
}

func (c *JWTClaims) SetExpiredAt(expiredAt int64) {
	c.ExpiresAt = expiredAt
}

func ObtainToken(user models.User) (string, error) {

	claims := JWTClaims{
		UserID:   user.ID,
		Username: user.UserName,
		//Role: user.Role,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.SetExpiredAt(time.Now().Add(time.Second * time.Duration(configs.Default.TokenExpireTime)).Unix())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(configs.Default.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func TokenRefresh(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.Default.SecretKey), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return "", errors.New("test")
	}
	if err := token.Claims.Valid(); err != nil {
		return "", err
	}

	//claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
	claims.IssuedAt = time.Now().Unix()
	claims.SetExpiredAt(time.Now().Add(time.Second * time.Duration(configs.Default.TokenExpireTime)).Unix())

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := newToken.SignedString([]byte(configs.Default.SecretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
