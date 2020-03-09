package main

import (
	"log"
	"net/rpc"
	"os/exec"

	"github.com/ProtonMail/go-autostart"
	"github.com/firefox-boxes/boxes"
)

func InitialSetupOfBoxes() {
	client, err := rpc.Dial("tcp", "127.0.0.1:6688")
	for err != nil {
		client, err = rpc.Dial("tcp", "127.0.0.1:6688")
	}
	defer client.Close()
	req := "box:new box.svg|Default|" + boxes.GetInstallations()[0].Exec
	log.Println("  Running " + req)
	res := ""
	client.Call("IPC.Handle", &req, &res)
	req = "default:set " + res
	log.Println("  Setting " + res + " as default box")
	res = ""
	client.Call("IPC.Handle", &req, &res)
}

func BinarySetup(p boxes.ProbeResult) {
	log.Println("Adding Boxes to autostart...")
	app := &autostart.App{
		Name: "Firefox (Boxes)",
		DisplayName: "Firefox (Boxes)",
		Exec: []string{p.GetRelDir("/bin/boxes-ipc")},
	}
	if err := app.Enable(); err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Starting boxes-ipc...")
	cmd := exec.Command(p.GetRelDir("/bin/boxes-ipc"))
	err := cmd.Start()
	if err != nil {
		log.Fatalln(err.Error())
	}
	cmd.Process.Release()

	log.Println("Setting up boxes-ipc...")
	InitialSetupOfBoxes()
}