package curl

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

var (
	NotValidError = errors.New("curl is not valid")
)

type Curl struct {
	Content
	HttpCommunication
}

type Content struct {
	Url    *url.URL
	Option Option
}

type Option map[string]string

type HttpCommunication interface {
	Request() (*http.Request, error)
	Response() (*http.Response, error)
}

var boolOptions = []string{
	"--anyauth", "--append", "-a", "--basic", "--cert-status", "--compressed-ssh", "--compressed", "--create-dirs",
	"--crlf", "--digest", "--disable-eprt", "--disable-epsv", "--disable", "-q", "--disallow-username-in-url",
	"--doh-cert-status", "--doh-insecure", "--fail-early", "--fail-with-body", "-f", "--fail", "--false-start",
	"--ftp-create-dirs", "--ftp-pasv", "--ftp-pret", "--ftp-skip-pasv-ip", "--ftp-ssl-ccc", "--ftp-ssl-control",
	"-G", "--get", "-g", "--globoff", "--haproxy-protocol", "-I", "--head", "--http0.9", "-0", "--http1.0", "--http1.1",
	"--http2-prior-knowledge", "--http2", "--http3", "--ignore-content-length", "-i", "--include", "-k", "--insecure",
	"-4", "--ipv4", "-6", "--ipv6", "-j", "--junk-session-cookies", "-l", "--list-only", "--location-trusted", "-L",
	"--location", "--mail-rcpt-allowfails", "-M", "--manual", "--metalink", "--negotiate", "--netrc-optional", "-n",
	"--netrc", "-:", "--next", "--no-alpn", "-N", "--no-buffer", "--no-keepalive", "--no-npn", "--no-progress-meter",
	"--no-sessionid", "--ntlm-wb", "--ntlm", "--parallel-immediate", "-Z", "--parallel", "--path-as-is", "--post301",
	"--post302", "--post303", "-#", "--progress-bar", "--proxy-anyauth", "--proxy-basic", "--proxy-digest",
	"--proxy-insecure", "--proxy-negotiate", "--proxy-ntlm", "--proxy-ssl-allow-beast", "--proxy-ssl-auto-client-cert",
	"--proxy-tlsv1", "-p", "--proxytunnel", "--raw", "-J", "--remote-header-name", "--remote-name-all", "-O",
	"--remote-name", "-R", "--remote-time", "--retry-all-errors", "--retry-connrefused", "--sasl-ir", "-S",
	"--show-error", "-s", "--silent", "--socks5-basic", "--socks5-gssapi-nec", "--socks5-gssapi", "--ssl-allow-beast",
	"--ssl-auto-client-cert", "--ssl-no-revoke", "--ssl-reqd", "--ssl-revoke-best-effort", "--ssl", "-2", "--sslv2",
	"-3", "--sslv3", "--styled-output", "--suppress-connect-headers", "--tcp-fastopen", "--tcp-nodelay",
	"--tftp-no-options", "--tlsv1.0", "--tlsv1.1", "--tlsv1.2", "--tlsv1.3", "-1", "--tlsv1", "--tr-encoding",
	"--trace-time", "-B", "--use-ascii", "-v", "--verbose", "-V", "--version", "--xattr",
}

func New(curl string) (*Curl, error) {
	content, err := parseCurl(curl)
	return &Curl{Content: *content}, err
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

func parseCurl(curl string) (*Content, error) {
	curl = strings.TrimLeft(curl, " ")
	curl = strings.TrimRight(curl, " ")
	arr := strings.Split(curl, " ")
	if arr[0] != "curl" {
		return nil, NotValidError
	}
	c := &Content{}
	for i := 1; i < len(arr); i++ {
		if arr[i][0] == '-' {
			if val, ok := c.Option[arr[i]]; ok {
				val = val + "," + arr[i+1]
				i++
				continue
			} else {
				c.Option[arr[i]] = arr[i+1]
				i++
				continue
			}
		} else {
			c.Url, _ = url.Parse(arr[i])
			continue
		}
	}
	return c, nil
}
