![Go](https://github.com/shitpostingio/alldebrid/workflows/Go/badge.svg)
[![DeepSource](https://static.deepsource.io/deepsource-badge-light-mini.svg)](https://deepsource.io/gh/shitpostingio/alldebrid/?ref=repository-badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/shitpostingio/alldebrid)](https://goreportcard.com/report/github.com/shitpostingio/alldebrid)
# Alldebrid Go Library

This is a Go wrapper for Alldebrid v4 APIs

# Installation

`go get -u github.com/shitpostingio/alldebrid`

# Usage
To get started import the library in your program and initialize a client

```go
import	"github.com/shitpostingio/alldebrid"

func main() {
	client, _ := alldebrid.New("friji3r124rnwer", "MyAppName")

	m := []string{"magnet:?xt=urn:btih:f95c371d5609d15f6615139be84edbb5b94a79bc"}

	resp, err := client.UploadMagnet(m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
```
