//+build linux

package main

import (
	"fmt"

	"github.com/firefox-boxes/boxes"
)

func CreateShortcut(p boxes.ProbeResult) {
	fmt.Println("\u001b[1;32mAdd desktop entry to get an icon, executable is installed at:\u001b[m")
	fmt.Println("\u001b[1;32m  " + p.GetRelDir("/bin/boxes") + "\u001b[m")
}