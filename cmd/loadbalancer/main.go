package main

import (
    "log"
    "net/http"
    "net/url"
    "time"

    "github.com/sebasnallar/go-loadbalancer/pkg/loadbalancer"
    "github.com/spf13/pflag"
)

func main() {
    var listenAddr string
    var backendURLs []string
    var healthCheckPath string
    var healthCheckInterval time.Duration

    pflag.StringVar(&listenAddr, "listen-addr", ":8080", "Load balancer listen address")
    pflag.StringSliceVar(&backendURLs, "backend-urls", nil, "Comma-separated list of backend URLs")
    pflag.StringVar(&healthCheckPath, "healthcheck-path", "/health", "Path for backend health check")
    pflag.DurationVar(&healthCheckInterval, "healthcheck-interval", 10*time.Second, "Interval between health checks")
    pflag.Parse()

    backends := []*loadbalancer.Backend{}

    for _, backendURL := range backendURLs {
        parsed, err := url.Parse(backendURL)
        if err != nil {
            log.Fatalf("Error parsing backend URL: %v", err)
        }
        backend := loadbalancer.NewBackend(parsed, healthCheckPath, healthCheckInterval)
        backends = append(backends, backend)
    }

    lb := loadbalancer.NewLoadBalancer(backends)

    log.Printf("Starting load balancer on %s", listenAddr)

    http.HandleFunc("/", lb.Proxy)
    log.Fatal(http.ListenAndServe(listenAddr, nil))
}
