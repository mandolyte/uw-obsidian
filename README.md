# uw-obsidian
An experiment to use Obsidian to show static linked resources from unfoldingWord

Goal: on a small scale demonstrate use of Obsidian to view linked, static resources.

Observations:

- this will be done using Go as the programming language
- it will use the Greek text for Titus 1:1-9; the text will be formatted as one verse per line with the verse reference as the first token.


Steps:

1. Clone the TW repo locally.
1. In the `xformtw` folder, run `sh run.sh`. This will fix links in the TW articles so they work locally.
1. In the `connecttw` folder, place the Greek text file and the "TWL" file (`Titus.md` and `twl_TIT.tsv` respectively)
Then Run `sh run.sh`. This will match up the TWL file and the Greek text and add footnotes in the text pointing to the TW articles. It also does a simple connection at the verse level to the translation notes. It will also fix the links in the TWL to work locally.
1. Clone the TN repo locally.
1. In the `xformtn` folder, run `sh run.sh`. This will convert the TSV file to a linkable markdown file suitable for Obsidian. It also changes the `\n` characters to newlines and will fix the links to work locally.
