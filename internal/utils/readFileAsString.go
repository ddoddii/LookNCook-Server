package utils

import (
	"io/ioutil"
)

func ReadFileAsString(filepath string) (string, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "Error Reading file as string", err
	}
	return string(data), nil
}
