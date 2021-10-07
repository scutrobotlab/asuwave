package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/scutrobotlab/asuwave/helper"
	"github.com/scutrobotlab/asuwave/option"
	"github.com/scutrobotlab/asuwave/serial"
	"github.com/scutrobotlab/asuwave/server"
	"github.com/scutrobotlab/asuwave/variable"
)

func main() {
	vFlag := false
	uFlag := false
	bFlag := false
	pFlag := -1
	flag.BoolVar(&vFlag, "v", false, "show version")
	flag.BoolVar(&uFlag, "u", false, "check update")
	flag.BoolVar(&bFlag, "b", false, "start browser")
	flag.IntVar(&pFlag, "p", 8000, "port to bind")
	flag.Parse()

	if vFlag {
		fmt.Println(helper.GetVersion())
		os.Exit(0)
	}

	if uFlag {
		helper.CheckUpdate()
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
		helper.StartBrowser("http://localhost:" + strconv.Itoa(option.Config.Port))
	}

	c := make(chan string, 10)
	go serial.GrReceive()
	go serial.GrTransmit()
	go serial.GrRxPrase(c)
	server.Start(c, &fsys)
}
