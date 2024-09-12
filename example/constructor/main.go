package constructor

import (
	"fmt"
	"strings"

	"github.com/robert-macwha/safe/pkg/safe"
)

type ConnType uint
type Client struct {
	url      string
	connType ConnType
}

const (
	HTTP ConnType = iota
	HTTPS
)

func SafeNewClient(url string) (res safe.Result[Client]) {
	safe.Handle(&res)

	if url == "" {
		return safe.Err[Client](fmt.Errorf("missing url"))
	}

	if !strings.Contains(url, ":") {
		return safe.Err[Client](fmt.Errorf("missing protocol"))
	}

	switch strings.Split(url, ":")[0] {
	case "http":
		return safe.Ok(Client{url: url, connType: HTTP})
	case "https":
		return safe.Ok(Client{url: url, connType: HTTPS})
	default:
		return safe.Err[Client](fmt.Errorf("unsupported protocol"))
	}
}

func NewClient(url string) (*Client, error) {
	if url == "" {
		return nil, fmt.Errorf("missing url")
	}

	if !strings.Contains(url, ":") {
		return nil, fmt.Errorf("missing protocol")
	}

	switch strings.Split(url, ":")[0] {
	case "http":
		return &Client{url: url, connType: HTTP}, nil
	case "https":
		return &Client{url: url, connType: HTTPS}, nil
	default:
		return nil, fmt.Errorf("unsupported protocol")
	}
}
