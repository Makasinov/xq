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
	for it := 0; it < len(text); it++ {
		var s string
		if it+1 >= len(text) {
			s = text[it:it+1] + " "
		} else {
			s = text[it : it+2]
		}

		if isInsideTag {
			if slash(s) {
				isInsideTag = true
				c = color.New(color.FgYellow, color.Bold)
				_, _ = c.Print(s[0:1])
				c = color.New(color.FgCyan)
				continue
			}
			_, _ = c.Print(s[0:1])
		}

		if al(s) {
			isInsideTag = true
			c = color.New(color.FgYellow, color.Bold)
			_, _ = c.Print(s[0:1])
			c = color.New(color.FgCyan)
			continue
		}
		if ar(s) {
			isInsideTag = false
			c = color.New(color.FgYellow, color.Bold)
			_, _ = c.Print(s[1:2])
			continue
		}
		if !isInsideTag {
			c = color.New(color.FgHiGreen, color.Bold)
			if s[0:1] != ">" && s[0:1] != "<" {
				_, _ = c.Print(s[0:1])
			} else {
				_, _ = c.Print("\n\t")
			}
		}
	}

}

const (
	AL    = "<"
	AR    = ">"
	SLASH = "/"
)

func al(s string) bool {
	if s[0:1] == AL {
		return true
	}
	return false
}

func ar(s string) bool {
	if s[1:2] == AR {
		return true
	}
	return false
}

func slash(s string) bool {
	if s[0:1] == SLASH {
		return true
	}
	return false
}
