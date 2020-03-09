package main

import (
	"log"
	"strings"

	"github.com/firefox-boxes/boxes"
)

const NATIVE_MESSAGING_MANIFEST = `{
	"name": "boxes_ext_native_shell",
	"description": "Boxes backend shell",
	"path": "{path}",
	"type": "stdio",
	"allowed_extensions": [ "boxes@whatsyouridea.com" ]
}`

func FirefoxSetup(p boxes.ProbeResult) {
	log.Println("Generating native manifest...")
	manifest := strings.ReplaceAll(NATIVE_MESSAGING_MANIFEST, "{path}", p.GetRelDir("/bin/boxes-ext-native-shell"))
	AddToFirefox(manifest, p)
}