package main

import (
	"bytes"
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	json "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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
const filterKey = "$filter"
const selectKey = "$select"
const topKey = "$top"
const skipKey = "$skip"
const size = 1000
const fetchMethod = "GET"
const authTokenHeaderKey = "Authorization"
const contentTypeHeaderKey = "Content-Type"
const bearer = "Bearer"
const nextLinkKey = "@odata.nextLink"
const valueKey = "value"
const defaultTimeout = time.Minute * 1

type Fetcher struct {
	authToken    string
	refreshToken string
	expiration   time.Time
	client       *http.Client
	username     string
	password     string
	url          string
}

type FetchInputs struct {
	Username  string
	Password  string
	Url       string
	Endpoint  string
	Filter    string
	Select    string
	ChunkSize int
	// we don't allow then to pass the `$top` value
	// because we are going to page the results
}

type FetchListOutputs struct {
	Count   int
	Records []JsonMap
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
	return fetch.execute(input)
}

func (fetch Fetcher) ExecuteJson(js string) (interface{}, error) {
	input := FetchInputs{}
	err := json.Unmarshal([]byte(js), &input)
	if err != nil {
		return nil, err
	}
	return fetch.execute(input)
}

func (fetch Fetcher) execute(input FetchInputs) (interface{}, error) {
	// set the username and pass on the fetcher
	fetch.username = input.Username
	fetch.password = input.Password
	fetch.url = input.Url
	// first thing we need to do is login to the FAMIS services
	// we require Username, Password, and Url so we do not have to validate the values
	var err error
	fetch.authToken, fetch.refreshToken, fetch.expiration, err = fetch.Login(input.Username, input.Password, input.Url)
	if err != nil {
		return nil, err
	}

	// if user didn't specify a chunk size
	// we are going to default to 1000
	if input.ChunkSize == 0 {
		input.ChunkSize = size
	}
	return fetch.FetchList(input.Url, input.Endpoint, input.Filter, input.Select, input.ChunkSize)
}

func (fetch Fetcher) FetchList(url, endpoint, filter, sel string, size int) (FetchListOutputs, error) {
	fetchUrl, err := getUrl(url, endpoint)
	if err != nil {
		return FetchListOutputs{}, err
	}
	// add filter query param if it applies
	if len(filter) > 0 {
		fetchUrl.Query().Add(filterKey, filter)
	}
	// add select query param if it applies
	if len(sel) > 0 {
		fetchUrl.Query().Add(selectKey, sel)
	}

	// add top filter
	fetchUrl.Query().Add(topKey, strconv.Itoa(size))
	// build fetch filter
	request := getFetchRequest(fetchUrl, fetch.authToken)

	records := make([]JsonMap, 0)
	data, err := fetch.exhaustFetch(request, records)
	if err != nil {
		return FetchListOutputs{}, err
	}
	return FetchListOutputs{Records: data, Count: len(data)}, nil
}

func (fetch Fetcher) exhaustFetch(req http.Request, records []JsonMap) ([]JsonMap, error) {
	// refreshed auth if we need to
	if err := fetch.RefreshIfNeeded(); err != nil {
		return nil, err
	}
	client := fetch.getHttpClient()
	// attempt fetch
	resp, err := client.Do(&req)
	if err != nil {
		return nil, err
	}
	// convert response from FAMIS to JsonMap
	data, err := handleFetchResponse(resp)
	// should be an array of interfaces
	values := data[valueKey]
	//
	if data, ok := values.([]interface{}); ok {
		addValuesToRecords(data, &records)
	}

	// if we have the next link value we know we need to continue fetching
	// else we assume we're done
	nextLink := data[nextLinkKey]
	if link, ok := nextLink.(string); ok {
		// the link is sent with http
		// and you can't connect to the famis service with http
		// replace http with https
		link = strings.Replace(link, "http", "https", 1)
		newUrl, err := url.Parse(link)
		if err != nil {
			return nil, err
		}
		// set the new url on link
		// and call self
		req.URL = newUrl
		return fetch.exhaustFetch(req, records)
	} else {
		// if we don't have a link we assume we're done
		return records, nil
	}
}

func (fetch *Fetcher) LogMeIn() error {
	auth, ref, expire, err := fetch.Login(fetch.username, fetch.password, fetch.url)
	if err != nil {
		return err
	}
	fetch.authToken = auth
	fetch.refreshToken = ref
	fetch.expiration = expire
	return nil
}

func (fetch Fetcher) Login(username, password, baseUrl string) (string, string, time.Time, error) {
	// validate and get login url
	loginUrl, err := getUrl(baseUrl, loginEndpoint)
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

func (fetch Fetcher) RefreshIfNeeded() error {
	now := time.Now()
	// time subtract a minute
	now = now.Add(time.Duration(-1) * 1)
	// we're NOT expired
	if !now.After(fetch.expiration) {
		// do nothing
		return nil
	} else {
		// login again and set values on fetcher
		err := fetch.LogMeIn()
		if err != nil {
			return err
		}
		return nil
	}
}

func addValuesToRecords(values []interface{}, records *[]JsonMap) {
	if len(values) < 1 {
		return
	}

	for _, val := range values {
		// TODO what do we do if we can't convert?
		data, ok := val.(map[string]interface{})
		if ok {
			// append value to values of pointer
			*records = append(*records, data)
		} else {
			// TODO no op?
		}
	}
}

func getFetchRequest(url *url.URL, authTok string) http.Request {
	fetchHeaders := getFetchHeaders(authTok, jsonContentType)
	return http.Request{
		Method: fetchMethod,
		URL:    url,
		Header: fetchHeaders,
	}
}

func getFetchHeaders(auth, content string) http.Header {
	headers := http.Header{}
	// we need to append `Bearer` to the auth token
	headers.Add(authTokenHeaderKey, fmt.Sprintf("%s %s", bearer, auth))
	headers.Add(contentTypeHeaderKey, content)
	return headers
}

func handleFetchResponse(resp *http.Response) (JsonMap, error) {
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := make(JsonMap, 0)
	err = json.Unmarshal(contents, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
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
		panic(err)
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

func getUrl(base, endpoint string) (*url.URL, error) {
	urlVal, err := url.Parse(base)
	if err != nil {
		return nil, err
	}
	urlVal.Path = endpoint
	return urlVal, nil
}

func (fetch *Fetcher) getHttpClient() *http.Client {
	if fetch.client == nil {
		fetch.client = &http.Client{
			Timeout: defaultTimeout,
		}
	}
	return fetch.client
}

func parseExpirationDate(date string) (time.Time, error) {
	return time.Parse(time.RFC1123, strings.TrimSpace(date))
}
