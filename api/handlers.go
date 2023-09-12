package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Jhnvlglmlbrt/oauth/utils"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<a href="/login/github/">LOGIN</a>`)
}

func GithubLoginHandler(w http.ResponseWriter, r *http.Request) {
	githubClientID := utils.GetGithubClientID()
	redirectURL := fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s",
		githubClientID,
		"http://localhost:3000/login/github/redirect",
	)
	http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)
}

func GithubRedirectHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	githubAccessToken := utils.GetGithubAccessToken(code)
	// fmt.Println("githubAccessToken:", githubAccessToken)
	githubData := utils.GetGithubData(githubAccessToken)
	LoggedinHandler(w, r, githubData)
}

func LoggedinHandler(w http.ResponseWriter, r *http.Request, githubData string) {
	if githubData == "" {
		fmt.Fprint(w, "Unauthorized access!")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var data interface{}
	if err := json.Unmarshal([]byte(githubData), &data); err != nil {
		log.Panic("Json Parse error: ", err)
	}

	formattedJson, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Panic("JSON format error: ", err)
	}

	fmt.Fprint(w, string(formattedJson))

}
