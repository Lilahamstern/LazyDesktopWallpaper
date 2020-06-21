package main

import (
	"github.com/joho/godotenv"
	"github.com/lilahamstern/automaticDesktopWallpaper/pkg/image"
	"github.com/lilahamstern/automaticDesktopWallpaper/pkg/win"
	"log"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file ", err)
	}

	for {
		image.Download()
		win.ChangeDesktopBackground()
		time.Sleep(20 * time.Minute)
	}
}
