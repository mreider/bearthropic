package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	initCmd := flag.Bool("init", false, "Initialize bthropic with API keys and settings")
	startCmd := flag.Bool("start", false, "Start an interactive Claude session")
	destroyCmd := flag.Bool("destroy", false, "Delete configuration and remove bthropic binary from /usr/local/bin/")
	flag.Parse()

	// Process --destroy flag first.
	if *destroyCmd {
		// Remove configuration file.
		if err := os.Remove(getConfigPath()); err != nil {
			fmt.Fprintf(os.Stderr, "Error removing config: %v\n", err)
		} else {
			fmt.Println("Configuration file removed.")
		}
		// Remove binary from /usr/local/bin/
		binaryPath := "/usr/local/bin/bthropic"
		if err := os.Remove(binaryPath); err != nil {
			fmt.Fprintf(os.Stderr, "Error removing binary: %v\n", err)
		} else {
			fmt.Println("Binary removed from /usr/local/bin.")
		}
		os.Exit(0)
	}

	// Default to start if neither --init nor --start is provided.
	if !*initCmd && !*startCmd {
		*startCmd = true
	}

	// If starting and config does not exist, prompt user to run --init.
	if *startCmd {
		if _, err := os.Stat(getConfigPath()); os.IsNotExist(err) {
			fmt.Fprintln(os.Stderr, "No Claude API key found. Please run --init")
			os.Exit(1)
		}
	}

	if *initCmd {
		if err := initializeApp(); err != nil {
			fmt.Fprintf(os.Stderr, "Error initializing: %v\n", err)
			os.Exit(1)
		}
	} else if *startCmd {
		session := NewSession()
		if err := session.Start(); err != nil {
			fmt.Fprintf(os.Stderr, "Error in session: %v\n", err)
			os.Exit(1)
		}
	} else {
		flag.Usage()
	}
}
