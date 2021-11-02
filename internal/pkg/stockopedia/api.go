package stockopedia

import (
	"encoding/json"
	"fmt"
	"github.com/google/martian/log"
	"golang.org/x/net/publicsuffix"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

const (
	LoginPostUri      = "auth/login/"
	ByToken           = "api/v1/users/bytoken"
	SearchSecurityUri = "api/v1/search/all?query=%s&type=security&isListed=1"
)

type ApiClientConfig struct {
	Client *http.Client
}

type ApiClient struct {
	http *http.Client
}

func NewApiClient(config func() ApiClientConfig) (*ApiClient, error) {
	c := config()

	if c.Client == nil {
		jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})

		if err != nil {
			return nil, err
		}

		c.Client = &http.Client{
			Timeout: time.Second * 10000,
			Jar:     jar,
		}
	}

	client := ApiClient{
		http: c.Client,
	}

	return &client, nil
}

func (c *ApiClient) SearchSecurity(q string) (*SearchResults, error) {
	if c.http.Jar == nil {
		return nil, fmt.Errorf("cookie jar is nil on http client")
	}

	u := fmt.Sprintf("%s/%s", "https://app.stockopedia.com", fmt.Sprintf(SearchSecurityUri, q))

	log.Debugf("fetching %s", u)

	resp, err := c.http.Get(u)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status code, expecte 200, got %d", resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		return nil, fmt.Errorf("invalid content-type returned, got %s", resp.Header.Get("Content-Type"))
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := SearchResults{}

	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (c *ApiClient) Auth(u string, p string) error {
	if c.http.Jar == nil {
		return fmt.Errorf("cookie jar is nil on http client")
	}

	data := url.Values{}
	data.Set("username", u)
	data.Set("password", p)

	postForm, err := c.http.PostForm(
		fmt.Sprintf("https://www.stockopedia.com/%s", LoginPostUri),
		data,
	)

	if err != nil {
		return err
	}

	if postForm.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to authenicate, non-200 response, %d given", postForm.StatusCode)
	}

	return nil
}
