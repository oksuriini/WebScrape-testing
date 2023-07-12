package pokemonproduct

type PokemonProduct struct {
	link  string
	name  string
	price string
}

func (p *PokemonProduct) AddPokemon(link, name, price string) {
	p.link = link
	p.name = name
	p.price = price
}
