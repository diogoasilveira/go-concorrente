package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

const charset = "abcdefghijklmnopqrstuvwxyz1234567890"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func randString(length int) string {
	return StringWithCharset(length, charset)
}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func generateContent(out chan string) {
	for i := 0; i < 100; i++{
		out <- randString(5)
	}
	close(out)
}

func filterContent(in chan string, out chan string) {
	for {
		word, ok := <-in
		if !ok{
			break
		}
		if isLetter(word) {
			out <- word
		}
	}
	close(out)
}

func main() {
	rawContent := make(chan string)
	filteredContent := make(chan string)

	go generateContent(rawContent)
	go filterContent(rawContent, filteredContent)

	for {
		alpha, ok := <-filteredContent
		if !ok {
			break
		}
		fmt.Printf("alpha: <%s>\n", alpha)
	}
}
