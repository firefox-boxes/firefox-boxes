//+build darwin

package main

import (
	"log"
	"os"

	"github.com/firefox-boxes/boxes"
	"github.com/markbates/pkger"
)

func CreateShortcut(p boxes.ProbeResult) {
	log.Println("Building app bundle...")
	os.MkdirAll("/Applications/Firefox (Boxes).app/Contents", os.FileMode(uint32(0774)))
	os.MkdirAll("/Applications/Firefox (Boxes).app/Contents/MacOS", os.FileMode(uint32(0774)))
	os.MkdirAll("/Applications/Firefox (Boxes).app/Contents/Resources", os.FileMode(uint32(0774)))

	f, err := pkger.Open("/res/Info.plist")
	if err != nil {
		log.Fatalln(err.Error())
	}
	Copy(f, "/Applications/Firefox (Boxes).app/Contents/Info.plist")
	f.Close()

	f, err = pkger.Open("/binary/dist/boxes")
	if err != nil {
		log.Fatalln(err.Error())
	}
	Copy(f, "/Applications/Firefox (Boxes).app/Contents/MacOS/boxes")
	f.Close()

	f, err = pkger.Open("/res/icon.icns")
	if err != nil {
		log.Fatalln(err.Error())
	}
	Copy(f, "/Applications/Firefox (Boxes).app/Contents/Resources/icon.icns")
	f.Close()
}