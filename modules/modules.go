package modules

import (
	"strings"
)

func FetchPokemonNames(pokemons string) (names []string) {
	for poke := pokemons; strings.Contains(poke, "<h2 class=\"woocommerce-loop-product__title\">"); _, poke, _ = strings.Cut(poke, "</h2>") {
		process, _, _ := strings.Cut(poke, "</h2>")

		_, name, _ := strings.Cut(process, "<h2 class=\"woocommerce-loop-product__title\">")
		names = append(names, name)
	}

	return
}
