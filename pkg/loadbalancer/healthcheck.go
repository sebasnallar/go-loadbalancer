package loadbalancer

import (
	"net/http"
	"time"
)

func StartHealthCheck(b *Backend, interval time.Duration) {
	ticker := time.NewTicker(interval)

	for {
		select {
		case <-ticker.C:
			res, err := http.Get(b.URL.String() + b.HealthCheckPath)

			b.Mutex.Lock()
			if err != nil || res.StatusCode != http.StatusOK {
				b.HealthCheckStatus = false
			} else {
				b.HealthCheckStatus = true
			}
			b.Mutex.Unlock()
		}
	}
}
