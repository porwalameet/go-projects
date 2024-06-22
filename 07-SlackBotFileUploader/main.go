package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	fmt.Println(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"example.txt"}

	for _, file := range fileArr {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     file,
		}
		fileDetails, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("Upload failed: %v\n", err)
			return
		}
		fmt.Printf("File upload details: Name: %s, URL: %s\n", fileDetails.Name, fileDetails.URL)
	}
}
