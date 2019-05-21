package ganblib

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	corporate "github.com/gmo-aozora-net-bank/go/api-go-openapi-example/gmo-aozora-api-go/corporateclient"
	personal "github.com/gmo-aozora-net-bank/go/api-go-openapi-example/gmo-aozora-api-go/personalclient"

	"github.com/gbrlsnchs/jwt/v2"
	"github.com/gmo-aozora-net-bank/go/api-go-openapi-example/libs"
	"github.com/google/uuid"
	try "gopkg.in/matryer/try.v1"
)

var config conf

const apiPersonalBasePath = `https://stg-api.gmo-aozora.com/ganb/api/personal/v1`
const apiCorporateBasePath = `https://stg-api.gmo-aozora.com/ganb/api/corporate/v1`

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
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope,omitempty"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	IDToken      string `json:"id_token,omitempty"`
	State        string `json:"state,omitempty"`
	Code         string `json:"code,omitempty"`
}

type conf struct {
	Salt        string `json:"SALT"`
	AuthBaseURL string `json:"AUTH_BASE_URL"`
	AuthPath    string `json:"AUTH_PATH"`
	TokenPath   string `json:"TOKEN_PATH"`
	JwtIssuer   string `json:"JWT_ISSUER"`
}

type jotToken struct {
	*jwt.JWT
	Nonce string `json:"nonce"`
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

// NewPersonalClient return a personal client API object
func NewPersonalClient(token Token) *personal.APIClient {
	cfg := personal.NewConfiguration()
	cfg.BasePath = apiPersonalBasePath
	cfg.AddDefaultHeader("x-refresh-token", token.RefreshToken)
	return personal.NewAPIClient(cfg)
}

// NewCorporateClient return a corporate client API object
func NewCorporateClient(token Token) *corporate.APIClient {
	cfg := corporate.NewConfiguration()
	cfg.BasePath = apiCorporateBasePath
	cfg.AddDefaultHeader("x-refresh-token", token.RefreshToken)
	return corporate.NewAPIClient(cfg)
}

// authorizationHeader return a Basic authorization header
func (g Ganb) authorizationHeader() string {
	msg := fmt.Sprintf("%v:%v", g.clientID, g.clientSecret)
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))

	// fmt.Println("== authorizationHeader")
	// fmt.Printf("\n client id + secret: %v", msg)
	// fmt.Printf("\n client encoded %v", encoded)
	return fmt.Sprintf("Basic %v", encoded)
}

