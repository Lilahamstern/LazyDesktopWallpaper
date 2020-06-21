package file

import (
	"fmt"
	"log"
	"path/filepath"
)

func GetFilePath() string {
	dir, err := filepath.Abs(filepath.Dir("%appdata%"))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func GetImageFilePath() string {
	return fmt.Sprintf("%s/%s", GetFilePath(), "wallpaper.jpg")
}
