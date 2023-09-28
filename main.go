package main

import (
	"jonnedu/hng_task5/handlers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Unable to load env file")
	}

	g := gin.Default()

	g.LoadHTMLFiles("index.html")

	api := g.Group("/api")
	api.GET("/video/:filename", handlers.ViewVideoHandler)
	api.GET("/video/page/:filename", handlers.VideoPageHandler)
	api.POST("/upload", handlers.UploadVideoHandler)

	// handle 404 routes
	g.NoRoute(handlers.NoRouteHandler)

	http.ListenAndServe(":8080", g)
}