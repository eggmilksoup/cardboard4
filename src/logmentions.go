// logmentions.go version 4.0.0

package main
import "fmt"
import "github.com/bwmarrin/discordgo"
import "os"
import "time"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s delay \n", os.Args[0])
		os.Exit(1)
	}
	discord, _ := discordgo.New("Bot " + os.Getenv("key"))
	discord.Open()
	discord.UpdateGameStatus(0, "always learning, ever growing")
	discord.AddHandler( func(discord *discordgo.Session, event *discordgo.MessageCreate) {
		for _, usr := range event.Message.Mentions {
			me, _ := discord.User("@me")
			if usr.ID == me.ID {
				fmt.Printf("%s:%s:%s\n", event.Message.ChannelID, event.Message.ID, event.Message.Content)
			}
		}
	})
	delay, _ := time.ParseDuration(os.Args[1])
	time.Sleep(delay)
	discord.UpdateGameStatus(0, "collecting eldritch data")
}
