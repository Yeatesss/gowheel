package gowheel

import (
	"image"
	"log"
	"net/http"
)

func GetImageSizeFromUrl(url string) (height,width float64) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	m, _, err := image.Decode(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	g := m.Bounds()
	// Get height and width
	height = float64(g.Dy())
	width = float64(g.Dx())
	return
}