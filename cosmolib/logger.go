package main

import (
	"log"
	"os"
)

// Trace log level
var Trace = log.New(os.Stdout, "[CosmoBackendTrace]: ", log.Ldate|log.Ltime)

// Info log level
var Info = log.New(os.Stdout, "[CosmoBackendInfo]: ", log.Ldate|log.Ltime)

// Warning log level
var Warning = log.New(os.Stdout, "[CosmoBackendWarning]: ", log.Ldate|log.Ltime|log.Lshortfile)

// Error log level
var Error = log.New(os.Stdout, "[CosmoBackendError]: ", log.Ldate|log.Ltime|log.Lshortfile)
