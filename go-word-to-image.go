/*
go-word-to-image translate a word to the image by google image search.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/bitly/go-simplejson"
)

const (
	regURL = "http://ajax.googleapis.com/ajax/services/search/images?v=1.0&start=10&q="
)

/*
GetImageURL gets image url for given word
*/
func GetImageURL(word string) (string, error) {

	imgURL := ""
	r := regURL + word
	resp, err := http.Get(r)
	if err != nil {
		return imgURL, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return imgURL, err
	}

	js, err := simplejson.NewJson(body)
	if err != nil {
		return imgURL, err
	}

	imgURL = js.GetPath("responseData", "results").GetIndex(1).Get("url").MustString()
	return imgURL, nil
}

func main() {
	if len(os.Args) > 1 {
		word := os.Args[1]
		imgURL, err := GetImageURL(word)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(imgURL)
	} else {
		fmt.Println("USAGE:\n go-word-to-image WORD\n WORD — a word to translate to image")
	}
}
