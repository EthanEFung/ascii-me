package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/fatih/color"
)

var inputPath string
var outputPath string

var cyan = color.New(color.FgCyan).SprintFunc()

func init() {
	flag.StringVar(&inputPath, "input", "", "path to input file to read")
	flag.StringVar(&inputPath, "i", "", "path to input file to read")
	flag.StringVar(&outputPath, "output", "", "path to output file to write")
	flag.StringVar(&outputPath, "o", "", "path to output file to write")
}

func main() {
	flag.Parse()

	toAsciiMap := make(map[rune]rune)
	reader := bufio.NewReader(os.Stdin)
	var input string
	var err error
	switch {
	case len(inputPath) > 0:
		input, err = readFile(inputPath)
		if err != nil {
			log.Fatal(err)
		}
	case len(os.Args) > 1 && len(os.Args[1]) > 0:
		input = os.Args[1]
	default:
		log.Fatal(errors.New("no input to copy"))
	}

	output := strings.Map(func(r rune) rune {
		if _, has := toAsciiMap[r]; !has && r > unicode.MaxASCII {
			fmt.Printf("\nNon-ASCII rune found: char: %s unicode: %+q", string(r), r)
			fmt.Printf("\nIn sentence: %s", findRuneContext(input, r))
			fmt.Print("\nReplace with ASCII char: ")
			str, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalf(fmt.Sprintf("could not read user input: %s", err))
			}
			newRune, _ := utf8.DecodeRuneInString(str)
			toAsciiMap[r] = newRune
			return toAsciiMap[r]

		} else if r > unicode.MaxASCII {
			return toAsciiMap[r]
		}
		return r
	}, input)

	switch {
	case len(outputPath) > 0:
		if err := writeFile(outputPath, output); err != nil {
			log.Fatal(err)
		}
		fmt.Println("\nwritten to file.")
	default:
		fmt.Println(output)
	}
}

func readFile(path string) (string, error) {
	b, err := os.ReadFile(inputPath)
	if err != nil {
		return "", fmt.Errorf("could not read from file %s: %s", path, err)
	}
	return string(b), nil
}

func writeFile(path, out string) error {
	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("could not write to file %s: %s", path, err)
	}
	defer f.Close()
	f.WriteString(out)
	f.Sync()
	return nil
}

func findRuneContext(s string, r rune) string {
	sentences := strings.Split(s, "\n")
	for _, sentence := range sentences {
		if strings.ContainsRune(sentence, r) {
			output := []string{}
			words := strings.Split(sentence, " ")
			for _, word := range words {
				if strings.ContainsRune(word, r) {
					output = append(output, cyan(word))
				} else {
					output = append(output, word)
				}
			}
			return strings.TrimSpace(strings.Join(output, " "))
		}
	}
	return ""
}
