package curl

import "net/http"

type Curl struct {
	Content string
	HttpCommunication
}

type HttpCommunication interface {
	Request() *http.Request
	Response() *http.Response
}

func New(curl string) *Curl {
	return &Curl{
		Content: curl,
	}
}

func (c *Curl) Request() *http.Request {
	return &http.Request{}
}

func (c *Curl) Response() *http.Response {
	return &http.Response{}
}
