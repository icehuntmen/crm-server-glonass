package middlewares

import (
	"log"
	"os"
)

var Version string

func GetVersion() (string, error) {
	file, err := os.Open("./VERSION")
	if err != nil {
		log.Fatal("Error opening file:", err)
		return "", err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Fatal("Error getting file information:", err)
		return "", err
	}

	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	if err != nil {
		log.Fatal("Error reading file:", err)
		return "", err
	}

	return string(data), nil
}
