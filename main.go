package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Flag defines the command line arguments
	// Use "go run . -help" to see the default values and usage information
	// Use "go run . -planet Mars" to see the result with given parameters
	path := flag.String("path", "myapp.log", "The path to log that should be analyzed")
	level := flag.String("planet", "Earth", "The planet name that will be searched for")

	// parse the command line arguments if they are provided; if none, use the default values
	flag.Parse()

	// since the path returns pointer for the path variable instead of literal, we need asterisk
	f, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	// Here, defer means execute this function once the other calls and methods are done.
	// With this, we make sure that we can still continue to work on the file
	// and once we are done in the main, it will be closed
	// if there are multiple defers, it works in LIFO
	defer f.Close()
	r := bufio.NewReader(f)
	// infinite loop
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			fmt.Print(s)
		}
	}
}
