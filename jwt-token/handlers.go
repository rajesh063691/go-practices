package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-jose/go-jose/v4/json"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("secret_key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {

	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	fmt.Printf("Credentials = %#v", credentials)
	// check if the credentials are valid
	expectedPassword, ok := users[credentials.UserName]
	if !ok || expectedPassword != credentials.Password {
		http.Error(w, "Invalid password provided", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	// create the claims
	claims := &Claims{
		UserName: credentials.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// create the token using claims and secret-key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Error in generating token", http.StatusInternalServerError)
		return

	}

	// write the token in the cookie
	cookies := &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	}
	http.SetCookie(w, cookies)
	w.Write([]byte("jwt token created successfully !!! "))
}

func Home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "No cookie found", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Error in reading cookie", http.StatusBadRequest)
		return
	}

	tokenString := cookie.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Welcome %s !!!", claims.UserName)))

}

func Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "No cookie found", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Error in reading cookie", http.StatusBadRequest)
		return
	}

	tokenString := cookie.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}
	// check for the expiration time

	// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
	// 	http.Error(w, "token is not expired yet", http.StatusBadRequest)
	// 	return
	// }

	newexpireTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = newexpireTime.Unix()
	// create the new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Error in generating token", http.StatusInternalServerError)
		return

	}
	// write the token in the cookie
	cookies := &http.Cookie{
		Name:    "refresh_token",
		Value:   tokenString,
		Expires: newexpireTime,
	}
	http.SetCookie(w, cookies)

	w.Write([]byte("token refreshed successfully !!!"))
}
