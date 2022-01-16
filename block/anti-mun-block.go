package main

import (
	"anti-toxic/osp"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	path       = osp.GetPath()
	backupPath = fmt.Sprintf("%s.old", path)
	hosts      = []string{
		"arcadespot.com",
		"backrooms-rewandered.fandom.com",
		"backrooms-sandbox-2.wikidot.com",
		"backrooms-wiki.wdfiles.com",
		"backrooms-wiki.wikidot.com",
		"backrooms.fandom.com",
		"devforum.roblox.com",
		"files.minecraftforge.net",
		"funhtml5games.com",
		"game-cdn.poki.com",
		"gamebai.man.vn",
		"gamejolt.com",
		"gamestips.site",
		"handleheldgame.vn",
		"java-for-minecraft.com",
		"kamiapp.wufoo.com",
		"kzsection.info",
		"linkneverdie.net",
		"minecraft-windows-10-edition.softonic.vn",
		"minefc.com",
		"poki.com",
		"stealthygaming.com",
		"supermarioemulator.com",
		"tlauncher.org",
		"tlauncher.softonic.vn",
		"tmodloader.softonic.vn",
		"web.roblox.com",
		"www.boredbro.com",
		"www.crazygames.com",
		"www.kinguin.net",
		"www.mariogames.be",
		"www.minecraft.net",
		"www.mixfreegames.com",
		"www.programarsiv.com",
		"www.ranker.com",
		"www.scaryexe.com",
		"www.trochoi.net",
		"www.ubestgames.com",
		"www.youtube.com",
		"youtube.com",
	}
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
