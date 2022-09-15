package gsuiteClient

import (
	"fmt"
	"log"
	"net/http"

	colour "github.com/logrusorgru/aurora/v3"
)

func Query(email string) {
	url := fmt.Sprintf("https://mail.google.com/mail/gxlu?email=%s", email)

	resp, err := http.Head(url)

	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	if resp.Header["Set-Cookie"] != nil {
		fmt.Println(colour.Green("[+] " + email + " Exists"))
	} else {
		fmt.Println(colour.Red("[-] " + email + " Does Not Exist"))
	}
}
