package main

import (
	"fmt"
	"gosseract-win-demo"
	"log"
)

func main() {
	log.Println(gosseract.Version())
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("./test.png")
	client.SetLanguage("eng")
	client.SetPageSegMode(gosseract.PSM_SINGLE_BLOCK)
	text, err := client.Text()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(text)
}
