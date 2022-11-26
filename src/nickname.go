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
	discord, err := discordgo.New("Bot " + os.Getenv("key"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "discordgo.New: %s\n", err.Error())
	}
	discord.GuildMemberNickname("478329250349449216", "@me", os.Args[1])
}
