// msg.go version 4.0.0

package main
import "fmt"
import "github.com/bwmarrin/discordgo"
import "os"

func main() {
	var txt string
	if len(os.Args) == 1 {
		bin, _ := os.ReadFile("/dev/stdin")
		txt = string(bin)
	} else {
		txt = ""
		for i := 1; i < len(os.Args); i ++ {
			txt = txt + os.Args[i] + " "
		}
	}
	var messages []string
	discord, _ := discordgo.New("Bot " + os.Getenv("key"))
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
			discord.ChannelMessageSend(os.Getenv("channel"), fmt.Sprintf("[%d/%d]\n", i + 1, len(messages) + 1) + messages[i])
		}
		msg, _ = discord.ChannelMessageSend(os.Getenv("channel"), fmt.Sprintf("[%d/%d]\n", len(messages) + 1, len(messages) + 1) + txt)
	} else {
		msg, _ = discord.ChannelMessageSend(os.Getenv("channel"), txt)
	}
	fmt.Println(msg.ID)
}
