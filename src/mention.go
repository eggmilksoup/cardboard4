// mention.go version 4.0.0

package main
import "bufio"
import "fmt"
import "github.com/bwmarrin/discordgo"
import "os"

func mention(discord *discordgo.Session, id string) {
			usr, err := discord.User(id)
			if err != nil {
				fmt.Fprintf(os.Stderr, "discordgo.Session.User: %s\n", err.Error())
				os.Exit(1)
			}
			fmt.Println(usr.Mention())
}

func main() {
	
	discord, err := discordgo.New("Bot " + os.Getenv("key"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "discordgo.New: %s\n", err.Error())
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i ++ {
			mention(discord, os.Args[i])
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			mention(discord, scanner.Text())
		}

		err = scanner.Err()
		if err != nil {
			fmt.Fprintf(os.Stderr, "bufio.Scanner.Scan: %s\n", err.Error())
			os.Exit(1)
		}
	}
}
