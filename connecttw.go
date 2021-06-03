package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBook := flag.String("book", "", "Input Book filename")
	inputTwl := flag.String("twl", "", "Input twl filename")
	outputDir := flag.String("dir", "vault", "Output folder; default 'vault'")
	flag.Parse()

	// check for args
	if *inputBook == "" {
		log.Fatal("book argument is missing")
	}
	if *inputTwl == "" {
		log.Fatal("TWL argument is missing")
	}

	// open output file (info: https://golangbot.com/write-files/)
	fo, foerr := os.Create(*outputDir + "/" + *inputBook)
	if foerr != nil {
		log.Fatal("os.Create() Error:" + foerr.Error())
	}
	defer fo.Close()

	// open input book (info: https://golang.org/pkg/io/ioutil/#ReadFile)
	fcontent, err := os.ReadFile(*inputBook)
	if err != nil {
		log.Fatal("os.Readfile() Error:" + err.Error())
	}

	var verseTable [][]string
	verses := strings.Split(string(fcontent), "\n")
	for _, verse := range verses {
		cells := strings.Split(verse, " ")
		if cells[0] == "" {
			continue
		}
		// remove punctuation
		for i := range cells {
			cells[i] = strings.TrimRight(cells[i], ",;.")
		}
		verseTable = append(verseTable, cells)
	}

	// open input twl file
	var r *csv.Reader
	fi, fierr := os.Open(*inputTwl)
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
	var footnote = ""
	var row uint64
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

		// match against book
		twlRef := cells[0]
		twlId := cells[1]
		twlQuote := cells[3]
		twlOccurrence := cells[4]
		twlLink := cells[5]
		for i := 0; i < len(verseTable); i++ {
			bookref := verseTable[i][0]
			if twlRef == bookref {
				occurrence := 0
				_footnote := ""
				for j := range verseTable[i] {
					if twlQuote == verseTable[i][j] {
						occurrence++
						if strconv.Itoa(occurrence) == twlOccurrence {
							log.Printf("Matched! ref:%v, occurence: %v, quote:%v", twlRef, twlOccurrence, twlQuote)
							verseTable[i][j], _footnote = rewrite(verseTable[i][j], twlId, twlLink)
							footnote += _footnote + "\n"
						}
					}
				}
				break
			}
		}
	}
	log.Printf("Number of rows in TWL: %v", row)
	//log.Printf("Verses:\n%v", verseTable)
	for i := 0; i < len(verseTable); i++ {
		for j := 0; j < len(verseTable[i]); j++ {
			fo.WriteString(verseTable[i][j] + " ")
		}
		fo.WriteString("\n")
	}
}

func usage(msg string) {
	fmt.Println(msg + "\n")
	fmt.Print("Usage: go run connecttw -book inputbook.md -twl inputtwl.tsv -dir outputDirectory \n")
	flag.PrintDefaults()
}

func rewrite(cell string, id string, link string) (string, string) {
	localLink := strings.TrimPrefix(link, "rc://*/tw/dict/")
	localLink = strings.ReplaceAll(localLink, "/", "-")
	_cell := cell + "[^" + id + "]"
	return _cell, "a footnote"
}

/* Code Graveyard


f, err := os.Create("test.txt")
if err != nil {
		fmt.Println(err)
		return
}
l, err := f.WriteString("Hello World")
if err != nil {
		fmt.Println(err)
		f.Close()
		return
}
fmt.Println(l, "bytes written successfully")
err = f.Close()
if err != nil {
		fmt.Println(err)
		return
}


*/
