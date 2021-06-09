#!/bin/sh

curl -X POST http://localhost:2468/gql -d 'query={docSets{document(bookCode:"TIT"){cvIndex(chapter:1){chapter verses{	verse{ text } } } } } }'
