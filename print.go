package httpprinter

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func WrapClient(client *http.Client) *http.Client {
	client.Transport = WrapTransport(client.Transport)
	return client
}

func WrapTransport(tspt http.RoundTripper) http.RoundTripper {
	return &printableTransport{wrapped: tspt}
}

type printableTransport struct {
	wrapped http.RoundTripper
}

func (pt *printableTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if bs, err := httputil.DumpRequestOut(req, true); err == nil {
		fmt.Printf("%s\n", bs)
	}

	resp, err := pt.wrapped.RoundTrip(req)
	if err == nil {
		if bs, err := httputil.DumpResponse(resp, true); err == nil {
			fmt.Printf("%s\n", bs)
		}
	}
	return resp, err
}
