package main

import (
        "fmt"

        "github.com/turnage/graw/reddit"
)

func main() {
        bot, err := reddit.NewBotFromAgentFile("login.agent", 0)
        if err != nil {
                fmt.Println("Failed to create bot handle: ", err)
                return
        }

        reddit := "www.reddit.com"

        harvest, err := bot.Listing("/r/golang", "")
        if err != nil {
                fmt.Println("Failed to fetch /r/golang: ", err)
                return
        }

        for index,  post := range harvest.Posts[:5] {
                num := index + 1
                fmt.Printf("%d. [%s](%s)\n", num, post.Title, reddit + post.Permalink)
        }
}
