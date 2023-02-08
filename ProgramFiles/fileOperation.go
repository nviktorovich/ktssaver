package programfiles

import (
	"log"
	"path/filepath"
)

func GetAbsPath() (string, error) {
	path, err := filepath.Abs(".")
	if err != nil {
		log.Print(err)
		return "", err
	}
	return path, nil
}
