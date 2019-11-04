package repo

import (
	"bufio"
	"errors"
	"log"
	"os"
)

type Message struct {
	Content string `json:"Message"`
}

func PostMessage(message Message) (interface{}, error) {
	if message.Content == "" {
		return nil, errors.New("Message can not be blank")
	}

	messageFile(message.Content)
	return message, nil
}

func AllMessage() interface{} {
	message := openMessageFile()
	return message
}

func messageFile(message string) {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("message.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(message + "\n")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func openMessageFile() interface{} {
	file, err := os.Open("message.log")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()
	return txtlines
}
