package builder

import (
	"fmt"
	"github.com/mtslzr/pokeapi-go"
)

const officialArtworkFormat string = "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/%d.png"

type PokemonS struct {
	Name     string
	Order    int
	ImageURL string
	Types    []string
}

func (p *PokemonS) Build(pokemonName string) error {
	if err := p.fillPokemonFields(pokemonName); err != nil {
		return err
	}
	return nil
}

func (p *PokemonS) fillPokemonFields(pokemonName string) error {
	pokemon, err := pokeapi.Pokemon(pokemonName)
	if err != nil {
		return err
	}
	p.Name = pokemon.Name
	p.Order = pokemon.Order
	p.ImageURL = fmt.Sprintf(officialArtworkFormat, pokemon.ID)
	types := make([]string, len(pokemon.Types))
	for _, t := range pokemon.Types {
		types[t.Slot-1] = t.Type.Name
	}
	p.Types = types
	return nil
}
