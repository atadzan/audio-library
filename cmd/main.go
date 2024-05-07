package main

import (
	"flag"
	"log"

	"github.com/atadzan/audio-library/internal/app"
)

var configPath *string

func init() {
	// app configuration path flag
	configPath = flag.String("config", "./configs/debug.yaml", "Default config path")
	flag.Parse()
}

// @title Audio library API docs
// @version 1.0
// @description Simple audio library
// @host localhost:8080
// @BasePath /
func main() {
	if err := app.Init(*configPath); err != nil {
		log.Println(err)
		return
	}
}
