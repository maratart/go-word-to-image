package main

import (
	"fmt"
	"os"

	wi "go-word-to-image/wordtoimage"
)

func main() {
	if len(os.Args) > 1 {
		word := os.Args[1]
		imgURL, err := wi.GetImageURL(word)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(imgURL)
	} else {
		fmt.Println("USAGE:\n go-word-to-image WORD\n WORD — a word to translate to image")
	}
}
