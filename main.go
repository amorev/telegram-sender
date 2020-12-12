package main

import (
	"bytes"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Panic(envErr)
	}
	botToken := ""
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) > 2 {
		botToken = argsWithoutProg[2]
	}
	botToken = os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Panic("TELEGRAM_BOT_TOKEN env is empty")
	}

	filePath := argsWithoutProg[0]
	chatId := ""
	if len(argsWithoutProg) > 1 {
		chatId = argsWithoutProg[1]
	}
	if chatId == "" {
		chatId = os.Getenv("TELEGRAM_CHAT_ID")
	}
	if chatId == "" {
		log.Panic("chatId is empty")
	}
	if filePath == "" {
		log.Panic("filePath is empty")
	}
	sendFile(botToken, chatId, filePath)
}

func sendFile(botToken string, chatId string, filePath string) {
	url := "https://api.telegram.org/bot" +
		botToken +
		"/sendDocument"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	path, _ := filepath.Abs(filePath)
	file, errFile1 := os.Open(path)
	defer file.Close()
	fmt.Println("Sending:", path)
	part1,
		errFile1 := writer.CreateFormFile("document", path)
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
	}
	_ = writer.WriteField("chat_id", chatId)
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
