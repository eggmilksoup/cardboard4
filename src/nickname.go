// nickname.go version 4.0.0

package main
import "fmt"
import "github.com/bwmarrin/discordgo"
import "os"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s nickname\n", os.Args[0])
		os.Exit(1)
	}
	discord, _ := discordgo.New("Bot " + os.Getenv("key"))
	discord.UpdateGameStatus(0, "")
	discord.GuildMemberNickname("478329250349449216", "@me", os.Args[1])
}
