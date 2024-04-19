package middleware

import (
	"errors"
	"log/slog"
	"strings"
	"xrChat_backend/pkg"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// key
var jwtSecret = []byte("xrchat")

// Claims
type Claims struct {
	Id uint `json:"id"`
	jwt.StandardClaims
}

var errAuth = errors.New("token 出错")

// GenerateToken generate token containing id.
func GenerateToken(id uint) (string, error) {
	claim := Claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			Issuer: "xrChat",
		},
	}
	// generate token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// transform string.
	res, err := token.SignedString(jwtSecret)
	if err != nil {
		slog.Error("generate token err:%s", err)
		return "", errors.New("服务器错误，请稍后重试")
	}
	return res, nil
}

// ParseToken parse token.
func ParseToken(token string) (*Claims, error) {
	claims := &Claims{}
	withClaims, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, errAuth
	}
	if myclaims, ok := withClaims.Claims.(*Claims); ok && withClaims.Valid {
		return myclaims, nil
	}
	return nil, errAuth
}

// writeToken write token to response header to notif frontend whether the auth is successful.
func writeToken(c *gin.Context, token string) {
	c.Header("Authorization", "Bearer "+token)
}

// Jwt middleware of json-web-token .
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get token from request header.
		tokenHeader := c.GetHeader("Authorization")
		// not token.
		if tokenHeader == "" {
			writeToken(c, "false")
			pkg.HandleError(c, errAuth)
			c.Abort()
			return
		}
		// judge token format.
		tokenString := strings.SplitN(tokenHeader, " ", 2)
		if len(tokenString) < 2 || tokenString[0] != "Bearer" {
			writeToken(c, "false")
			pkg.HandleError(c, errAuth)
			c.Abort()
			return
		}
		// parse token.
		claims, err := ParseToken(tokenString[1])
		if err != nil {
			writeToken(c, "false")
			pkg.HandleError(c, err)
			c.Abort()
			return
		}
		writeToken(c, "true")
		// transfer id
		c.Set("id", claims.Id)
	}
}
