// logdms.go version 4.0.0

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
    discord, _ := discordgo.New("Bot " + os.Getenv("key"))
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
    duration, _ := time.ParseDuration(os.Args[1])
    time.Sleep(duration)
}
