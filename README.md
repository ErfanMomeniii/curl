<p align="center">
<img src="./assets/photo/logo.png" width=50% height=50%>
</p>
<p align="center">
<a href="https://pkg.go.dev/github.com/mehditeymorian/koi/v3?tab=doc"target="_blank">
    <img src="https://img.shields.io/badge/Go-1.19+-00ADD8?style=for-the-badge&logo=go" alt="go version" />
</a>

<img src="https://img.shields.io/badge/license-MIT-magenta?style=for-the-badge&logo=none" alt="license" />
<img src="https://img.shields.io/badge/Version-1.0.0-red?style=for-the-badge&logo=none" alt="version" />
</p>

# curl
curl is a lightweight package for rendering http request and response from curl string with Go


# Documentation

## Install

```bash
go get github.com/erfanmomeniii/curl
```   

Next, include it in your application:

```bash
import "github.com/erfanmomeniii/curl"
``` 

## Quick Start

The following example demonstrates how to use this package for generating request and response from curl:

```go
package main

import (
	"fmt"
	"github.com/erfanmomeniii/curl"
)

func main() {
	c, _ := curl.New("curl -H \"Test1:no\" -H \"Test2:yes\" -d \"User=foobar\" www.google.com")

	request := c.Request()
	response, _ := c.Response()

	fmt.Println(request, response)
}

```
## Contributing
Pull requests are welcome. For changes, please open an issue first to discuss what you would like to change.
Please make sure to update tests as appropriate.
