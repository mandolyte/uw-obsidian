#!/bin/sh
# remove .obsidian
rm -rf ../vault/.obsidian

# folder of just new format TSV for TN content
INPUTDIR=$HOME/Projects/git.door43.org/unfoldingword/en_tn

for i in `ls $INPUTDIR`
do
    bid=`echo $i | cut -c4-6`
    echo $bid
    go run xformtn.go -bookId $bid -tntsv $INPUTDIR/$i 
done