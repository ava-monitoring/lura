package client

import (
	"context"
	"fmt"
	"net/http"
)

// HTTPRequestExecutor defines the interface of the request executor for the HTTP transport protocol
type HTTPRequestExecutor func(ctx context.Context, req *http.Request) (*http.Response, error)

// DefaultHTTPRequestExecutor creates a HTTPRequestExecutor with the received HTTPClientFactory
func DefaultHTTPRequestExecutor(clientFactory HTTPClientFactory) HTTPRequestExecutor {
	return func(ctx context.Context, req *http.Request) (*http.Response, error) {
		return clientFactory(ctx).Do(req.WithContext(ctx))
	}
}

// HTTPClientFactory creates http clients based with the received context
type HTTPClientFactory func(ctx context.Context) *http.Client

// NewHTTPClient just returns the http default client
func NewHTTPClient(ctx context.Context) *http.Client { return defaultHTTPClient }

func checkRedirect (req *Request, via []*Request) error {
	fmt.Println("DEBUG: checkRedirect")
        return http.ErrUseLastResponse
}

var defaultHTTPClient = &http.Client{CheckRedirect : checkRedirect}
