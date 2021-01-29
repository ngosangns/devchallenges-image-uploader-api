package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// func connectDB() (*mongo.Client, context.CancelFunc, error) {
// 	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	if err = client.Connect(ctx); err != nil {
// 		return client, cancel, err
// 	}
// 	return client, cancel, err
// }

func main() {
	// Get port
	port := os.Getenv("PORT")
	// port := "80"
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	// Limit size
	e.Use(middleware.BodyLimit("4M"))
	// Disable CROS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodOptions, http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	// Routes
	e.GET("/", Hello)
	e.Static("/static", "images")
	e.POST("/images", Create)

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}
