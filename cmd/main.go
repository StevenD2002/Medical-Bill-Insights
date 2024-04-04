package main

import (
	"context"
	"net/http"

	"github.com/StevenD2002/Medical-Bill-Insights/web/template/partials"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		component := partials.Hello("Steven")
		component.Render(context.Background(), w)
	})
	http.ListenAndServe(":8080", nil)
}
