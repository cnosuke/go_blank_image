package main

import (
	"net/http"

	"github.com/cnosuke/go_blank_image/server"
)

func main() {
	http.HandleFunc("/blank.png", server.PngHandler)
	http.HandleFunc("/blank.gif", server.GifHandler)

	http.ListenAndServe("127.0.0.1:8080", nil)
}
