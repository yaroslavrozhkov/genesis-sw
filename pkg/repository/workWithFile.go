package repository

import (
	models "GenesisProject/models"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func writeFile(filePath, content string) error {
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return err
	}
	fmt.Println("File written successfully.")
	return nil
}

func readFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func checkAndWrite(content string) (string, error) {
	f, err := os.OpenFile("storage.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	f.Close()

	readContent, err := readFile("storage.txt")
	if err != nil {
		log.Fatal(err)
	}
	splits := strings.Split(readContent, ",")

	for _, item := range splits {
		if content == strings.Trim(item, "\n") {
			//return "Такий E-mail вже було додано", errors.New("Email already exists")
			return "Такий E-mail вже було додано",
				&models.RequestError{
					StatusCode: 409,
					Err:        errors.New("unavailable"),
				}
		}
	}

	newEmail := content

	if len(readContent) > 1 {
		newEmail = ",\n" + content
	}

	newContent := readContent + newEmail

	err = writeFile("storage.txt", newContent)
	if err != nil {
		log.Fatal(err)
	}

	return "E-mail додано", nil
}

func getEmailsList() []string {
	readContent, err := readFile("storage.txt")
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(readContent, ",")
}
