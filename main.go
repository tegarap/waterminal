package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

var cli struct {
	Login struct {
		Force bool   `help:"Force login." short:"f"`
		Phone string `help:"Your phone number." short:"p" required:"true"`
	} `cmd:"" help:"WhatsApp login"`
	Message struct {
		Phone       string `help:"Your phone number." short:"p" required:"true"`
		Destination string `help:"Phone number destination." short:"d" required:"true"`
		Text        string `help:"Text messages." short:"t" required:"true"`
	} `cmd:"" help:"WhatsApp send message"`
}

func login(phone string, force bool) string {
	return fmt.Sprintf("Login with %v -> force = %v", phone, force)
}

func message(sender, destination, text string) string {
	return fmt.Sprintf("%v Sending to -> %v, message = %v", sender, destination, text)

}

func main() {
	ctx := kong.Parse(&cli,
		kong.Name("wacli"),
		kong.Description("A WhatsApp chat run on terminal."),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}))
	switch ctx.Command() {
	case "login":
		result := login(cli.Login.Phone, cli.Login.Force)
		fmt.Println(result)

	case "message":
		result := message(cli.Message.Phone, cli.Message.Destination, cli.Message.Text)
		fmt.Println(result)
	}
}