// OAuthAuthorization methods Authorization used with oAuth
func (g Ganb) OAuthAuthorization(sessionID string, scope string, redirectURI string) string {
	fmt.Printf("### oAuth authorization")
	fmt.Println(config.AuthBaseURL)

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
func (g Ganb) OAuthGetToken(redirectURI string, code string, state string, authMethod string) Token {
	fmt.Println("\t- oAuth get token")

	var token Token
	values := url.Values{}
	values.Add("grant_type", "authorization_code")

	tokenURL := fmt.Sprintf("%s%s", config.AuthBaseURL, config.TokenPath)
	authHeader := ""

	if authMethod == "basic" {
		authHeader = g.authorizationHeader()
	} else {
		values.Add("client_id", g.clientID)
		values.Add("client_secret", g.clientSecret)
	}

	values.Add("redirect_uri", redirectURI)
	values.Add("code", code)

	body, err := libs.Request("post", tokenURL, values, authHeader)

	if err != nil {
		log.Fatalln(err)
	}

	jsonErr := json.Unmarshal(body, &token)

	// fmt.Println("OAUTH TOKEN token")
	// fmt.Printf("%s", body)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	token.Code = code
	token.State = state
	// fmt.Printf("%v", token)
	return token
}

// OpenIDAuthorization methods Authorization used with OpenID
func (g Ganb) OpenIDAuthorization(sessionID string, scope string, redirectURI string) (string, error) {
	// fmt.Printf("### OpenID authorization")
	// fmt.Println(config.AuthBaseURL)

	state := getState(sessionID, config.Salt)
	newNonce := uuid.New()

	if err := g.nonce.Save(newNonce.String()); err != nil {
		return "", errors.New("Could not save Nonce value")
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
func (g Ganb) OpenIDGetToken(redirectURI string, code string, state string) (Token, error) {
	var token Token
	values := url.Values{}

	tokenURL := fmt.Sprintf("%s%s", config.AuthBaseURL, config.TokenPath)
	authHeader := g.authorizationHeader()

	values.Add("grant_type", "authorization_code")
	values.Add("client_id", g.clientID)
	values.Add("client_secret", g.clientSecret)
	values.Add("code", code)
	values.Add("state", state)
	values.Add("redirect_uri", redirectURI)

	body, err := libs.Request("post", tokenURL, values, authHeader)

	if err != nil {
		log.Print(err)
	}

	jsonErr := json.Unmarshal(body, &token)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	err = g.isTokenValid(token.IDToken)

	if token.IDToken == "" || err != nil {
		return token, err
	}

	return token, nil
}

// Check if the token is a valid one with JWT and against the Nonce value
func (g Ganb) isTokenValid(IDToken string) error {
	// Check the validity of the JWT and it's properties
	validator := jwt.NewHS256(g.clientSecret)
	payload, sig, err := jwt.Parse(IDToken)

	if err != nil {
		return err
	}

	if err = validator.Verify(payload, sig); err != nil {
		return err
	}

	var jot jotToken
	if err = jwt.Unmarshal(payload, &jot); err != nil {
		return err
	}

	audValidator := jwt.AudienceValidator(g.clientID)
	issValidator := jwt.IssuerValidator(config.JwtIssuer)

	if err = jot.Validate(issValidator, audValidator); err != nil {
		return err
	}

	// Check the nonce
	if err = g.nonce.checkAndDelete(jot.Nonce); err != nil {
		return err
	}

	return nil
}

// RefreshSession refresh the session using a refresh token
func (g Ganb) RefreshSession(httpStatusCode int, err error, response *http.ResponseWriter, token Token) (Token, error) {
	// fmt.Printf("\n\n- RefreshSession\n")
	// fmt.Printf("\n- requestErr: %v", httpStatusCode)

	// log.Println("Server error message")
	// log.Println(string(err.(personal.GenericSwaggerError).Body()))

	// Get the error message from the response
	var errResponse personal.ErrorResponse

	if err != nil {
		errResponseBody := err.(personal.GenericSwaggerError).Body()
		json.Unmarshal([]byte(errResponseBody), &errResponse)
	}

	// token provided is expired, revoked, malformed, or invalid
	// 401 error
	if !(httpStatusCode == http.StatusUnauthorized && errResponse.ErrorCode == "WG_ERR_105") {
		return token, nil
	}

	const CounterMaxTries = 3

	retryErr := try.Do(func(attempt int) (bool, error) {
		// fmt.Printf("\n- trying to refresh token. try #%v", attempt)
		// fmt.Printf("\n  token access #%v", token.AccessToken)
		// fmt.Printf("\n  token refresh #%v", token.RefreshToken)

		var err error

		tokenURL := fmt.Sprintf("%v%v", config.AuthBaseURL, config.TokenPath)
		authHeader := g.authorizationHeader()

		values := url.Values{
			"client_id":     {g.clientID},
			"client_secret": {g.clientSecret},
			"grant_type":    {"refresh_token"},
			"refresh_token": {token.RefreshToken},
		}

		body, err := libs.Request("post", tokenURL, values, authHeader)

		if err != nil {
			fmt.Printf("\n\nRefresh token request error: %v", err)
			return true, err
		}

		err = json.Unmarshal(body, &token)

		// fmt.Println("Refresh TOKEN value: ")
		// fmt.Printf("%s", body)

		return attempt < CounterMaxTries, err // try 3 times
	})

	return token, retryErr
}
