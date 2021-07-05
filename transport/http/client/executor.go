/* Package client provides some http helpers to create http clients and executors
 */
// SPDX-License-Identifier: Apache-2.0
package client

import (
	"context"
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
func NewHTTPClient(ctx context.Context) *http.Client {
	return defaultHTTPClient
}

func checkRedirect (req *http.Request, via []*http.Request) error {
        var redirects = via[0].Header.Get("X-Krakend-Follow-Redirects")

        if redirects != "" && redirects == "false" {
                return http.ErrUseLastResponse
        } else {
               cookie := req.Response.Header.Get("Set-Cookie")
               if len(cookie) > 0 {
                       req.Header.Add("Cookie", cookie)
               }
               return nil
	}
}

var defaultHTTPClient = &http.Client{CheckRedirect : checkRedirect}
