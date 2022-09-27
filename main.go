package main

import (
	"fmt"
	"log"
	"os"

	gsuiteClient "gsuite-enum/pkg"

	"github.com/akamensky/argparse"
)

func init() {
	parser := argparse.NewParser("gsuite-enum", "A simple tool to enumerate existing users within gSuite or Gmail")

	email := parser.String("e", "email", &argparse.Options{Required: false, Help: "Email address to query. Example: user@domain.com"})
	userList := parser.String("U", "userlist", &argparse.Options{Required: false, Help: "Specify userlist to enumerate"})
	outFile := parser.String("o", "outfile", &argparse.Options{Required: false, Help: "Output validated emails to file"})

	err := parser.Parse(os.Args)
	if err != nil {
		log.Fatal(err.Error())
	}

	if *email != "" {
		gsuiteClient.Query(*email, *outFile)
	} else if *userList != "" {
		users, err := gsuiteClient.ReadFile(*userList)
		if err != nil {
			log.Fatal(err.Error())
		}
		for _, line := range users {
			gsuiteClient.Query(line, *outFile)
		}

	} else {
		fmt.Print(parser.Usage(err))
	}
}

func main() {
}
