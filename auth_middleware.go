package main

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

type AuthError error

type LoginClaims struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var secretKey string = os.Getenv("SECRET_KEY")

func AuthMiddleware(next http.Handler) http.Handler {
	log.Println("MIDDLEWARE")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawToken, err := AuthBearerChecker(r)
		if err != nil {
			log.Panic(err)
			return
		}
		log.Println("rawToken Key: ", rawToken)
		log.Println("Secret Key: ", secretKey)
		loginClaims := &LoginClaims{Login: "", Password: ""}
		token, err := jwt.ParseWithClaims(rawToken, loginClaims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})
		log.Println("Signature: ", token.Signature)
		log.Println("Validad", token.Claims)
		claims := token.Claims
		log.Println(claims)
		if token.Valid {
			// Access context values in handlers like this
			// props, _ := r.Context().Value("props").(jwt.MapClaims)
			log.Println("LOGADO")
			next.ServeHTTP(w, r)
		} else {
			fmt.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		}
		log.Println(token)

		// next.ServeHTTP(w, r)

	})
}
func AuthBearerChecker(r *http.Request) (string, error) {
	head := strings.Split(r.Header.Get("Authorization"), "Bearer ")
	if len(head) != 2 {
		return "", errors.New("malformad token")
	}
	return head[1], nil
}
func GenerateToken(w http.ResponseWriter, r *http.Request) {
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

func MalformedTokenError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Malformad token"))
}
