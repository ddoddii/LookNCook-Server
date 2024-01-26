package utils

import "io"

import (
	"github.com/labstack/echo/v4"
)

func ExtractImageData(c echo.Context, formKey string) ([]byte, error) {
	file, err := c.FormFile(formKey)
	if err != nil {
		return nil, err
	}

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	return io.ReadAll(src)
}
