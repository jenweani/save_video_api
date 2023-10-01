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

	handlers.VideoDataMap = map[string][]byte{}
	g := gin.Default()

	g.LoadHTMLFiles("index.html")

	g.GET("/health", handlers.HealthHandler)
	
	api := g.Group("/api")
	// view video data endpoints
	api.GET("/video/:videoID", handlers.ViewVideoHandler)
	api.GET("/transcript/:videoID", handlers.ViewTranscriptHandler)


	// stream video endpoints
	api.GET("/startStream", handlers.StartStream)
	api.POST("/streamupload/:videoID", handlers.StreamUpload)
	api.POST("/endstream/:videoID", handlers.StopStream)


	// handle 404 routes
	g.NoRoute(handlers.NoRouteHandler)

	http.ListenAndServe(":8080", g)
}