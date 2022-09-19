package gsuiteClient

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	colour "github.com/logrusorgru/aurora/v3"
)

func Query(email string, path string) {
	url := fmt.Sprintf("https://mail.google.com/mail/gxlu?email=%s", email)

	resp, err := http.Head(url)

	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	if resp.Header["Set-Cookie"] != nil {
		fmt.Println(colour.Green("[+] " + email + " Exists"))
		if path != "" {
			WriteFile(path, email)
		}
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

func WriteFile(path string, data string) error {
	if _, err := os.Stat(path); err == nil {
		f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err.Error())
		}

		defer f.Close()
		if _, err := f.WriteString(data + "\n"); err != nil {
			log.Fatal(err.Error())
		}
		return nil
	} else if errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(path)
		if err != nil {
			log.Fatal(err.Error())
		}

		defer f.Close()
		w := bufio.NewWriter(f)
		fmt.Fprintln(w, data)
		return w.Flush()
	} else {
		log.Fatal("Schrondingers error.... How did you manage this?")
		return nil
	}
}
