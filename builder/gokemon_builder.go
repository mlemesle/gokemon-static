package builder

type GokemonPartS struct {
	Order int
	Name  string
	Path  string
}

type GokemonS struct {
	Pokemons []*GokemonPartS
}

func NewGokemonS() *GokemonS {
	return &GokemonS{}
}

func (g *GokemonS) Build(pokemons []*GokemonPartS) {
	g.Pokemons = pokemons
}
