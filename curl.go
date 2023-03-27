package curl

import (
	"net/http"
	"net/url"
)

type Curl struct {
	Content
	HttpCommunication
}

type Content struct {
	Url     *url.URL
	Options []Option
}

type Option map[string]string

type HttpCommunication interface {
	Request() (*http.Request, error)
	Response() (*http.Response, error)
}

func New(curl string) *Curl {
	return &Curl{
		Content: *generateContentFromCurl(curl),
	}
}

func (c *Curl) Request() (*http.Request, error) {
	return &http.Request{}, nil
}

func (c *Curl) Response() (*http.Response, error) {
	request, err := c.Request()
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(request)
}

func generateContentFromCurl(curl string) *Content {
	return nil
}
