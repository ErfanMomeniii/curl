package curl_test

import (
	"github.com/erfanmomeniii/curl"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_Request(t *testing.T) {
	// Test request single header
	c, err := curl.New("curl -H \"Test:yes\" www.google.com")

	req := c.Request()

	assert.NoError(t, err)

	assert.Equal(t, "www.google.com", req.URL.String())
	assert.Equal(
		t, http.Header{"Test": {"yes"}, "Content-Type": {"application/x-www-form-urlencoded"}}, req.Header,
	)

	// Test request with multiple header
	c, err = curl.New("curl -H \"Test1:yes\" -H \"Test2:yes\" www.google.com")

	req = c.Request()

	assert.NoError(t, err)

	assert.Equal(t, "www.google.com", req.URL.String())
	assert.Equal(
		t, http.Header{"Test1": {"yes"}, "Test2": {"yes"}, "Content-Type": {"application/x-www-form-urlencoded"}}, req.Header,
	)

	// Test request with body and header
	c, err = curl.New("curl -H \"Test1:yes\" -H \"Test2:yes\" -d \"user=foobar\" www.google.com")

	req = c.Request()

	body, _ := ioutil.ReadAll(req.Body)

	assert.NoError(t, err)

	assert.Equal(t, "www.google.com", req.URL.String())
	assert.Equal(t, "\"user=foobar\"", string(body))
	assert.Equal(
		t, http.Header{"Test1": {"yes"}, "Test2": {"yes"}, "Content-Type": {"application/x-www-form-urlencoded"}}, req.Header,
	)

	// Test request with body and header and set post method
	c, err = curl.New("curl -X POST -H \"Test1:yes\" -H \"Test2:yes\" -d \"user=foobar\" www.google.com")

	req = c.Request()

	body, _ = ioutil.ReadAll(req.Body)

	assert.NoError(t, err)

	assert.Equal(t, "www.google.com", req.URL.String())
	assert.Equal(t, "\"user=foobar\"", string(body))
	assert.Equal(
		t, http.Header{"Test1": {"yes"}, "Test2": {"yes"}, "Content-Type": {"application/x-www-form-urlencoded"}}, req.Header,
	)
	assert.Equal(t, req.Method, "POST")
}
