package main

import (
	"os"
)

func main() {
	os.Exit(0) // want "os.Exit is in main function"
}

func otherfunc() {
	os.Exit(0)
}
