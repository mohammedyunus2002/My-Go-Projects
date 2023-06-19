package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analytics <-chan *slacker.CommandEvent) {
	for events := range analytics {
		fmt.Println("Command Events")
		fmt.Println(events.Timestamp)
		fmt.Println(events.Command)
		fmt.Println(events.Parameters)
		fmt.Println(events.Event)
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5366412634135-5378043256133-Zo8OK2HrK9nQbGZDohQf9MaP")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A05B6QK0ABV-5380837882355-4d66c280fa6826ae7be4aa3944f6db4ee93e430afb596b4993e807ffe1ccaee4")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples:    []string{"my yob is 2020"},
		Handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
			year := r.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				w.Reply("Invalid year rpovided")
			}
			age := 2023 - yob
			res := fmt.Sprintf("age is %d", age)
			w.Reply(res)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
