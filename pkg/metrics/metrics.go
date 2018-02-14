package metrics

import (
	prom "github.com/prometheus/client_golang/prometheus"
)

func New() *prom.Registry {
	return prom.NewRegistry()
}

func Register(reg *prom.Registry) error {
	// TODO: random metric for now.
	httpRequestsTotal := prom.NewCounterVec(prom.CounterOpts{
		Name: "http_requests_total",
		Help: "Count of all HTTP requests",
	}, []string{"code", "method"})

	reg.MustRegister(httpRequestsTotal)

	return nil
}
