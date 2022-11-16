package line

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"net/http"
	"os"
	"regexp"
)

const (
	accessTokenEndpoint string = "https://api.line.me/oauth2/v2.1/token"
	verifyEndpoint      string = "https://api.line.me/oauth2/v2.1/verify"
	EnvLineClientId     string = "LINE_CLIENT_ID"
	EnvLineClientSecret string = "LINE_CLIENT_SECRET"
	EnvLineRedirectUri  string = "LINE_REDIRECT_URI"
)

func RequestAccessToken(code string) (*AccessTokenResponse, error) {
	clientId := os.Getenv(EnvLineClientId)
	clientSecret := os.Getenv(EnvLineClientSecret)
	redirectUri := os.Getenv(EnvLineRedirectUri)

	//prevent injection vulnerability
	reNum := regexp.MustCompile("^[0-9A-Za-z]+$")
	if !reNum.MatchString(code) {
		return nil, errors.New("INVALID CODE")
	}

	var credential = new(AccessTokenResponse)
	client := resty.New()
	resp, err := client.R().
		SetFormData(map[string]string{
			"grant_type":    "authorization_code",
			"code":          code,
			"redirect_uri":  redirectUri,
			"client_id":     clientId,
			"client_secret": clientSecret,
		}).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetResult(credential).
		Post(accessTokenEndpoint)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	return credential, nil
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
