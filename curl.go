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
	"--trace-time", "-B", "--use-ascii", "-v", "--verbose", "-V", "--version", "--xattr", "--no-anyauth", "--no-basic",
	"--no-cert-status", "--no-compressed-ssh", "--no-compressed", "--no-create-dirs", "--no-crlf", "--no-digest",
	"--no-disable-eprt", "--no-disable-epsv", "--no-disable", "--no-disallow-username-in-url", "--no-doh-cert-status",
	"--no-doh-insecure", "no-fail-early", "--no-fail-with-body", "--no-fail", "--no-false-start",
	"--no-ftp-create-dirs", "--no-ftp-pasv", "--no-ftp-pret", "--no-ftp-skip-pasv-ip", "--no-ftp-ssl-ccc",
	"--no-ftp-ssl-control", "--no-get", "--no-globoff", "--no-haproxy-protocol", "--no-head", "--no-http0.9", "--no-http1.0",
	"--no-http1.1", "--no-http2-prior-knowledge", "--no-http2", "--no-http3", "--no-ignore-content-length", "--no-include",
	"--no-insecure", "--no-ipv4", "--no-ipv6", "--no-junk-session-cookies", "--no-list-only", "--no-location-trusted",
	"--no-location", "--no-mail-rcpt-allowfails", "--no-manual", "--no-metalink", "--no-negotiate", "--no-netrc-optional",
	"--no-netrc", "--no-next", "--no-ntlm-wb", "--no-ntlm", "--no-parallel-immediate", "--no-parallel", "--no-path-as-is",
	"--no-post301", "--no-post302", "--no-post303", "--no-progress-bar", "--no-proxy-anyauth", "--no-proxy-basic",
	"--no-proxy-digest", "--no-proxy-insecure", "--no-proxy-negotiate", "--no-proxy-ntlm", "--no-proxy-ssl-allow-beast",
	"--no-proxy-ssl-auto-client-cert", "--no-proxy-tlsv1", "--no-proxytunnel", "--no-raw", "--no-remote-header-name",
	"--no-remote-name-all", "--no-remote-name", "--no-remote-time", "--no-retry-all-errors", "--no-retry-connrefused",
	"--no-sasl-ir", "--no-show-error", "--no-silent", "--no-socks5-basic", "--no-socks5-gssapi-nec", "--no-socks5-gssapi",
	"--no-ssl-allow-beast", "--no-ssl-auto-client-cert", "--no-ssl-no-revoke", "--no-ssl-reqd", "--no-ssl-revoke-best-effort",
	"--no-ssl", "--no-sslv2", "--no-sslv3", "--no-styled-output", "--no-suppress-connect-headers", "--no-tcp-fastopen",
	"--no-tcp-nodelay", "--no-tftp-no-options", "--no-tlsv1.0", "--no-tlsv1.1", "--no-tlsv1.2", "--no-tlsv1.3",
	"--no-tlsv1", "--no-tr-encoding", "--no-trace-time", "--no-use-ascii", "--no-verbose", "--no-version", "--no-xattr",
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
	var st Stack
	for i := 1; i < len(arr); i++ {
		if arr[i][0] == '-' {
			if !arrayExist(boolOptions, arr[i]) {
				st.Push(arr[i])
			}
		} else {
			if len(st) == 0 {
				c.Url, _ = url.Parse(arr[i])
			} else {
				if val, ok := c.Option[st[0]]; ok {
					val = val + "," + arr[i]
					c.Option[st[0]] = val
				} else {
					c.Option[st[0]] = arr[i]
				}

				st.Pop()
			}
		}
	}
	return c, nil
}
