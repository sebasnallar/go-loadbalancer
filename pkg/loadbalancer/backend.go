package loadbalancer

import (
	"net/url"
	"sync"
    "time"

)

type Backend struct {
	URL               *url.URL
	HealthCheckPath   string
	HealthCheckStatus bool
	Mutex             *sync.RWMutex
}

func NewBackend(url *url.URL, healthCheckPath string, healthCheckInterval time.Duration) *Backend {
	backend := &Backend{
		URL:               url,
		HealthCheckPath:   healthCheckPath,
		HealthCheckStatus: false,
		Mutex:             &sync.RWMutex{},
	}

	go StartHealthCheck(backend, healthCheckInterval)

	return backend
}
