package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode"
)

type Count struct {
	count int
	sync.Mutex
}

func getWords(line string) []string {
	return strings.Split(line, " ")
}

func counterLetters(word string) int {
	letter := 0
	for _, ch := range word {
		if unicode.IsLetter(ch) {
			letter++
		}
	}
	return letter
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	totalLetters := Count{}

	var wg sync.WaitGroup
	for {
		if scanner.Scan() {
			line := scanner.Text()
			words := getWords(line)

			for _, word := range words {
				wordCopy := word
				wg.Add(1)
				go func() {
					totalLetters.Lock()
					defer totalLetters.Unlock()
					defer wg.Done()
					sum := counterLetters(wordCopy)

					totalLetters.count += sum
				}()
			}
		} else {
			break
		}
	}

	wg.Wait()
	totalLetters.Lock()
	sum := totalLetters.count
	totalLetters.Unlock()

	fmt.Println("total letters:", sum)
}
