package main

import (
	"fmt"
	"os"
)

func start_dbus() {
	fmt.Fprintln(os.Stderr, "Not implemented")
}

func main() {
	fmt.Println("twinits: the worst init system")
	var pid int = os.Getpid()
	var targetpid int = 1
	if pid != targetpid {
		fmt.Fprintln(os.Stderr, "Not init system, cannot operate")
		os.Exit(1)
	}
	fmt.Println("Starting twinits")
}
