package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	_ "github.com/dgrijalva/jwt-go"
)

type LoginClaims struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var secretKey string = os.Getenv("SECRET_KEY")

func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawToken, err := AuthBearerChecker(r)
		if err != nil {
			log.Panic(err)
			return
		}
		token := parseToken(rawToken)
		if token.Valid {
			next.ServeHTTP(w, r)
			return
		}
		fmt.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))

	})
}
func parseToken(rawToken string) *jwt.Token {
	token, err := jwt.ParseWithClaims(rawToken, &LoginClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		log.Panic(err)
	}
	return token
}

func AuthBearerChecker(r *http.Request) (string, error) {
	head := strings.Split(r.Header.Get("Authorization"), "Bearer ")
	if len(head) != 2 {
		return "", errors.New("malformad token")
	}
	return head[1], nil
}
func GenerateToken(w http.ResponseWriter, r *http.Request) {
	// Verificação de login deverá vir aqui nesse método

	claims := &LoginClaims{
		Login:    "asdffdas",
		Password: "asdffdsa",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(1)).Unix(),
			Issuer:    "AuthService",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(signedToken))
}
