#!/bin/sh

INPUT=$HOME/Projects/git.door43.org/unfoldingword/en_tn/tn_1CO.tsv

go run xformtn.go -bookId 1co -tntsv $INPUT 
