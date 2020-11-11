package main

import (
	"github.com/gin-contrib/gzip"

	"go-swag-demo-api/pkg/routing"
)

// @title Demo API
// @version 1.0
// @description This is a demo API
// @host https://api.demo.com
// @BasePath /v1
func main() {
	router := routing.NewRouter()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Run(":" + "3001")
}
