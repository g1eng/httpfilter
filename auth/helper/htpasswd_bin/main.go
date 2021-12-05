package main

import (
	"bufio"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		println("usage: htpasswd <username>")
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := s.Text()
		cryptPassword, err := bcrypt.GenerateFromPassword([]byte(t), 12)
		if err != nil {
			log.Fatalln(err)
		} else {
			log.Println(string(cryptPassword))
		}
	}
}
