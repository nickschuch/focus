package main

import (
	"fmt"
	"time"

	"github.com/TheCreeper/go-notify"
	"github.com/alecthomas/kingpin"
	"github.com/slack-go/slack"
)

var (
	cliToken   = kingpin.Flag("token", "Token used to authenticate with Slack.").Envar("SLACK_TOKEN").String()
	cliEmoji   = kingpin.Flag("emoji", "Emoji to identify to users that you are in a focus session.").Default(":tomato:").String()
	cliStatus  = kingpin.Flag("status", "Status message to post to Slack.").Default("Focus Session").String()
	cliMinutes = kingpin.Arg("minutes", "How many minutes to snooze on events for.").Required().Int()
)

func main() {
	kingpin.Parse()

	client := slack.New(*cliToken)

	fmt.Println("Pausing your notifications")

	_, err := client.SetSnooze(*cliMinutes)
	if err != nil {
		panic(err)
	}

	available := time.Now().Local().Add(time.Minute * time.Duration(*cliMinutes))

	fmt.Println("Updating your status")

	err = client.SetUserCustomStatus(*cliStatus, *cliEmoji, available.Unix())
	if err != nil {
		panic(err)
	}

	fmt.Println("Time to focus. Will send a sustem notification when session is complete.")

	time.Sleep(time.Minute * time.Duration(*cliMinutes))

	ntf := notify.NewNotification("Focus", "Session complete. Go for a walk, stretch your legs and maybe load up on coffee.")
	ntf.Timeout = notify.ExpiresNever
	if _, err := ntf.Show(); err != nil {
		panic(err)
	}

	fmt.Println("Session Complete!")
}
