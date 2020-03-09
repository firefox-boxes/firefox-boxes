package main

import (
	"log"

	"github.com/firefox-boxes/boxes"
)

func main() {
	p := boxes.Probe()
	CopyBinaries(p)
	BinarySetup(p)
	FirefoxSetup(p)
	CreateShortcut(p)
	log.Println("Done!")
}