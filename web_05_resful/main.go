package main

import (
	"net/http"

	"github.com/ecosy/webGo/web_05/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
