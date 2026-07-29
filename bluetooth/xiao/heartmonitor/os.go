//go:build !baremetal

package main

import "os"

func connectName() string {
	if len(os.Args) < 2 {
		println("usage: heartratemonitor [name]")
		os.Exit(1)
	}

	// look for device with specific name
	name := os.Args[1]

	return name
}

// done just prints a message and allows program to exit.
func done() {
	println("Done.")
}
