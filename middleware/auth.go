package middleware

import (
	"fmt"
	"net/http"
	"portfolio/service"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func getToken(c *gin.Context) (string, error) {
	err := fmt.Errorf("Error: %s", "invalid authrization")

	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", err
	}

	token := strings.Split(authHeader, " ")
	if token[0] != "Bearer" {
		return "", err
	}

	return token[1], nil
}

func parseToken(tokenString string) (parsedToken *jwt.Token, err error) {
	parsedToken, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("takurinton"), nil
	})

	if err != nil {
		return
	}
	if !parsedToken.Valid {
		return
	}
	return
}

// check the token here
// JWT token
func checkUser(token *jwt.Token) (isAuth bool) {
	user := service.User{}
	// Claimsだけではmapになっていないので注意
	// jwt.MapClaimsでキャストしてあげる必要がある
	_username := token.Claims.(jwt.MapClaims)["user"]
	username := fmt.Sprintf("%s", _username)
	_, err := user.Me(username)
	if err != nil {
		return
	}

	isAuth = true
	return
}

func AuthMiddlewere(c *gin.Context) {
	token, err := getToken(c)
	if err != nil {
		c.JSONP(http.StatusUnauthorized, gin.H{})
	}

	parsedToken, err := parseToken(token)
	if err != nil {
		c.JSONP(http.StatusUnauthorized, gin.H{})
	}

	isAuthUser := checkUser(parsedToken)
	if !isAuthUser {
		c.JSONP(http.StatusUnauthorized, gin.H{})
	}
}
