// mkthread.go version 4.0.0

package main
import "fmt"
import "github.com/bwmarrin/discordgo"
import "os"
import "strconv"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s duration title", os.Args[0])
		os.Exit(1)
	}

	discord, err := discordgo.New("Bot " + os.Getenv("key"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "discordgo.New: %s", err.Error())
		os.Exit(1)
	}

	dur, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "strconv.Atoi: %s", err.Error())
		os.Exit(1)
	}

	thread, err := discord.ThreadStart(os.Getenv("channel"), os.Args[2], 11, dur)
	if err != nil {
		fmt.Fprintf(os.Stderr, "discordgo.Session.ThreadStart: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println(thread.ID)
}
