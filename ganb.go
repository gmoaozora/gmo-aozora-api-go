package ganblib

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gmoaozora/gmo-aozora-api-go/libs"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	try "gopkg.in/matryer/try.v1"
)

var config conf

// Ganb is our main struct
type Ganb struct {
	clientID     string
	clientSecret string
	nonce        nonce
}

type nonce struct {
	checkAndDelete func(string) error
	Save           func(string) error
}

// Token data sctucture for both oAuth and openID tokens
type Token struct {
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	Scope            string `json:"scope,omitempty"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	IDToken          string `json:"id_token,omitempty"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorUri         string `json:"error_uri"`
}

type conf struct {
	Salt        string `json:"SALT"`
	AuthBaseURL string `json:"AUTH_BASE_URL"`
	AuthPath    string `json:"AUTH_PATH"`
	TokenPath   string `json:"TOKEN_PATH"`
	JwtIssuer   string `json:"JWT_ISSUER"`
}

func init() {
	// maybe read and load env variables from conf file
	jsonFile, err := os.Open("./conf.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &config)
}

func getState(sessionID string, salt string) string {
	sum := sha256.Sum256([]byte(sessionID + salt))
	return fmt.Sprintf("%x", sum)
}

// New setup the lib
func New(clientID string, clientSecret string, nonceSave func(string) error, nonceCheck func(string) error) (Ganb, error) {
	n := nonce{nonceCheck, nonceSave}
	return Ganb{clientID, clientSecret, n}, nil
}

// authorizationHeader return a Basic authorization header
func (g Ganb) authorizationHeader() string {
	msg := fmt.Sprintf("%v:%v", g.clientID, g.clientSecret)
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	return fmt.Sprintf("Basic %v", encoded)
}

// OAuthAuthorization methods Authorization used with oAuth
func (g Ganb) OAuthAuthorization(sessionID string, scope string, redirectURI string) string {
	fmt.Println("### oAuth authorization")

	state := getState(sessionID, config.Salt)

	values := url.Values{
		"response_type": {"code"},
		"scope":         {scope},
		"state":         {state},
		"client_id":     {g.clientID},
		"redirect_uri":  {redirectURI},
	}

	return config.AuthBaseURL + config.AuthPath + "?" + values.Encode()
}

// OAuthGetToken methods used with oAuth
func (g Ganb) OAuthGetToken(redirectURI string, code string, authMethod string) (Token, error) {
	fmt.Println("\t- oAuth get token")

	var token Token

	tokenURL := fmt.Sprintf("%s%s", config.AuthBaseURL, config.TokenPath)
	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("redirect_uri", redirectURI)
	values.Add("code", code)

	authHeader := ""
	if authMethod == "basic" {
		authHeader = g.authorizationHeader()
	} else {
		values.Add("client_id", g.clientID)
		values.Add("client_secret", g.clientSecret)
	}

	body, err := libs.Request("post", tokenURL, values, authHeader)
	if err != nil {
		return token, errors.Wrap(err, "failed to Http Request.")
	}

	err = json.Unmarshal(body, &token)
	if err != nil {
		return token, errors.Wrap(err, "failed to unmarshal json.")
	}
	return token, nil
}

// OpenIDAuthorization methods Authorization used with OpenID
func (g Ganb) OpenIDAuthorization(sessionID string, scope string, redirectURI string) (string, error) {
	fmt.Println("### OpenID authorization")

	state := getState(sessionID, config.Salt)
	newNonce := uuid.New()

	if err := g.nonce.Save(newNonce.String()); err != nil {
		return "", errors.Wrap(err, "Could not save Nonce value")
	}

	values := url.Values{
		"response_type": {"code"},
		"scope":         {scope},
		"state":         {state},
		"client_id":     {g.clientID},
		"redirect_uri":  {redirectURI},
		"nonce":         {newNonce.String()},
	}

	return config.AuthBaseURL + config.AuthPath + "?" + values.Encode(), nil
}

// OpenIDGetToken methods used with OpenID
func (g Ganb) OpenIDGetToken(redirectURI string, code string, authMethod string) (Token, error) {
	fmt.Println("\t- OpenID get token")

	var token Token

	tokenURL := fmt.Sprintf("%s%s", config.AuthBaseURL, config.TokenPath)
	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("redirect_uri", redirectURI)
	values.Add("code", code)

	authHeader := ""
	if authMethod == "basic" {
		authHeader = g.authorizationHeader()
	} else {
		values.Add("client_id", g.clientID)
		values.Add("client_secret", g.clientSecret)
	}

	body, err := libs.Request("post", tokenURL, values, authHeader)
	if err != nil {
		return token, errors.Wrap(err, "failed to Http Request.")
	}

	err = json.Unmarshal(body, &token)
	if err != nil {
		return token, errors.Wrap(err, "failed to unmarshal json.")
	}

	err = g.isTokenValid(token.IDToken)
	if err != nil {
		return token, errors.Wrap(err, "IDToken is invalid.")
	}

	return token, nil
}

// Check if the token is a valid one with JWT and against the Nonce value
func (g Ganb) isTokenValid(IDToken string) error {
	if IDToken == "" {
		return errors.New("IDToken has no value")
	}
	token, err := jwt.Parse(IDToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(g.clientSecret), nil
	})
	if err != nil {
		return errors.Wrap(err, "failed to parse and validation IDToken")
	}

	nonce := token.Claims.(jwt.MapClaims)["nonce"].(string)
	err = g.nonce.checkAndDelete(nonce)
	if err != nil {
		return errors.Wrap(err, "nonce is invalid")
	}
	return nil
}

// RefreshToken refresh the session using a refresh token
func (g Ganb) RefreshTokens(refreshToken string) (Token, error) {
	fmt.Println("### refreshTokens")
	const CounterMaxTries = 3

	var token Token
	retryErr := try.Do(func(attempt int) (bool, error) {
		tokenURL := fmt.Sprintf("%v%v", config.AuthBaseURL, config.TokenPath)
		authHeader := g.authorizationHeader()

		values := url.Values{
			"client_id":     {g.clientID},
			"client_secret": {g.clientSecret},
			"grant_type":    {"refresh_token"},
			"refresh_token": {refreshToken},
		}

		body, err := libs.Request("post", tokenURL, values, authHeader)

		if err != nil {
			fmt.Printf("\n\nRefresh token request error: %v", err)
			return true, err
		}

		err = json.Unmarshal(body, &token)

		return attempt < CounterMaxTries, err // try 3 times
	})

	return token, retryErr
}
