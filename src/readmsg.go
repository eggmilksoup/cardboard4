// readmsg.go version 4.0.0

package main
import "bufio"
import "fmt"
import "github.com/bwmarrin/discordgo"
import "os"

func readmsg(discord *discordgo.Session, id string) {

	msg, err := discord.ChannelMessage(os.Getenv("channel"), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "discordgo.Session.ChannelMessage: %s",
			err.Error())
		os.Exit(1)
	}

	fmt.Println(msg.Content)

}

func main() {
	
	discord, err := discordgo.New("Bot " + os.Getenv("key"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "discordgo.New: %s", err.Error())
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i ++ {
			readmsg(discord, os.Args[i])
		}
	} else {

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			readmsg(discord, scanner.Text())
		}

		err = scanner.Err()
		if err != nil {
			fmt.Fprintf(os.Stderr, "bufio.Scanner.Scan: %s\n", err.Error())
			os.Exit(1)
		}

	}
}
