package curl

import "net/http"

type Curl struct {
	Content string
	HttpCommunication
}

type HttpCommunication interface {
	Request() (*http.Request, error)
	Response() (*http.Response, error)
}

func New(curl string) *Curl {
	return &Curl{
		Content: curl,
	}
}

func (c *Curl) Request() (*http.Request, error) {
	return &http.Request{}, nil
}

func (c *Curl) Response() (*http.Response, error) {
	return &http.Response{}, nil
}
