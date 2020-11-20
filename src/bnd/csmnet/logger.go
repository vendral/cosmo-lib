package csmnet

import (
	"os"
    "log"
)

var Trace = log.New(os.Stdout, "[CosmoBackendTrace]: ", log.Ldate|log.Ltime)
var Info = log.New(os.Stdout, "[CosmoBackendInfo]: ", log.Ldate|log.Ltime)
var Warning = log.New(os.Stdout, "[CosmoBackendWarning]: ", log.Ldate|log.Ltime|log.Lshortfile)
var Error = log.New(os.Stdout, "[CosmoBackendError]: ", log.Ldate|log.Ltime|log.Lshortfile)
