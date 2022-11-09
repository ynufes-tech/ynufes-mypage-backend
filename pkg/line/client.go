package line

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"regexp"
)

const (
	accessTokenEndpoint string = "https://api.line.me/oauth2/v2.1/token"
	verifyEndpoint      string = "https://api.line.me/oauth2/v2.1/verify"
)

func RequestAccessToken(code string) (*AccessTokenResponse, error) {
	clientId := os.Getenv("LINE_CLIENT_ID")
	clientSecret := os.Getenv("LINE_CLIENT_SECRET")
	redirectUri := os.Getenv("REDIRECT_URI")

	//prevent injection vulnerability
	reNum := regexp.MustCompile("^\\d+$")
	if !reNum.MatchString(code) {
		return nil, errors.New("INVALID CODE")
	}

	client := &http.Client{}
	uriIssueAccessToken := accessTokenEndpoint + "?grant_type=authorization_code&code=" + code
	uriIssueAccessToken += "&redirect_uri=" + redirectUri
	uriIssueAccessToken += "&client_id=" + clientId
	uriIssueAccessToken += "&client_secret=" + clientSecret
	req, err := http.NewRequest("GET", uriIssueAccessToken, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var credential AccessTokenResponse
	err = json.Unmarshal(body, &credential)
	if err != nil {
		return nil, err
	}
	return &credential, nil
}

func VerifyAccessToken(accessToken string) (*VerifyResponse, error) {
	verifyUri := verifyEndpoint + "?access_token=" + accessToken
	resp, err := http.Get(verifyUri)
	if err != nil {
		return nil, err
	}
	var verifyResponse VerifyResponse
	body, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &verifyResponse)
	if err != nil {
		return nil, err
	}
	return &verifyResponse, nil
}
