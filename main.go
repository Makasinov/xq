package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {

	fullStr, err := getInput()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	err = validateXML(fullStr)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	beautifyOutput(fullStr)
	return
}

func getInput() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var fullStr string
	for scanner.Scan() {
		fullStr += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return fullStr, nil
}

func validateXML(text string) error {
	var i *interface{}
	return xml.Unmarshal([]byte(text), &i)
}

func beautifyOutput(text string) {
	c := color.New(color.FgYellow)
	isInsideTag := false

	for _, char := range text {
		if isInsideTag && !al(char) && !ar(char) {
			c = color.New(color.FgYellow, color.Bold)
			c.Printf("%c", char)
			continue
		}
		if ar(char) {
			isInsideTag = false
			c = color.New(color.FgYellow)
			c.Printf("%c", char)
			continue
		}
		if al(char) {
			isInsideTag = true
			c = color.New(color.FgYellow)
			c.Printf("%c", char)
			continue
		}
		if !isInsideTag {
			c = color.New(color.FgBlue)
			c.Printf("%c", char)
			continue
		}
	}

}
