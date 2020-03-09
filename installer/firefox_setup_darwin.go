//+build darwin

package main

import (
	"log"
	"os"
	"path"

	"github.com/firefox-boxes/boxes"
)

func AddToFirefox(manifest string, p boxes.ProbeResult) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err.Error())
	}
	manifestDir := path.Join(homeDir, "/Library/Application Support/Mozilla/NativeMessagingHosts/")
	MkdirIfNotExists(manifestDir, os.FileMode(uint32(0774)))
	manifestPath := path.Join(manifestDir, "boxes-ext-native-shell.json")

	log.Println("Writing to '" + manifestPath + "'...")
	
	f, err := os.OpenFile(manifestPath, os.O_WRONLY|os.O_CREATE, os.FileMode(uint32(0664)))
	defer f.Close()
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = f.WriteString(manifest)
	if err != nil {
		log.Fatalln(err.Error())
	}
}