package line

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"net/http"
	"regexp"
	"ynufes-mypage-backend/pkg/setting"
)

const (
	accessTokenEndpoint string = "https://api.line.me/oauth2/v2.1/token"
	verifyEndpoint      string = "https://api.line.me/oauth2/v2.1/verify"
)

func RequestAccessToken(code string) (*AccessTokenResponse, error) {

	//prevent injection vulnerability
	reNum := regexp.MustCompile("^[0-9A-Za-z]+$")
	if !reNum.MatchString(code) {
		return nil, errors.New("INVALID CODE")
	}

	config := setting.Get()
	var credential = new(AccessTokenResponse)
	client := resty.New()
	resp, err := client.R().
		SetFormData(map[string]string{
			"grant_type":    "authorization_code",
			"code":          code,
			"redirect_uri":  config.ThirdParty.LineLogin.CallbackURI,
			"client_id":     config.ThirdParty.LineLogin.ClientID,
			"client_secret": config.ThirdParty.LineLogin.ClientSecret,
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
