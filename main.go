package main

import(
	"fmt"
	"os"

	"github.com/turnage/graw/reddit"
)

func main(){
	bot, err := reddit.NewBotFromAgentFile("useragent.txt", 0)
	if err != nil {
		//fmt.Println("Failed to create bot: ", err)
		fmt.Fprintln(os.Stderr, err)
		return
	}

	subreddits := []string{"Art", "EarthPorn"}
	
	for _, sub := range subreddits {
		lp_file, err := os.OpenFile(fmt.Sprintf("lastpost_%s", sub), os.O_RDWR | os.O_CREATE, os.FileMode(0666))
		lastpost := make([]byte, 9)
		_, err = lp_file.Read(lastpost)
	
		art, err := bot.Listing(fmt.Sprintf("/r/%s", sub), string(lastpost))
		if err != nil {
			//fmt.Println("Failed to fetch /r/Art: ", err)
			fmt.Fprintln(os.Stderr, "Failed to fetch ", sub, err)
			return
		}
		
		if len(art.Posts) > 0 {
			for _, post := range art.Posts {
				fmt.Println(post.URL)
			}
	
			_, err = lp_file.WriteAt([]byte(art.Posts[0].Name), 0)
			if err != nil {
	//			fmt.Println(err)
				fmt.Fprintln(os.Stderr, err)
			}
		}
		err = lp_file.Close()
	}
}

