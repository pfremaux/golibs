package main

import (
	"fmt"
	"net/http"

	"github.com/pfremaux/golibs/web/pkg/web"
)

func main() {
	config := web.NewWebServer("", 8080)
	config.RegisterPublicEndpoint("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"message": "Hello World"}`)
	})
	config.RegisterPublicEndpoint("/hi", web.JsonResponse(`{"message": "Hi world!"}`))

	config.Listen()
}
