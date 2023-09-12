package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Jhnvlglmlbrt/oauth/api"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file not found")
	}
}

func main() {
	http.HandleFunc("/", api.RootHandler)
	http.HandleFunc("/login/github/", api.GithubLoginHandler)
	http.HandleFunc("/login/github/redirect", api.GithubRedirectHandler)

	http.HandleFunc("/loggedin", func(w http.ResponseWriter, r *http.Request) {
		api.LoggedinHandler(w, r, "")
	})

	fmt.Println("[UP ON PORT 4000]")
	log.Panic(
		http.ListenAndServe(":4000", nil),
	)
}
