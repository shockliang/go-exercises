package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	CmdHello   = "hello"
	CmdGoodBye = "bye"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	numLine := 0
	numCommands := 0

	for scanner.Scan() {
		if strings.ToUpper(scanner.Text()) == "Q" {
			break
		} else {
			text := strings.TrimSpace(scanner.Text())

			switch text {
			case CmdHello:
				numCommands++
				fmt.Println("command response: hi!")
			case CmdGoodBye:
				numCommands++
				fmt.Println("command response: bye!")
			}
			if text != "" {
				numLine++
			}
		}
	}

	fmt.Printf("Entered %v lines\n", numLine)
	fmt.Printf("Entered %v commands\n", numCommands)
}
