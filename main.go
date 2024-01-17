package main

import "github.com/saaitt/snappfood_test/app"

func main() {
	orderManager := app.New()
	orderManager.Initialize()
	orderManager.Start()
}
