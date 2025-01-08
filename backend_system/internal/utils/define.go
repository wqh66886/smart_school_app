package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/wqh/smart/school/system/internal/domain"
	"github.com/wqh/smart/school/system/internal/errorx"
	"github.com/wqh/smart/school/system/internal/initiate"
	"golang.org/x/crypto/bcrypt"
	"time"
)

/**
* description:
* author: wqh
* date: 2025/1/8
 */

type JwtCustomClaims struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	jwt.RegisteredClaims
}

func GetPwd(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ComparePwd 比对密码
func ComparePwd(pwd1 string, pwd2 string) bool {
	// Returns true on success, pwd1 is for the database.
	err := bcrypt.CompareHashAndPassword([]byte(pwd1), []byte(pwd2))
	return err != nil
}

func ParseTime(date string) time.Time {
	if len(date) == 0 {
		return time.Now()
	}
	if t, err := time.Parse("2006-01-02", date); err == nil {
		return t
	}
	return time.Now()
}

func CreateAccessToken(user *domain.User) (accessToken string, err error) {
	claims := &JwtCustomClaims{
		Name: user.Username,
		ID:   user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(initiate.Expiry))),
			IssuedAt:  jwt.NewNumericDate(time.Now()), // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()), // 生效时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(initiate.SecretKey))
	if err != nil {
		return "", err
	}
	return t, err
}

func IsAuthorized(requestToken string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(requestToken, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(initiate.SecretKey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errorx.NewError(500, "that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errorx.NewError(500, "token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errorx.NewError(500, "token not active yet")
			} else {
				return nil, errorx.NewError(500, "couldn't handle this token")
			}
		}
	}
	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errorx.NewError(500, "couldn't handle this token")
}
