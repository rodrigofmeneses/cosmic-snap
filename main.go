package main

import (
	"github.com/rodrigofmeneses/cosmic-snap/config"
	"github.com/rodrigofmeneses/cosmic-snap/router"
)

var logger *config.Logger

func main() {
	logger = config.GetLogger("main")
	// Initialize Config
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
