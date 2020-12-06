package gateways

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/julioc98/openbanking/pkg/entities"
	"github.com/julioc98/openbanking/pkg/interfaces"
	"github.com/julioc98/openbanking/pkg/jwt"
	"github.com/julioc98/openbanking/pkg/requesthelp"
)

const (
	accountAccessConsentsPath = "/account-access-consents"
	authPath                  = "/auth/realms/provider/protocol/openid-connect"
)

// Obiebank gateway
type Obiebank struct {
	authURL     string
	baseURL     string
	redirectURI string
	client      interfaces.Requester
}

// NewObiebank factory
func NewObiebank(authURL, baseURL, redirectURI string, client interfaces.Requester) *Obiebank {
	return &Obiebank{
		authURL:     authURL,
		baseURL:     baseURL,
		client:      client,
		redirectURI: redirectURI,
	}
}

// GetToken returns a token
func (g *Obiebank) GetToken(user, pass string) (*entities.Token, error) {
	uri := fmt.Sprint(g.authURL, authPath, "/token")

	form := url.Values{}
	form.Add("grant_type", "client_credentials")
	form.Add("scope", "accounts")

	req, err := http.NewRequest("POST", uri, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(user, pass)

	return g.requestToken(req)
}

// CreateAccountAccessConsents ...
func (g *Obiebank) CreateAccountAccessConsents(accountAccessConsents *entities.AccountAccessConsents, token string) (*entities.AccountAccessConsents, error) {
	uri := fmt.Sprint(g.baseURL, accountAccessConsentsPath)

	jsonBody, err := json.Marshal(accountAccessConsents)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, bytes.NewReader(jsonBody))
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	// The unique id of the ASPSP to which the request is issued. The unique id will be issued by OB.
	req.Header.Add("x-fapi-financial-id", "484848")
	requesthelp.SetAuthorization(req, token)
	resp, err := g.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Status response:%s", http.StatusText(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var res entities.AccountAccessConsents
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GenerateAuthURL return a URL auth
func (g *Obiebank) GenerateAuthURL(user, prvtk, consentID string) (string, error) {
	valueID := "urn:obiebank:accounts:" + consentID
	scope := "openid profile email accounts"
	state := 1125107479227
	stateStr := fmt.Sprintf("%d", state)
	nonce := int64(1400476011498)
	nonceStr := fmt.Sprintf("%d", nonce)
	uri := fmt.Sprint(g.authURL, authPath, "/auth")

	payload := &entities.JWTPayload{
		Aud:          fmt.Sprint(g.authURL, "/auth/realms/provider"),
		Iss:          user,
		ClientID:     user,
		RedirectURI:  g.redirectURI,
		Scope:        scope,
		Nonce:        nonce,
		Exp:          time.Now().Add(30 * 24 * time.Hour).Unix(),
		ResponseType: "code id_token",
		Claims: entities.Claims{
			Userinfo: entities.Userinfo{
				OpenbankingIntentID: entities.OpenbankingIntentID{
					Value:     valueID,
					Essential: true,
				},
			},
			IDToken: entities.IDToken{
				OpenbankingIntentID: entities.OpenbankingIntentID{
					Value:     valueID,
					Essential: true,
				},
				Acr: entities.Acr{
					Essential: true,
					Values: []string{
						"urn:openbanking:psd2:sca",
						"urn:openbanking:psd2:ca",
					},
				},
			},
		},
		Iat: time.Now().Unix(),
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	tokenJWT := jwt.NewJWT(prvtk)
	token, err := tokenJWT.Generate(jsonPayload)
	if err != nil {
		return "", err
	}

	params := url.Values{
		"response_type": {"code"},
		"client_id":     {user},
		"redirect_uri":  {g.redirectURI},
		"scope":         {scope},
		"state":         {stateStr},
		"nonce":         {nonceStr},
		"request":       {token},
	}
	url := fmt.Sprintf("%s?%s", uri, params.Encode())

	return url, nil
}

// GetTokenByCode returns a token
func (g *Obiebank) GetTokenByCode(user, pass, code string) (*entities.Token, error) {
	uri := fmt.Sprint(g.authURL, authPath, "/token")

	form := url.Values{}
	form.Add("grant_type", "authorization_code")
	form.Add("code", strings.TrimSpace(code))
	form.Add("client_id", user)
	form.Add("client_secret", pass)
	form.Add("redirect_uri", g.redirectURI)

	req, err := http.NewRequest("POST", uri, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return g.requestToken(req)

}

// using not to repeat code in the token call
func (g *Obiebank) requestToken(req *http.Request) (*entities.Token, error) {
	resp, err := g.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status response: %s | Error message: %s", http.StatusText(resp.StatusCode), body)
	}

	var token entities.Token
	err = json.Unmarshal(body, &token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}
