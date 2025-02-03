package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		// Get the current working directory
		location, _ := os.Getwd()
		split := strings.Split(location, "/")
		currLocation := split[len(split)-1]

		// Print the current directory prompt
		fmt.Printf("> %s ", currLocation)
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			continue
		}

		input = strings.TrimSpace(input)
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		cmd := parts[0]
		arg := strings.Join(parts[1:], " ")

		switch cmd {
		case "crib":
			// Change directory
			if err := os.Chdir(arg); err != nil {
				fmt.Fprintln(os.Stderr, "Error changing directory:", err)
			}

		case "peek":
			// List files in the current directory
			files, _ := os.ReadDir(".")
			for _, file := range files {
				fmt.Println(file.Name())
			}

		case "whereami":
			// Print the current working directory
			location, _ := os.Getwd()
			fmt.Println("Current directory:", location)

		case "cribup":
			// Create a new directory
			os.Mkdir(arg, 0755)

		case "yell":
			// Output the string
			if len(input) > 5 {
				fmt.Println(input[6:])
			} else {
				fmt.Println("Nothing to yell!")
			}

		case "new":
			// Create a new file
			file, err := os.Create(arg)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error creating file:", err)
			}
			file.Close()

		case "cribdown":
			// Remove the directory and its contents
			if err := os.RemoveAll(arg); err != nil {
				fmt.Fprintln(os.Stderr, "Error removing directory:", err)
			}

		default:
			fmt.Println("Unknown command:", cmd)
		}
	}
}
