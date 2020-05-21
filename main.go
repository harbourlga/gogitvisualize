package main

import (
	"flag"
)

func main() {
	var folder string
	var email string
	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")
	flag.Parse()

	InitZapLog()
	defer logger.Sync()

	if folder != "" {
		scan(folder)
		return
	}
	Level("Scan Email",email)
	stats(email)
}
