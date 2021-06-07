package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	inputTw := flag.String("tw", "", "Input tw filename")
	outputTw := flag.String("output", "../vault/bible", "Output TW filename")
	flag.Parse()

	// check for args
	if *inputTw == "" {
		usage("")
		log.Fatal("tw filename argument is missing")
	}

	// open input tw file
	content, fierr := os.ReadFile(*inputTw)
	if fierr != nil {
		log.Fatal("os.ReadFile() Error:" + fierr.Error())
	}

	// open output file (info: https://golangbot.com/write-files/)
	fo, foerr := os.Create(*outputTw)
	if foerr != nil {
		log.Fatal("os.Create() Error:" + foerr.Error())
	}
	defer fo.Close()

	_content := fixLinks(string(content))
	fo.WriteString(_content)
	log.Printf("Done with %v to %v", *inputTw, *outputTw)
}

func usage(msg string) {
	fmt.Println(msg + "\n")
	fmt.Print("Usage: go run xformtn -bookId bookId -tntsv inputtn.tsv -dir outputDirectory \n")
	flag.PrintDefaults()
}

func fixLinks(content string) string {
	_content := strings.Replace(content, `rc://*/ta/man/`, "../../", -1)
	_content = strings.Replace(_content, `rc://en/ta/man/`, "../../", -1)
	_content = strings.Replace(_content, `rc://*/tw/dict/`, "../../", -1)
	return _content
}
