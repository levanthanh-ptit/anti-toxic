package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	// path := "%SystemRoot%\\System32\\drivers\\etc\\hosts"
	path       = "/etc/hosts"
	backupPath = fmt.Sprintf("%s.old", path)
	hosts      = []string{"youtube.com", "www.youtube.com"}
)

func getHostsString() string {
	data := "# Block hosts\n"
	for _, host := range hosts {
		data += fmt.Sprintf("127.0.0.1\t%s\n", host)
	}
	data += "\n"
	return data
}

func writeBackUp(path string) {
	// prevent double block
	if _, err := os.Stat(backupPath); err == nil {
		log.Fatalln("Already blocked!")
	}
	// backup a host file
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	file, err := os.Create(backupPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	file.Write(content)
}

func main() {
	writeBackUp(path)
	hostFile, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalln(err)
	}
	defer hostFile.Close()
	content := getHostsString()
	fmt.Println("# Hosts blocked:")
	fmt.Println(strings.Join(hosts, "\n"))
	if _, err := hostFile.WriteString(content); err != nil {
		log.Fatalln(err)
	}
}
