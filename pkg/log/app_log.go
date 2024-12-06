package applog

import (
	"log"
	"os"
)

var Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
var Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

func Load() {
	log.SetOutput(os.Stdout)
}
