package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	inputTnTsv := flag.String("tntsv", "", "Input tn tsv filename")
	inputBookId := flag.String("bookId", "", "Input book id")
	outputDir := flag.String("dir", "../vault/tn", "Output folder; default 'vault'")
	flag.Parse()

	// check for args
	if *inputTnTsv == "" {
		log.Fatal("tn tsv filename argument is missing")
	}
	if *inputBookId == "" {
		log.Fatal("book Id argument is missing")
	}

	// open output file (info: https://golangbot.com/write-files/)
	fo, foerr := os.Create(*outputDir + "/" + *inputBookId + ".md")
	if foerr != nil {
		log.Fatal("os.Create() Error:" + foerr.Error())
	}
	defer fo.Close()

	// open input tn file
	var r *csv.Reader
	fi, fierr := os.Open(*inputTnTsv)
	if fierr != nil {
		log.Fatal("os.Open() Error:" + fierr.Error())
	}
	defer fi.Close()
	r = csv.NewReader(fi)

	// ignore expectations of fields per row
	r.FieldsPerRecord = -1
	r.LazyQuotes = true
	r.Comma = '\t'

	// read loop for CSV
	var row uint64
	var markdown string = ""
	var prevH1 = ""
	for {
		// read the csv file
		cells, rerr := r.Read()
		if rerr == io.EOF {
			break
		}
		if rerr != nil {
			log.Fatalf("csv.Read:\n%v\n", rerr)
		}
		if row == 0 {
			row = 1
			log.Printf("header row skipping %v", cells)
			continue
		}
		row++
		// header level 1 is the reference with bookId
		if cells[0] != prevH1 {
			markdown += "# " + cells[0] + "\n"
			prevH1 = cells[0]
		}
		// header level 2 is the row id with content of tags and spt ref
		markdown += "## " + cells[1] + "\n"
		if cells[2] != "" {
			markdown += "Tags:" + cells[2] + "\n"
		}
		/* support reference not needed since it will be in the text content
		if cells[3] != "" {
			markdown += cells[3] + "\n"
		}
		*/
		if cells[4] != "" {
			markdown += "### " + cells[4] + " (" + cells[5] + ")\n"
		}
		if cells[6][0] == '#' {
			fnote, fnerr := os.Create(*outputDir + "/" + *inputBookId + "-" + cells[1] + ".md")
			if fnerr != nil {
				log.Fatal("os.Create() Error:" + fnerr.Error())
			}
			defer fnote.Close()
			_content := fixLinks(cells[6])
			fnote.WriteString(_content)
			markdown += "See [[" + *inputBookId + "-" + cells[1] + "]]"
		} else {
			_content := fixLinks(cells[6])
			markdown += _content + "\n"
		}
		markdown += "\n"
	}
	fo.WriteString(markdown)

	log.Printf("Number of rows in TN: %v", row)
}

func usage(msg string) {
	fmt.Println(msg + "\n")
	fmt.Print("Usage: go run xformtn -bookId bookId -tntsv inputtn.tsv -dir outputDirectory \n")
	flag.PrintDefaults()
}

func fixLinks(content string) string {
	_content := strings.Replace(content, "\\n", "\n", -1)
	_content = strings.Replace(_content, `rc://*/ta/man/translate/`, "", -1)
	_content = strings.Replace(_content, `rc://en/ta/man/translate/`, "", -1)
	_content = strings.Replace(_content, `rc://*/tw/dict/`, "../", -1)
	return _content
}
