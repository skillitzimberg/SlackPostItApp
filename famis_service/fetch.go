package main

import (
	"bytes"
	"encoding/json"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const loginEndpoint = "/MobileWebServices/api/Login"
const jsonContentType = "application/json"
const usernameKey = "Username"
const passwordKey = "Password"
const loginItemKey = "Item"
const accessTokenKey = "access_token"
const refreshTokenKey = "refresh_token"
const expirationDateKey = ".expires"
const dateFormatConstant = "Fri, 22 Mar 2019 20:46:23 GMT"

type Fetcher struct {
	authToken    string
	refreshToken string
	expiration   time.Time
	_client      *http.Client
}

type FetchInputs struct {
	Username string
	Password string
	Url      string
	Endpoint string
	Filter   string
	Select   string
	// we don't allow then to pass the `$top` value
	// because we are going to page the results
}

func (Fetcher) Name() string {
	return "fetch_list"
}

func (Fetcher) Version() string {
	return "1.0"
}

func (fetch Fetcher) Execute(in step.Context) (interface{}, error) {
	input := FetchInputs{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (fetch Fetcher) Login(username, password, baseUrl string) (string, string, time.Time, error) {
	// validate and get login url
	loginUrl, err := getLoginUrl(baseUrl)
	if err != nil {
		return "", "", time.Time{}, err
	}
	// get login body from username and password
	loginBody, err := getReaderFromLoginBody(username, password)
	if err != nil {
		return "", "", time.Time{}, err
	}
	client := fetch.getHttpClient()
	resp, err := client.Post(loginUrl.String(), jsonContentType, loginBody)
	if err != nil {
		return "", "", time.Time{}, err
	}
	return handleLoginResponse(resp)
}

func handleLoginResponse(resp *http.Response) (string, string, time.Time, error) {
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", time.Time{}, err
	}
	data := make(JsonMap, 0)
	err = json.Unmarshal(contents, &data)
	itemData := data[loginItemKey]
	if itemData == nil {
		return "", "", time.Time{}, errors.New("invalid login response")
	}
	return handleLoginItem(itemData)
}

func handleLoginItem(item interface{}) (string, string, time.Time, error) {
	data, ok := item.(map[string]interface{})
	if !ok {
		return "", "", time.Time{}, errors.New("invalid login response")
	}

	accessToken, ok := data[accessTokenKey].(string)
	if !ok {
		return "", "", time.Time{}, errors.New("invalid access token")
	}

	refreshToken, ok := data[refreshTokenKey].(string)
	if !ok {
		return "", "", time.Time{}, errors.New("invalid refresh token")
	}

	expirationDateString, ok := data[expirationDateKey].(string)
	if !ok {
		return "", "", time.Time{}, errors.New("invalid expiration date value")
	}
	expirationDate, err := parseExpirationDate(expirationDateString)
	if err != nil {
		return "", "", time.Time{}, errors.New("invalid expiration expiration date value")
	}

	return accessToken, refreshToken, expirationDate, nil
}

func getReaderFromLoginBody(username, password string) (io.Reader, error) {
	loginContents, err := getLoginBody(username, password)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(loginContents), nil
}

func getLoginBody(username, password string) ([]byte, error) {
	body := make(JsonMap, 0)
	body[usernameKey] = username
	body[passwordKey] = password
	return getContents(body)
}

func getContents(data JsonMap) ([]byte, error) {
	return json.Marshal(data)
}

func getLoginUrl(base string) (*url.URL, error) {
	urlVal, err := url.Parse(base)
	if err != nil {
		return nil, err
	}
	urlVal.Path = loginEndpoint
	return urlVal, nil
}

func (fetch *Fetcher) getHttpClient() *http.Client {
	if fetch._client == nil {
		fetch._client = &http.Client{}
	}
	return fetch._client
}

func parseExpirationDate(date string) (time.Time, error) {
	return time.Parse(dateFormatConstant, date)
}
