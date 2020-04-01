package main

import (
	"fmt"
	"git.nosugar.io/gosseract"
	"os"
)

func main() {
	os.Setenv("TESSDATA_PREFIX", `./gosseract`)
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("./test.png")
	client.SetLanguage("eng")
	client.SetPageSegMode(gosseract.PSM_SINGLE_BLOCK)
	text, err := client.Text()
	fmt.Println(text, err)
	fmt.Println(gosseract.Version())
}
