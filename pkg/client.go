package gsuiteClient

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

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

func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
