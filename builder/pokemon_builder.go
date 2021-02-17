package builder

type PokemonS struct {
	Name                  string
	PokemonCard           *PokemonCardS
	PokemonDescriptions   *PokemonDescriptionsS
	PokemonEvolutionGridS *PokemonEvolutionGridS
}

func NewPokemonS() *PokemonS {
	return &PokemonS{
		Name:                  "",
		PokemonCard:           NewPokemonCardS(),
		PokemonDescriptions:   NewPokemonDescriptionsS(),
		PokemonEvolutionGridS: NewPokemonEvolutionGridS(),
	}
}

func (p *PokemonS) Build(pokemonName string) error {
	if err := p.PokemonCard.Build(pokemonName); err != nil {
		return err
	}
	if err := p.PokemonDescriptions.Build(pokemonName); err != nil {
		return err
	}
	if err := p.PokemonEvolutionGridS.Build(pokemonName); err != nil {
		return err
	}
	p.Name = p.PokemonCard.NameEN
	return nil
}
