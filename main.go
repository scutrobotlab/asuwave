package main

import (
	"flag"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"time"

	"github.com/scutrobotlab/asuwave/option"
	"github.com/scutrobotlab/asuwave/serial"
	"github.com/scutrobotlab/asuwave/server"
	"github.com/scutrobotlab/asuwave/variable"
)

func main() {
	bFlag := false
	pFlag := flag.Int("p", -1, "port to bind")
	flag.BoolVar(&bFlag, "b", false, "start browser")
	flag.Parse()

	option.Load()
	variable.Load()

	if pFlag != nil && *pFlag > 0 && *pFlag < 65535 {
		option.Config.Port = *pFlag
	}

	option.Save()

	if bFlag {
		startBrowser()
	}

	c := make(chan string, 10)
	go serial.GrReceive()
	go serial.GrTransmit()
	go serial.GrRxPrase(c)
	server.Init(c)
}

func startBrowser() {
	var commands = map[string]string{
		"windows": "explorer.exe",
		"darwin":  "open",
		"linux":   "xdg-open",
	}
	run, ok := commands[runtime.GOOS]
	if !ok {
		log.Printf("don't know how to open things on %s platform", runtime.GOOS)
	} else {
		go func() {
			log.Println("Your browser will start in 3 seconds")
			time.Sleep(3 * time.Second)
			exec.Command(run, "http://localhost:"+strconv.Itoa(option.Config.Port)).Start()
		}()
	}
}
