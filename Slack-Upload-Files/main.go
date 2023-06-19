package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5366412634135-5379228644567-IydqyWwQTpBtyqjn07DGL5Ul")
	os.Setenv("CHANNEL_ID", "C05B9B5MRU4")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelsArr := []string{os.Getenv("CHANNEL_ID")}
	filesArr := []string{"ZIPL.pdf"}

	for i := 0; i < len(filesArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelsArr,
			File:     filesArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		fmt.Printf("Name: %s, URL: %s", file.Name, file.URL)
	}
}
