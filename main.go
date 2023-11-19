package main

import (
	"fmt"
	"htmxgo/routes"
	"htmxgo/settings"
	"net/http"
)

func main() {
	PORT := fmt.Sprintf(":%d", settings.PORT)

	routing := http.HandlerFunc(routes.ServeRoute)

	fmt.Printf("Running server http://localhost:%d/", settings.PORT)
	http.ListenAndServe(PORT, routing)
}
