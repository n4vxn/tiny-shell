package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		input = strings.TrimSpace(input)
		fmt.Println(input)

		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		cmd := fmt.Sprintf(parts[0])
		_ = fmt.Sprint(parts[1:])

		switch cmd {
		case "show":
			files, _ := os.ReadDir(".")
			for _, file := range files {
				fmt.Println(file.Name())
			}
		}
	}
}
