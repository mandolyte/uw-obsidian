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
1. Clone the greek and hebrew lexicons
1. In the `lexicon` folder, run `sh copyStrongArticles.sh`. This will copy all documents into the `vault/lexicon` folder with proper names (hebrew is ok; greek must be adjusted)
1. Then run `sh run.sh`. This will update all TW articles to point to the lexicon articles in the "Strongs" paragrapch (last one).


Source for Greek Text: https://door43.org/u/unfoldingWord/UGNT/69478b8734/57-TIT.html#056-ch-001

Notes on an impoved method of obtaining the Greek Text.
1. See [here](https://unfoldingword.zulipchat.com/#narrow/stream/209457-SOFTWARE--.20UR/topic/Proskomma/near/242032634)
1. The idea is to query Proskomma to parse the USFM format and return the text in the format needed for this little project.
1. Essential GraphQl query:
```gql
{
  docSets {
    document(bookCode: "TIT") {
      cvIndex(chapter:1) {
        chapter
        verses {
          verse {
            text
          }
        }
      }
    }
  }
}
```
1. Steps (also see the read me in the PK node express project)...
1. cloned proskomma node express and ran `npm install`
2. downloaded a zip of the UFW Greek NT and copied the USFM files to the data subdirectory (actually `./data/usfm`)
3. added all the files to the "toLoad" array in `index.js`
4. ran: `npm run dev`; this took a few minutes
5. pointed my browser to `http://localhost:2468/gql_form` and pasted the query above into the form and clicked the Run Query button


## Miscellaneous Notes

Proskomma docs and tutorial [here](https://doc.proskomma.bible/en/dev/tutorial/index.html)

Other things that could be done:
- use PK (see getText folder) to iterate over all NT
- expand to OT
- expand ULT / UST
- how to do alignment
- move to DCS
- learn how to exploit aliases in the front matter

