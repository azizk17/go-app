package main

import (
	"fmt"

	"github.com/techschool/myApp/config"
)

func main() {

	c, err := config.LoadConfig(".")
	if err != nil {
		fmt.Errorf("Could not load configs", err)
	}
	fmt.Printf("VAR %v\n", c.DB.Host)
	if err := c.DB.Setup(); err != nil {
		fmt.Errorf("Could not init database", err)
	}
	c.Server.Setup()
	if err := c.Server.ServeWithGraceFullShutdown(); err != nil {
		fmt.Errorf("Could not run server", err)
	}
	// logger, _ := zap.NewProduction()
	// defer logger.Sync() // flushes buffer, if any
	// sugar := logger.Sugar()
	// sugar.Infow("failed to fetch URL",
	// 	// Structured context as loosely typed key-value pairs.
	// 	"url", url,
	// 	"attempt", 3,
	// 	"backoff", time.Second,
	// )
	// sugar.Infof("Failed to fetch URL: %s", url)
	// fmt.Printf("VAR %v\n", c)

	// c.Setup()
	// c.Server.ServeWithGraceFullShutdown()s
}
