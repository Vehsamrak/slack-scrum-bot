package main

import (
    "fmt"
    "os"

    "github.com/nlopes/slack"
)

const SLACK_TOKEN_ENVIRONMENT = "IQOPTION_SLACK_TOKEN_SCRUM_BOT"
const BOT_NAME = "scrum"
const CHANNEL_BOTTESTING = "G6J5NPD4Z"
const CHANNEL_AFFILIATE_DEV = "C0FTH89AT"

func main() {
    api := slack.New(os.Getenv(SLACK_TOKEN_ENVIRONMENT))

    groups, err := api.GetGroups(false)

    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("I'm now available in groups:")
    for _, group := range groups {
        fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
    }

    fmt.Println("")

    channelId, timestamp, err := api.PostMessage(
        CHANNEL_BOTTESTING, getMeetingInvitation(),
        slack.PostMessageParameters{
            Username: BOT_NAME,
            AsUser:   true,
        },
    )

    if err != nil {
        fmt.Printf("%s\n", err)
        return
    }

    fmt.Printf("Message successfully sent to channel %s at %s\n", channelId, timestamp)
}

func getMeetingInvitation() string {
    return "Ребята, пойдем соберёмся! (с)"
}
