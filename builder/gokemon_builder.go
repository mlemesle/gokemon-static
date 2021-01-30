package builder

type GokemonPartS struct {
	Order int
	Name  string
	Path  string
}

type GokemonS struct {
	Pokemons []GokemonPartS
}

func (g *GokemonS) Build(pokemons []GokemonPartS) {
	g.Pokemons = pokemons
}
