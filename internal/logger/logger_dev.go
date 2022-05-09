//go:build !release

package logger

import (
	"log"

	"github.com/scutrobotlab/asuwave/internal/helper"
)

var Log *log.Logger

func init() {
	Log = log.New()
	Log.Println(helper.GetVersion())
	Log.Println("Develop Mode.")
}
