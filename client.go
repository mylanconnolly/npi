package npi

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"go.uber.org/zap"
)

// ErrNotFound is used when the record simply wasn't found. This does not
// indicate that there was a problem with the request, simply that the record
// does not exist, or the criteria was incorrect.
var ErrNotFound = errors.New("record not found")

// The different possible log levels.
const (
	LogLevelDevel int = 1 << iota
	LogLevelProd
)

const uri = "https://npiregistry.cms.hhs.gov/api"

// Client is the client used to interact with the NPI API.
type Client struct {
	log *zap.Logger
}

// NewClient is used to create and prepare a new NPI client for use.
func NewClient(logLevel int) (*Client, error) {
	var (
		log *zap.Logger
		err error
	)
	switch logLevel {
	case LogLevelDevel:
		log, err = zap.NewDevelopment()
	case LogLevelProd:
		log, err = zap.NewProduction()
	default:
		return nil, errors.New("unknown log level")
	}
	if err != nil {
		return nil, err
	}
	return &Client{log: log}, nil
}

// Get is used to retrieve an NPI result based on the NPI number provided.
func (c Client) Get(id int) (Result, error) {
	url := buildURL(url.Values{"number": []string{strconv.Itoa(id)}})
	client := newHTTPClient()
	req, err := newHTTPReq("GET", url)

	if err != nil {
		return Result{}, err
	}
	res := listResult{}
	resp, err := client.Do(req)

	if err != nil {
		return Result{}, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return Result{}, err
	}
	if res.ResultCount == 1 {
		return res.Results[0], nil
	}
	return Result{}, ErrNotFound
}

// SearchName is used to find NPI results with the given names. Note that the
// criteria here is fuzzy (by design on the NPI API), so there's no need for
// asterisks or percent signs.
func (c Client) SearchName(firstName, lastName string) ([]Result, error) {
	url := buildURL(url.Values{"first_name": []string{firstName}, "last_name": []string{lastName}})
	client := newHTTPClient()
	req, err := newHTTPReq("GET", url)

	if err != nil {
		return nil, err
	}
	res := listResult{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return res.Results, nil
}

func buildURL(query url.Values) string {
	return uri + "?" + query.Encode()
}

func newHTTPReq(method, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "github.com/mylanconnolly/npi")

	return req, nil
}

func newHTTPClient() http.Client {
	return http.Client{Timeout: 2 * time.Second}
}
