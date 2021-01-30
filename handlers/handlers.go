package handlers

import (
	"io"
	"math/rand"
	"net/http"
	"strings"

	"image-uploader/libs"
	"image-uploader/models"

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
		libs.LogError(err)
		return c.JSON(http.StatusNotImplemented, models.Response{
			Message: libs.ErrorUploadingResponse(),
		})
	}

	// -----

	// Source
	file, err := fileHeader.Open()
	if err != nil {
		libs.LogError(err)
		return c.JSON(http.StatusNotImplemented, models.Response{
			Message: libs.ErrorUploadingResponse(),
		})
	}
	// Checking file MIME header
	buff := make([]byte, 512) // Docs tell that it take only first 512 bytes into consideration
	if _, err := file.Read(buff); err != nil {
		libs.LogError(err)
		return c.JSON(http.StatusNotImplemented, models.Response{
			Message: libs.ErrorUploadingResponse(),
		})
	}
	mimeType := http.DetectContentType(buff)
	if "image" != strings.Split(mimeType, "/")[0] {
		return c.JSON(http.StatusUnsupportedMediaType, models.Response{
			Message: "File should be image",
		})
	}
	// Close file
	defer file.Close()

	// -----

	// Create new random name
	newFileName := ""
	for i := 0; i < 48; i++ {
		newFileName += string(rune(rand.Intn(26) + 65))
	}

	// -----

	// Copy temp file to server
	// Source
	src, err := fileHeader.Open()
	if err != nil {
		libs.LogError(err)
		return c.JSON(http.StatusNotImplemented, models.Response{
			Message: libs.ErrorUploadingResponse(),
		})
	}
	// Destination
	dst, err := libs.CreateFile("./images/" + newFileName)
	if err != nil {
		libs.LogError(err)
		return c.JSON(http.StatusNotImplemented, models.Response{
			Message: libs.ErrorUploadingResponse(),
		})
	}
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		libs.LogError(err)
		return c.JSON(http.StatusNotImplemented, models.Response{
			Message: libs.ErrorUploadingResponse(),
		})
	}
	// Close file
	defer src.Close()
	defer dst.Close()

	// Return response
	res := map[string]interface{}{
		"link": "http://" + c.Request().Host + "/static/" + newFileName,
	}
	return c.JSON(http.StatusOK, res)
}
