package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"

	"github.com/scutrobotlab/asuwave/option"
	"github.com/scutrobotlab/asuwave/serial"
	"github.com/scutrobotlab/asuwave/server"
	"github.com/scutrobotlab/asuwave/variable"
)

var (
	githash   string
	buildtime string
	goversion string
)

func main() {
	vFlag := false
	bFlag := false
	pFlag := -1
	flag.BoolVar(&vFlag, "v", false, "show version")
	flag.BoolVar(&bFlag, "b", false, "start browser")
	flag.IntVar(&pFlag, "p", -1, "port to bind")
	flag.Parse()

	if vFlag {
		fmt.Printf("asuwave %s\n", githash)
		fmt.Printf("build time %s\n", buildtime)
		fmt.Println(goversion)
		os.Exit(0)
	}

	option.Load()
	variable.Load()

	if pFlag >= 0 && pFlag <= 65535 {
		option.Config.Port = pFlag
	}

	option.Save()

	fsys := getFS()

	if bFlag {
		startBrowser()
	}

	c := make(chan string, 10)
	go serial.GrReceive()
	go serial.GrTransmit()
	go serial.GrRxPrase(c)
	server.Start(c, &fsys)
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
