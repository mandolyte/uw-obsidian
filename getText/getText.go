package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	//Encode the data
	queryString := `
	{
		docSets {
			document(bookCode: "TIT") {
				cvIndex {
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
	`

	v := url.Values{}
	v.Set("query", queryString)

	//Leverage Go's HTTP Post function to make request
	resp, err := http.PostForm("http://localhost:2468/gql", v)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured:\n%v\nresponse=%v", err.Error(), resp)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, body, "", "  ")

	if err != nil {
		log.Fatalf("json.MarshalIndent() Error:%v", err.Error())
	}
	log.Println(string(prettyJSON.Bytes()))
}
