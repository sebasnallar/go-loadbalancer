# go-loadbalancer
Golang LoadBalancer!

First build the loadbalancer =) remember to choose the port you want in the main.go file

go build -o loadbalancer ./cmd/loadbalancer

To run it you should use:

./loadbalancer --backend-urls "http://localhost:8000,http://localhost:8001"

Test this just by starting two basic servers, you can use this structure

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        id := "server-1"
        fmt.Fprintf(w, "Hello from %s", id)
    })

    http.ListenAndServe(":8000", nil)
}
```
