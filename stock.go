package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Quote struct {
	Symbol        string
	Last          string
	Change        string
	ChangePercent string
}

func main() {
	symbols := os.Args[1:]
	if len(symbols) == 0 {
		fmt.Printf("Usage stock [symbol ...]\n")
		return
	}
	resp, err := http.Get(fmt.Sprintf("http://api.addata.wallst.com/net/a5878a3d6f2be40db26311f6f8fb21a3/MiniQuotes/json?symbols=%s", strings.Join(symbols, "|")))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var quotes []Quote
	if json.Unmarshal(body, &quotes) != nil {
		log.Fatal(fmt.Sprintf("Couldn't decode json: '%s'", body))
	}
	for _, quote := range quotes {
		fmt.Printf("%s:\t%s\t%s\t(%s)\n", quote.Symbol, quote.Last, quote.Change, quote.ChangePercent)
	}
}
