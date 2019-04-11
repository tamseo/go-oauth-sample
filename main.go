package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

type post struct {
	Title   string `json:"title"`
	Message string `json:"tag"`
}

func main() {
	r := mux.NewRouter()

	r.Handle("/public", public)
	r.Handle("/private", jwtMiddleware.Handler(private))

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("ListenAndServe:", nil)
	}
}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte("SIGNINGKEY"), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

var public = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &post{
		Title:   "Public API",
		Message: "Hi this is a public API",
	}
	json.NewEncoder(w).Encode(post)
})

var private = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &post{
		Title:   "Private API",
		Message: "Hi this is a private API",
	}
	json.NewEncoder(w).Encode(post)
})
