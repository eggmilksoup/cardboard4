// logdms.go version 4.1.0

package main
import "fmt"
import "github.com/bwmarrin/discordgo"
import "os"
import "time"

func main() {

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "usage: %s duration admin ...\n", os.Args[0])
		os.Exit(1)
	}
	
	discord, err := discordgo.New("Bot " + os.Getenv("key"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "discordgo.New: %s\n", err.Error())
		os.Exit(1)
	}

	discord.Open()
	discord.UpdateGameStatus(0, "")

	discord.AddHandler( func(discord *discordgo.Session, event *discordgo.MessageCreate) {
		// check whether we are in a DM
		if event.Message.GuildID == "" {
		    for i :=2; i < len(os.Args); i ++ {
			if os.Args[i] == event.Message.Author.ID {
			    fmt.Printf("%s:%s:%s:%s\n",
				event.Message.ChannelID,
				event.Message.Author.ID,
				event.Message.ID,
				event.Message.Content)
			}
		    }
		}
	})
	
	duration, err := time.ParseDuration(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "time.ParseDuration: %s\n", err.Error())
		os.Exit(1)
	}
	
	time.Sleep(duration)
}
