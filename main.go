package main

import (
	"SsmtpGUI/modules"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {

	maxpage := 48

	var allnames []string

	for index := 1; index < maxpage; index++ {

		// Fetch website to be scraped
		r, err := http.Get(fmt.Sprintf("https://scrapeme.live/shop/page/%d", index))
		if err != nil {
			log.Fatal(err)
		}

		// Read website body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer r.Body.Close()

		var bodystr string
		// Do processing of data
		bodystr = string(body)
		// Products under div "content" -> div "col-full" -> div "primary" -> main "main" -> ul "product columns-4" -> li "product"

		slaissi := strings.Split(bodystr, "<ul class=\"products columns-4\">")
		slaissi = strings.Split(slaissi[1], "<div class=\"storefront-sorting\">")
		allnames = append(allnames, modules.FetchPokemonNames(slaissi[0])...)
	}
	fmt.Println(allnames)
}
