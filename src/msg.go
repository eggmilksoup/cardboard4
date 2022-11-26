// msg.go version 4.1.0

package main
import "fmt"
import "github.com/bwmarrin/discordgo"
import "os"

func main() {
	var txt string
	if len(os.Args) == 1 {

		bin, err := os.ReadFile("/dev/stdin")
		if err != nil {
			fmt.Fprintf(os.Stderr, "os.ReadFile: %s\n", err.Error())
			os.Exit(1)
		}

		txt = string(bin)
	} else {
		txt = ""
		for i := 1; i < len(os.Args); i ++ {
			txt = txt + os.Args[i] + " "
		}
	}
	var messages []string

	discord, err := discordgo.New("Bot " + os.Getenv("key"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "discordgo.New %s\n", err.Error())
		os.Exit(1)
	}

	// the whole mess below is to handle the 2000 character limit imposed by
	// Discord, breaking messages into <2000-char chunks labelled like [2/3]
	var msg *discordgo.Message
	if len(txt) >= 2000 {
		for len(txt) >= 1994 {
			var i int
			for i = 1992; i > 1; i -- {
				if txt[i] == '\n' {
					break
				}
			}
			if i == 1 {
				for i = 1992; i > 1; i -- { 
					if txt[i] == ' ' {
						break
					}
				}
				if i == 1 {
					i = 1992
				}
			}
			messages = append(messages, txt[0:i])
			txt = txt[i:len(txt)]
		}
		for i := 0; i < len(messages); i ++ {
			discord.ChannelMessageSend(os.Getenv("channel"),
				fmt.Sprintf("[%d/%d]\n",
					i + 1,
					len(messages) + 1) + messages[i])
		}

		msg, err = discord.ChannelMessageSend(os.Getenv("channel"),
			fmt.Sprintf("[%d/%d]\n",
				len(messages) + 1,
				len(messages) + 1) + txt)
		if err != nil {
			fmt.Fprintf(os.Stderr,
				"discordgo.Session.ChannelMessageSend: %s\n",
				err.Error())
			os.Exit(1)
		}

	} else {

		msg, err = discord.ChannelMessageSend(os.Getenv("channel"), txt)
		if err != nil {
			fmt.Fprintf(os.Stderr,
				"discordgo.Session.ChannelMessageSend: %s\n",
				err.Error())
			os.Exit(1)
		}

	}
	fmt.Println(msg.ID)
}
