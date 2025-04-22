package telebot

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

// RequestOption represents a functional option for requests
type RequestOption func(*requestOptions)

type requestOptions struct {
	client *http.Client
	url    string

	toBotApiLatency   prometheus.Observer
	fromBotApiLatency prometheus.Observer
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

func WithFileLatencyMetric(toBotApiLatency prometheus.Observer, fromBotApiLatency prometheus.Observer) RequestOption {
	return func(opts *requestOptions) {
		opts.toBotApiLatency = toBotApiLatency
		opts.fromBotApiLatency = fromBotApiLatency
	}
}
