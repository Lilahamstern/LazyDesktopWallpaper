package image

import (
	"github.com/lilahamstern/automaticDesktopWallpaper/internal/unsplash"
	"io"
	"net/http"
	"os"
)

var (
	fileName = "wallpaper.jpg"
	photoUrl = ""
)

func Download() {
	photoUrl = unsplash.GetRandomImageUrl()
	file := createFile()

	saveImage(file, httpClient())
}

func saveImage(file *os.File, client *http.Client) {
	resp, err := client.Get(photoUrl)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)

	defer file.Close()

	if err != nil {
		panic(err)
	}
}

func httpClient() *http.Client {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	return &client
}

func createFile() *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	return file
}
