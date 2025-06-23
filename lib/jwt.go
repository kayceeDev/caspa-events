package lib

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/kayceeDev/caspa-events/pkg/config"
)

func SignJwtRecoveryToken(id int32, ExpiredAt time.Duration) (string, error) {
	env := config.EnvVars()

	expiredAt := time.Now().Add(ExpiredAt).Unix()

	jwtSecretKey := env.JWT_SECRET

	// metadata for your jwt
	claims := jwt.MapClaims{}
	claims["exp"] = expiredAt
	claims["authorization"] = true
	claims["userID"] = id
	return getToken(claims, jwtSecretKey)
}

type AcessTokenPayload struct {
	ID             int32
	FirstName      string
	LastName       string
	Email          string
	IsInterested   bool
	IsVendor       bool
	Pics           string
	UserUUID       uuid.UUID
	VendorID       int32
	BusinessHandle *string
}

func SignJwtAccessToken(payload AcessTokenPayload, ExpiredAt time.Duration) (string, error) {
	env := config.EnvVars()

	expiredAt := time.Now().Add(ExpiredAt).Unix()

	jwtSecretKey := env.JWT_SECRET

	// metadata for your jwt
	claims := jwt.MapClaims{}
	claims["exp"] = expiredAt
	claims["authorization"] = true
	claims["userID"] = payload.ID
	claims["firstName"] = payload.FirstName
	claims["lastName"] = payload.LastName
	claims["email"] = payload.Email
	claims["user_uuid"] = payload.UserUUID
	return getToken(claims, jwtSecretKey)
}

func getToken(claims jwt.MapClaims, secret string) (string, error) {
	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := to.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return token, nil
}

func GetHeaderToken(ctx *gin.Context) *string {
	tokenHeader := ctx.GetHeader("Authorization")
	headerSplit := strings.SplitAfter(tokenHeader, "Bearer ")

	if len(headerSplit) > 1 {
		accessToken := strings.SplitAfter(tokenHeader, "Bearer ")[1]
		return &accessToken
	}
	return nil

}

func VerifyJwtTokenHeader(ctx *gin.Context) (*jwt.Token, error) {
	accessToken := GetHeaderToken(ctx)

	if accessToken != nil {
		return VerifyToken(*accessToken)
	}
	return nil, errors.New("Error verifying token")

}

func VerifyToken(accessToken string) (*jwt.Token, error) {
	env := config.EnvVars()

	token, err := jwt.Parse(strings.Trim(accessToken, " "), func(token *jwt.Token) (interface{}, error) {
		return []byte(env.JWT_SECRET), nil
	})

	if token == nil {
		return nil, errors.New("empty token set")
	}

	switch {
	case token.Valid:
		return token, nil
	case errors.Is(err, jwt.ErrTokenMalformed):
		return nil, errors.New("malformed token")
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		return nil, errors.New("invalid signature")
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		return nil, errors.New("token expired")
	default:
		return nil, errors.New("Error verifying token")
	}
}

func DecodeJwtToken(accessToken, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}
