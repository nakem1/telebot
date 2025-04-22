package telebot

import "net/http"

// RequestOption represents a functional option for requests
type RequestOption func(*requestOptions)

type requestOptions struct {
	client *http.Client
	url    string
}

// WithCustomHTTPClient allows to use custom HTTP client for requests
func WithCustomHTTPClient(client *http.Client) RequestOption {
	return func(opts *requestOptions) {
		opts.client = client
	}
}

func WithCustomURL(url string) RequestOption {
	return func(opts *requestOptions) {
		opts.url = url
	}
}
