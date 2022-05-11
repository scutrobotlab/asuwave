package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/scutrobotlab/asuwave/internal/helper"
	"github.com/scutrobotlab/asuwave/internal/option"
	"github.com/scutrobotlab/asuwave/internal/serial"
	"github.com/scutrobotlab/asuwave/internal/server"
	"github.com/scutrobotlab/asuwave/pkg/elffile"
)

func main() {
	vFlag := false
	uFlag := false
	bFlag := false
	flag.BoolVar(&vFlag, "v", false, "show version")
	flag.BoolVar(&uFlag, "u", false, "check update")
	flag.BoolVar(&bFlag, "b", true, "start browser")
	flag.IntVar(&helper.Port, "p", 8000, "port to bind")
	flag.Parse()

	if vFlag {
		fmt.Println(helper.GetVersion())
		os.Exit(0)
	}

	if uFlag {
		helper.CheckUpdate(false)
		os.Exit(0)
	}

	option.Load()

	fsys := getFS()

	if bFlag {
		helper.StartBrowser("http://localhost:" + strconv.Itoa(helper.Port))
	}

	c := make(chan string, 10)
	go serial.GrReceive()
	go serial.GrTransmit()
	go serial.GrRxPrase(c)
	go elffile.FileWatch()
	server.Start(c, &fsys)
}
