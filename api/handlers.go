package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jhnvlglmlbrt/oauth/utils"
	"github.com/gorilla/sessions"
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

var store = sessions.NewCookieStore([]byte(utils.GetSessionKey()))

func GithubRedirectHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	// fmt.Println("auth code:", code)
	githubAccessToken := utils.GetGithubAccessToken(code)
	// fmt.Println("githubAccessToken:", githubAccessToken)
	githubData := utils.GetGithubData(githubAccessToken)

	session, err := store.Get(r, "github-session")
	if err != nil {
		http.Error(w, "Session error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["githubData"] = githubData
	session.Save(r, w)

	http.Redirect(w, r, "/loggedin", http.StatusFound)

}

func LoggedinHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "github-session")
	githubData, ok := session.Values["githubData"].(string)

	if !ok {
		http.Error(w, "Unauthorized access!", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var data interface{}
	if err := json.Unmarshal([]byte(githubData), &data); err != nil {
		http.Error(w, "JSON Parse error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	formattedJson, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		http.Error(w, "JSON format error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(formattedJson)
}
