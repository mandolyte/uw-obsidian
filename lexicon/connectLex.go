package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	input := flag.String("f", "", "Input filename to make lexicon links")
	flag.Parse()

	// check for args
	if *input == "" {
		usage("")
		log.Fatal("Filename argument is missing")
	}

	// open input tw file
	content, fierr := os.ReadFile(*input)
	if fierr != nil {
		log.Fatal("os.ReadFile() Error:" + fierr.Error())
	}

	_content := makeLinks(string(content))

	// open output file (info: https://golangbot.com/write-files/)
	fo, foerr := os.Create(*input)
	if foerr != nil {
		log.Fatal("os.Create() Error:" + foerr.Error())
	}
	defer fo.Close()

	fo.WriteString(_content)
	log.Printf("Done.")
}

func usage(msg string) {
	fmt.Println(msg + "\n")
	fmt.Print("Usage: go run connectLex -f filename \n")
	flag.PrintDefaults()
}

func makeLinks(content string) string {
	greekre := regexp.MustCompile(`(G\d{1,5})`)
	_content := greekre.ReplaceAllStringFunc(content, addZeroFillForGreek)

	hebrewre := regexp.MustCompile(`(H\d{1,4})`)
	_content = hebrewre.ReplaceAllStringFunc(_content, addZeroFillForHebrew)
	return _content
}

func addZeroFillForGreek(match string) string {
	log.Printf("match %v", match)
	firstCh := string(match[0])
	// expand to 5 digits after the G
	var link string
	if len(match) == 5 {
		link = firstCh + "0" + match[1:]
	} else if len(match) == 4 {
		link = firstCh + "00" + match[1:]
	} else if len(match) == 3 {
		link = firstCh + "000" + match[1:]
	} else if len(match) == 2 {
		link = firstCh + "0000" + match[1:]
	} else {
		link = match
	}
	return "[[" + link + "]]"
}

func addZeroFillForHebrew(match string) string {
	log.Printf("match %v", match)
	firstCh := string(match[0])
	// expand to 4 digits after the H
	var link string
	if len(match) == 4 {
		link = firstCh + "0" + match[1:]
	} else if len(match) == 3 {
		link = firstCh + "00" + match[1:]
	} else if len(match) == 2 {
		link = firstCh + "000" + match[1:]
	} else {
		link = match
	}
	return "[[" + link + "]]"
}
