#!/bin/sh
# remove .obsidian
rm -rf ../vault/.obsidian
OUTPUTDIR=../vault/lexicon

# copy the greek
INPUTDIR=$HOME/Projects/git.door43.org/unfoldingword/en_ugl/content
echo working on Greek Lexicon
for i in `ls $INPUTDIR`
do
    cp $INPUTDIR/$i/01.md $OUTPUTDIR/$i.md
done

# copy the hebrew
INPUTDIR=$HOME/Projects/git.door43.org/unfoldingword/en_uhal/content

echo working on Hebrew Lexicon
for i in `ls $INPUTDIR`
do
  cp $INPUTDIR/$i $OUTPUTDIR/$i
done
