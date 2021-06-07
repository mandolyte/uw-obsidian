#!/bin/sh
# remove .obsidian
rm -rf ../vault/.obsidian

# folder of just new format TSV for TN content
INPUTDIR=$HOME/Projects/git.door43.org/unfoldingword/en_tw/bible/kt
OUTPUTDIR=../vault/bible/kt
for i in `ls $INPUTDIR/*.md`
do
    F=`basename $i`
    echo $F from $i
    go run xformtw.go -tw $i -output $OUTPUTDIR/$F
done

# folder of just new format TSV for TN content
INPUTDIR=$HOME/Projects/git.door43.org/unfoldingword/en_tw/bible/names
OUTPUTDIR=../vault/bible/names
for i in `ls $INPUTDIR/*.md`
do
    F=`basename $i`
    echo $F from $i
    go run xformtw.go -tw $i -output $OUTPUTDIR/$F
done

# folder of just new format TSV for TN content
INPUTDIR=$HOME/Projects/git.door43.org/unfoldingword/en_tw/bible/other
OUTPUTDIR=../vault/bible/other
for i in `ls $INPUTDIR/*.md`
do
    F=`basename $i`
    echo $F from $i
    go run xformtw.go -tw $i -output $OUTPUTDIR/$F
done