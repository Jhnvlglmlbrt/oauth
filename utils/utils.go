package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func GetFilePath() string {
	absPath, _ := filepath.Abs(".")
	baseDir := filepath.Dir(absPath)
	return baseDir
}

func GetGithubClientID() string {
	githubClientId, exists := os.LookupEnv("CLIENT_ID")
	if !exists {
		log.Fatal("Github ClientID not found in .env file")
	}
	return githubClientId
}

func GetGithubClientSecret() string {
	githubClientSecret, exists := os.LookupEnv("CLIENT_SECRET")
	if !exists {
		log.Fatal("Github ClientSecret not found in .env file")
	}
	return githubClientSecret
}

func GetGithubAccessToken(code string) string {
	clientID := GetGithubClientID()
	clientSecret := GetGithubClientSecret()

	requestBodyMap := map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"code":          code,
	}

	requestJSON, _ := json.Marshal(requestBodyMap)

	req, err := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		bytes.NewBuffer(requestJSON),
	)

	if err != nil {
		log.Panic("Request creation failed")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, perr := http.DefaultClient.Do(req)
	if perr != nil {
		log.Panic("Request failed")
	}

	respbody, _ := io.ReadAll(resp.Body)

	type githubAccessTokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}

	var ghresp githubAccessTokenResponse
	json.Unmarshal(respbody, &ghresp)

	return ghresp.AccessToken

}

func GetGithubData(accessToken string) string {
	req, err := http.NewRequest(
		"GET",
		"https://api.github.com/user",
		nil,
	)

	if err != nil {
		log.Panic("API Request creation failed")
	}

	authorizationHeaderValue := fmt.Sprintf("token %s", accessToken)
	req.Header.Set("Authorization", authorizationHeaderValue)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic("Request failed")
	}

	respbody, _ := io.ReadAll(resp.Body)

	return string(respbody)
}
