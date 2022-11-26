// username.go version 4.1.0

package main
import "bufio"
import "fmt"
import "github.com/bwmarrin/discordgo"
import "os"

func username(discord *discordgo.Session, id string) {

	usr, err := discord.User(id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "discordgo.Session.User: %s\n", err.Error())
		os.Exit(1);
	}

	fmt.Println(usr.Username)
}

func main() {

	discord, err := discordgo.New("Bot " + os.Getenv("key"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "discordgo.New: %s\n", err.Error())
		os.Exit(1)
	}

	if len(os.Args) == 1 {

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			username(discord, scanner.Text());
		}

		err = scanner.Err()
		if err != nil {
			fmt.Fprintf(os.Stderr, "bufio.Scanner.Scan: %s\n", err.Error())
			os.Exit(1)
		}

	} else {
		for i := 1; i < len(os.Args); i ++ {
			username(discord, os.Args[i])
		}
	}
}
