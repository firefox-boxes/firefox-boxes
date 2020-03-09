package main

import (
	"io"
	"log"

	"os"

	"github.com/firefox-boxes/boxes"
	"github.com/markbates/pkger"
)

func Copy(src io.Reader, dst string) {
	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, os.FileMode(uint32(0764)))
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, src)
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = dstFile.Sync()
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func MkdirIfNotExists(path string, mode os.FileMode) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, mode)
	}
}

func copyBinaries(p boxes.ProbeResult) {
	f, err := pkger.Open("/binary/dist/boxes")
	if err != nil {
		log.Fatalln(err.Error())
	}
	Copy(f, p.GetRelDir("/bin/boxes"))
	f.Close()

	f, err = pkger.Open("/binary/dist/boxes-ext-native-shell")
	if err != nil {
		log.Fatalln(err.Error())
	}
	Copy(f, p.GetRelDir("/bin/boxes-ext-native-shell"))
	f.Close()

	f, err = pkger.Open("/binary/dist/boxes-ipc")
	if err != nil {
		log.Fatalln(err.Error())
	}
	Copy(f, p.GetRelDir("/bin/boxes-ipc"))
	f.Close()

	f, err = pkger.Open("/binary/dist/boxes-shell")
	if err != nil {
		log.Fatalln(err.Error())
	}
	Copy(f, p.GetRelDir("/bin/boxes-shell"))
	f.Close()

	f, err = pkger.Open("/res/icon.png")
	if err != nil {
		log.Fatalln(err.Error())
	}
	Copy(f, p.GetRelDir("/icon.png"))
	f.Close()
}

func CopyBinaries(p boxes.ProbeResult) {
	log.Println("Creating bin directory...")
	MkdirIfNotExists(p.GetRelDir("bin"), os.FileMode(uint32(0774)))
	log.Println("Copying binaries...")
	copyBinaries(p)
}