package pexels

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"
)

const (
	baseURL      = "https://api.pexels.com/v1"
	videoBaseURL = "https://api.pexels.com"
)

type Client struct {
	PhotoService      PhotoService
	VideoService      VideoService
	CollectionService CollectionService

	HTTPClient  *http.Client
	AccessToken string
	BaseURL     string
}

func NewClient(accessToken string) *Client {
	var cli Client
	cli.PhotoService = &photoService{cli: &cli}
	cli.VideoService = &videoService{cli: &cli}
	cli.CollectionService = &collectionService{cli: &cli}

	cli.AccessToken = accessToken
	cli.BaseURL = baseURL

	return &cli
}

func (cli *Client) httpClient() *http.Client {
	if cli.HTTPClient != nil {
		return cli.HTTPClient
	}

	return http.DefaultClient
}

func (cli *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", cli.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("User-Agent", "Pexels/Go")
	return cli.httpClient().Do(req)
}

func (cli *Client) get(ctx context.Context, path string, params url.Values, v interface{}) error {
	r := regexp.MustCompile("^videos")
	if r.MatchString(path) {
		cli.BaseURL = videoBaseURL
	}

	reqURL := cli.BaseURL + "/" + path
	if len(params) > 0 {
		reqURL += "?" + params.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return fmt.Errorf("cannot create HTTP request: %w", err)
	}

	resp, err := cli.do(ctx, req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if !(resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices) {
		return cli.error(resp.StatusCode, resp.Body)
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return fmt.Errorf("cannot parse HTTP body: %w", err)
	}

	return nil
}

func (cli *Client) error(statusCode int, body io.Reader) error {
	var aerr APIError
	if err := json.NewDecoder(body).Decode(&aerr); err != nil {
		return &APIError{HTTPStatus: statusCode}
	}
	aerr.HTTPStatus = statusCode
	return &aerr
}

func StructToMap(i interface{}) (values url.Values) {
	values = url.Values{}
	if reflect.ValueOf(i).IsNil() {
		return
	}
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		if !iVal.Field(i).IsZero() {
			values.Set(toSnakeCase(typ.Field(i).Name), fmt.Sprint(iVal.Field(i)))
		}
	}
	return
}

func toSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
