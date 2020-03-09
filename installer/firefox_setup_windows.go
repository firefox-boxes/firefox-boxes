//+build windows

package main

import (
	"os"
	"log"
	"golang.org/x/sys/windows/registry"
	"github.com/firefox-boxes/boxes"
)

func AddToFirefox(manifest string, p boxes.ProbeResult) {
	log.Println("Writing to '" + p.GetRelDir("native-manifest.json") + "'...")

	f, err := os.OpenFile(p.GetRelDir("native-manifest.json"), os.O_WRONLY|os.O_CREATE, os.FileMode(uint32(0664)))
	defer f.Close()
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = f.WriteString(manifest)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Adding manifest path to 'SOFTWARE\\Mozilla\\NativeMessagingHosts\\BoxesExtNativeShell'...")

	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Mozilla\NativeMessagingHosts\BoxesExtNativeShell`, registry.SET_VALUE)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer k.Close()

	err = k.SetStringValue("", p.GetRelDir("native-manifest.json"))
	if err != nil {
		log.Fatalln(err.Error())
	}
}