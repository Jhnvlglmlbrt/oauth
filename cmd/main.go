package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Jhnvlglmlbrt/oauth/api"
	"github.com/Jhnvlglmlbrt/oauth/utils"
	"github.com/joho/godotenv"
)

func init() {
	envFile := utils.GetFilePath() + "/.env"
	if err := godotenv.Load(envFile); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	http.HandleFunc("/", api.RootHandler)
	http.HandleFunc("/login/github/", api.GithubLoginHandler)
	http.HandleFunc("/login/github/redirect", api.GithubRedirectHandler)

	http.HandleFunc("/loggedin", func(w http.ResponseWriter, r *http.Request) {
		api.LoggedinHandler(w, r, "")
	})

	fmt.Println("[UP ON PORT 3000]")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}
