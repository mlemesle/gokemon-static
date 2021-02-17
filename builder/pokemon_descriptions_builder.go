package builder

import "github.com/mtslzr/pokeapi-go"

type pokemonDescriptionKeyS struct {
	GenerationName   string
	GenerationNameEN string
}

type PokemonDescriptionsS struct {
	DescriptionAndGameNameByGeneration map[pokemonDescriptionKeyS]map[string]string
}

func NewPokemonDescriptionsS() *PokemonDescriptionsS {
	return &PokemonDescriptionsS{
		DescriptionAndGameNameByGeneration: make(map[pokemonDescriptionKeyS]map[string]string),
	}
}

func (p *PokemonDescriptionsS) Build(pokemonName string) error {
	pokemon, err := pokeapi.Pokemon(pokemonName)
	if err != nil {
		return err
	}
	pokemonSpecie, err := pokeapi.PokemonSpecies(pokemon.Species.Name)
	if err != nil {
		return err
	}

	for _, flavorTextEntry := range pokemonSpecie.FlavorTextEntries {
		if flavorTextEntry.Language.Name == "en" {
			version, err := pokeapi.Version(flavorTextEntry.Version.Name)
			if err != nil {
				return err
			}
			versionGroup, err := pokeapi.VersionGroup(version.VersionGroup.Name)
			if err != nil {
				return err
			}
			generation, err := pokeapi.Generation(versionGroup.Generation.Name)
			if err != nil {
				return err
			}
			key := pokemonDescriptionKeyS{
				GenerationName:   generation.Name,
				GenerationNameEN: extractENName(generation.Names),
			}
			if _, ok := p.DescriptionAndGameNameByGeneration[key]; !ok {
				p.DescriptionAndGameNameByGeneration[key] = make(map[string]string)
			}
			p.DescriptionAndGameNameByGeneration[key][getVersionNameEN(version.Name, version.Names)] += " " + flavorTextEntry.FlavorText
		}
	}

	return nil
}
