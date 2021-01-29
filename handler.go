package main

import (
	"io"
	"math/rand"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// Handler
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello! This is API of project: Dev challenges - Image uploader.")
}

func Create(c echo.Context) error {
	// Get form file value
	fileHeader, err := c.FormFile("image")
	if err != nil {
		LogError(err)
		return c.JSON(http.StatusNotImplemented, Response{ErrorUploadingResponse()})
	}

	// -----

	// Source
	file, err := fileHeader.Open()
	if err != nil {
		LogError(err)
		return c.JSON(http.StatusNotImplemented, Response{ErrorUploadingResponse()})
	}
	// Checking file MIME header
	buff := make([]byte, 512) // Docs tell that it take only first 512 bytes into consideration
	if _, err := file.Read(buff); err != nil {
		LogError(err)
		return c.JSON(http.StatusNotImplemented, Response{ErrorUploadingResponse()})
	}
	mimeType := http.DetectContentType(buff)
	if "image" != strings.Split(mimeType, "/")[0] {
		return c.JSON(http.StatusUnsupportedMediaType, Response{"File should be image"})
	}
	// Close file
	defer file.Close()

	// -----

	// Create new random name
	newFileName := ""
	for i := 0; i < 48; i++ {
		newFileName += string(rand.Intn(26) + 65)
	}

	// -----

	// Copy temp file to server
	// Source
	src, err := fileHeader.Open()
	if err != nil {
		LogError(err)
		return c.JSON(http.StatusNotImplemented, Response{ErrorUploadingResponse()})
	}
	// Destination
	dst, err := CreateFile("./images/" + newFileName)
	if err != nil {
		LogError(err)
		return c.JSON(http.StatusNotImplemented, Response{ErrorUploadingResponse()})
	}
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		LogError(err)
		return c.JSON(http.StatusNotImplemented, Response{ErrorUploadingResponse()})
	}
	// Close file
	defer src.Close()
	defer dst.Close()

	// -----

	// // Connect db
	// client, cancelDB, err := connectDB()
	// if err != nil {
	// 	logError(err)
	// 	return err
	// }
	// // Get collection
	// imagesCollection := client.Database("image-uploader").Collection("images")
	// // Add file info to db
	// ash := Image{newFileName}
	// insertResult, err := imagesCollection.InsertOne(context.TODO(), ash)
	// if err != nil {
	// 	logError(err)
	// 	return c.JSON(http.StatusNoContent, "{err: \"Error while uploading...\"}")
	// }
	// // Disconnect database
	// defer cancelDB()

	// -----

	// Return response
	res := Image{newFileName}
	return c.JSON(http.StatusOK, res)
}
