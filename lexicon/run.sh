#!/bin/sh
# remove .obsidian
rm -rf ../vault/.obsidian

TWBASE=../vault/bible

INPUTDIR=$TWBASE/kt
for i in `ls $INPUTDIR/*.md`
do
    echo working on $i
    go run connectLex.go -f $i
done

INPUTDIR=$TWBASE/names
for i in `ls $INPUTDIR/*.md`
do
    echo working on $i
    go run connectLex.go -f $i
done

INPUTDIR=$TWBASE/other
for i in `ls $INPUTDIR/*.md`
do
    echo working on $i
    go run connectLex.go -f $i
done