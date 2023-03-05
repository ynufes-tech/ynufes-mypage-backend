package line

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"io"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

const (
	accessTokenEndpoint string = "https://api.line.me/oauth2/v2.1/token"
	verifyEndpoint      string = "https://api.line.me/oauth2/v2.1/verify"
	ErrorInvalidCode    string = "INVALID CODE"
)

type AuthVerifier struct {
	callbackURI  string
	clientID     string
	clientSecret string
	stateCache   map[string]int64
	random       *rand.Rand
}

func NewAuthVerifier(callbackURI, clientID, clientSecret string) AuthVerifier {
	return AuthVerifier{
		callbackURI:  callbackURI,
		clientID:     clientID,
		clientSecret: clientSecret,
		stateCache:   make(map[string]int64),
		random:       rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (v AuthVerifier) IssueNewState() string {
	var newState string
	newState = strconv.FormatUint(v.random.Uint64(), 10)
	for _, duplicate := v.stateCache[newState]; duplicate; {
		newState = strconv.FormatUint(rand.Uint64(), 10)
	}
	v.stateCache[newState] = time.Now().Unix()
	return newState
}

func (v AuthVerifier) verifyState(entry string) bool {
	r, res := v.stateCache[entry]
	if !res {
		return false
	}
	delete(v.stateCache, entry)
	if time.Now().Unix()-r > 120000 {
		return false
	}
	return true
}

func (v AuthVerifier) RevokeOldStates() {
	for s, t := range v.stateCache {
		if time.Now().Unix()-t > 120000 {
			delete(v.stateCache, s)
		}
	}
}

func (v AuthVerifier) RequestAccessToken(code string, state string) (*AccessTokenResponse, error) {
	if !v.verifyState(state) {
		return nil, errors.New(ErrorInvalidCode)
	}

	//prevent injection vulnerability
	reNum := regexp.MustCompile("^[0-9A-Za-z]+$")
	if !reNum.MatchString(code) {
		return nil, errors.New("INVALID CODE")
	}

	var credential = new(AccessTokenResponse)
	client := resty.New()
	req := client.R().
		SetFormData(map[string]string{
			"grant_type":    "authorization_code",
			"code":          code,
			"redirect_uri":  v.callbackURI,
			"client_id":     v.clientID,
			"client_secret": v.clientSecret,
		}).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetResult(&credential)
	_, err := req.Post(accessTokenEndpoint)
	if err != nil {
		log.Printf("Failed to request access token... %v", err)
		return nil, err
	}
	return credential, nil
}

func (v AuthVerifier) VerifyAccessToken(accessToken string) (*VerifyResponse, error) {
	verifyUri := verifyEndpoint + "?access_token=" + accessToken
	resp, err := http.Get(verifyUri)
	if err != nil {
		log.Printf("Failed to access verify endpoint... %v", err)
		return nil, err
	}
	var verifyResponse VerifyResponse
	body, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &verifyResponse)
	if err != nil {
		log.Printf("Failed to parse verify response... %v", err)
		return nil, err
	}
	return &verifyResponse, nil
}
