package main

import (
	"SsmtpGUI/modules"
	"SsmtpGUI/pokemonproduct"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {

	// maxpage currently hardcorded, later make function to fetch maxpages
	maxpage := 10

	// variables to hold raw data
	var allnames []string
	var allprices []string

	for index := 1; index < maxpage; index++ {

		// Fetch website to be scraped for page [index]
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

		// Do processing of data
		bodystr := string(body)

		// process unnecessary data out, then pass it down to functions to get wanted data
		slaissi := strings.Split(bodystr, "<ul class=\"products columns-4\">")
		slaissi = strings.Split(slaissi[1], "<div class=\"storefront-sorting\">")

		allnames = append(allnames, modules.FetchPokemonNames(slaissi[0])...)
		allprices = append(allprices, modules.FetchPokemonPrice(slaissi[0])...)
	}

	// create empty struct to hold all the data
	pokemons := []pokemonproduct.PokemonProduct{}

	// assign fetched data to our own struct
	for i, _ := range allnames {
		pokemon := pokemonproduct.PokemonProduct{}
		pokemon.AddPokemon("TODO", allnames[i], allprices[i])
		pokemons = append(pokemons, pokemon)
	}

	// print fetched data
	fmt.Println(pokemons)
}
