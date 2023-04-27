package loadbalancer

import (
	"net/http"
	"net/http/httputil"
)

type LoadBalancer struct {
	Backends []*Backend
	RR       *RoundRobin
	Proxy    http.HandlerFunc
}

func NewLoadBalancer(backends []*Backend) *LoadBalancer {
	rr := NewRoundRobin(len(backends))

	lb := &LoadBalancer{
		Backends: backends,
		RR:       rr,
	}

	lb.Proxy = func(w http.ResponseWriter, r *http.Request) {
		backend := lb.getNextBackend()
		if backend == nil {
			http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(backend.URL)
		proxy.ServeHTTP(w, r)
	}

	return lb
}

func (lb *LoadBalancer) getNextBackend() *Backend {
	var backend *Backend
	for i := 0; i < len(lb.Backends); i++ {
		index := lb.RR.Next()
		backend = lb.Backends[index]

		backend.Mutex.RLock()
		if backend.HealthCheckStatus {
			backend.Mutex.RUnlock()
			break
		}
		backend.Mutex.RUnlock()
	}

	return backend
}
