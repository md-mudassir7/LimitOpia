package handlers

import (
	"fmt"
	"net/http"
)

func ExampleRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Request successful!")
}
