package main

import(
	"fmt"

	"github.com/turnage/graw/reddit"
)

func main(){
	bot, err := reddit.NewBotFromAgentFile("useragent.txt", 0)
	if err != nil {
		fmt.Println("Failed to create bot: ", err)
		return
	}

	posts, err := bot.Listing("/r/Art", "")
	if err != nil {
		fmt.Println("Failed to fetch /r/Art: ", err)
		return
	}

	for _, post := range posts.Posts[:5] {
		fmt.Printf("[%s] posted [%s] - [%s]\n", post.Author, post.Title, post.URL)
	}
}
