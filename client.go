package pixels

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	baseURL      = "https://api.pexels.com/v1"
	videoBaseURL = "https://api.pexels.com"
)

type RateLimit struct {
	Limit     int64
	Remaining int64
	Reset     time.Time
}

type Client struct {
	PhotoService      PhotoService
	VideoService      VideoService
	CollectionService CollectionService

	HTTPClient    *http.Client
	AccessToken   string
	BaseURL       string
	LastRateLimit *RateLimit
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
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cli.AccessToken))
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
	fmt.Println(reqURL)

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

	rl, err := RateLimitFromHeader(resp.Header)
	if err != nil {
		return err
	}

	cli.LastRateLimit = rl

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return fmt.Errorf("cannot parse HTTP body: %w", err)
	}

	return nil
}

func RateLimitFromHeader(h http.Header) (*RateLimit, error) {
	ls := h.Get("X-Ratelimit-Limit")
	if ls == "" {
		return nil, errors.New("cannot get X-Ratelimit-Limit from header")
	}

	l, err := strconv.ParseInt(ls, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("X-Ratelimit-Limit is invalid value: %w", err)
	}

	rs := h.Get("X-Ratelimit-Remaining")
	if rs == "" {
		return nil, errors.New("cannot get X-Ratelimit-Remaining from header")
	}

	r, err := strconv.ParseInt(rs, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("X-Rate-Limit-Remaining is invalid value: %w", err)
	}

	ts := h.Get("X-Ratelimit-Reset")
	if ts == "" {
		return nil, errors.New("cannot get X-Ratelimit-Reset from header")
	}

	t, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("X-Ratelimit-Reset is invalid value: %w", err)
	}

	return &RateLimit{
		Limit:     l,
		Remaining: r,
		Reset:     time.Unix(t, 0),
	}, nil
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
