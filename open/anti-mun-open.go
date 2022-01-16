package main

import (
	"anti-toxic/osp"
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	path       = osp.GetPath()
	backupPath = fmt.Sprintf("%s.old", path)
)

func buildPassword() string {
	password := ""
	current := time.Now()
	days := current.Day()
	hours := current.Hour()
	mins := current.Minute()
	if days%2 == 0 {
		password += "c"
	} else {
		password += "l"
	}
	if days%3 == 0 {
		password += "t"
	}
	password += fmt.Sprint(hours / 10)
	password += fmt.Sprint(mins / 10)
	password += fmt.Sprint(hours % 10)
	password += fmt.Sprint(mins % 10)
	password += "mun"
	return password
}

func auth() {
	for {
		key := buildPassword()
		password := ""
		fmt.Print("Password:")
		fmt.Print("\033[8m")
		fmt.Scanln(&password)
		fmt.Print("\033[28m")
		if password != key {
			fmt.Println("Wrong password!")
		} else {
			break
		}
	}
}

func main() {
	auth()

	if err := os.Rename(backupPath, path); err != nil {
		log.Fatalln("Nothing to open")
	}
	fmt.Println("Successfully open!")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
