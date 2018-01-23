package server

import (
	"net/http"
	"github.com/cnosuke/go_blank_image/blank_image"
)

func PngHandler(w http.ResponseWriter, r *http.Request) {
	b := blank_image.PngBytes()
	w.Write(b)
}

func GifHandler(w http.ResponseWriter, r *http.Request) {
	b := blank_image.GifBytes()
	w.Write(b)

}