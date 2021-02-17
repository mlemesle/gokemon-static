package builder

import (
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

const genderRateRatio float32 = 12.5

type GenderRatesS struct {
	FemaleRate float32
	MaleRate   float32
}

func NewGenderRatesS() *GenderRatesS {
	return &GenderRatesS{
		FemaleRate: 0.0,
		MaleRate:   0.0,
	}
}

type PokemonCardS struct {
	Order              int
	NameFR             string
	NameEN             string
	NameJP             string
	Types              []string
	Gender             string
	Weight             float32
	Height             float32
	Abilities          []string
	EggGroups          []string
	StepsUntilHatch    int
	EffortPoints       map[string]int
	BaseExperience     int
	Level100Experience int
	GenderRates        *GenderRatesS
	Color              string
	CaptureRate        int
	ArtworkURL         string
}

func NewPokemonCardS() *PokemonCardS {
	return &PokemonCardS{
		Order:              0,
		NameFR:             "",
		NameEN:             "",
		NameJP:             "",
		Types:              []string{},
		Gender:             "",
		Weight:             0.0,
		Height:             0.0,
		Abilities:          []string{},
		EggGroups:          []string{},
		StepsUntilHatch:    0,
		EffortPoints:       make(map[string]int),
		BaseExperience:     0,
		Level100Experience: 0,
		GenderRates:        NewGenderRatesS(),
		Color:              "",
		CaptureRate:        0,
		ArtworkURL:         "",
	}
}

func extractTypes(p *structs.Pokemon) []string {
	result := make([]string, len(p.Types))
	for i, t := range p.Types {
		result[i] = getTypePictoURL(t.Type.Name)
	}
	return result
}

func extractAbilities(p *structs.Pokemon) ([]string, error) {
	result := make([]string, len(p.Abilities))
	for i, a := range p.Abilities {
		ability, err := pokeapi.Ability(a.Ability.Name)
		if err != nil {
			return nil, err
		}
		result[i] = extractENName(ability.Names)
	}
	return result, nil
}

func extractEggGroups(p *structs.PokemonSpecies) ([]string, error) {
	result := make([]string, len(p.EggGroups))
	for i, e := range p.EggGroups {
		eggGroup, err := pokeapi.EggGroup(e.Name)
		if err != nil {
			return nil, err
		}
		result[i] = extractENName(eggGroup.Names)
	}
	return result, nil
}

func extractEffortPoints(p *structs.Pokemon) (map[string]int, error) {
	result := make(map[string]int)
	for _, s := range p.Stats {
		if s.Effort != 0 {
			stat, err := pokeapi.Stat(s.Stat.Name)
			if err != nil {
				return nil, err
			}
			result[extractENName(stat.Names)] = s.Effort
		}
	}
	return result, nil
}

func computeStepsUntilHatch(p *structs.PokemonSpecies) int {
	return (p.HatchCounter + 1) * 255
}

func extractGenderRates(p *structs.PokemonSpecies) *GenderRatesS {
	if p.GenderRate < -1 {
		return nil
	}
	femaleGenderRate := float32(p.GenderRate) * genderRateRatio
	return &GenderRatesS{
		FemaleRate: femaleGenderRate,
		MaleRate:   100 - femaleGenderRate,
	}
}

func extractColorName(p *structs.PokemonSpecies) (string, error) {
	pokemonColor, err := pokeapi.PokemonColor(p.Color.Name)
	if err != nil {
		return "", err
	}
	return extractENName(pokemonColor.Names), nil
}

func (p *PokemonCardS) Build(pokemonName string) error {
	pokemon, err := pokeapi.Pokemon(pokemonName)
	if err != nil {
		return err
	}
	pokemonSpecie, err := pokeapi.PokemonSpecies(pokemon.Species.Name)
	if err != nil {
		return err
	}
	growthRate, err := pokeapi.GrowthRate(pokemonSpecie.GrowthRate.Name)
	if err != nil {
		return err
	}

	nameJP1, nameJP2 := "", ""
	for _, pokemonSpecieName := range pokemonSpecie.Names {
		switch pokemonSpecieName.Language.Name {
		case "fr":
			p.NameFR = pokemonSpecieName.Name
		case "en":
			p.NameEN = pokemonSpecieName.Name
		case "ja":
			nameJP1 = pokemonSpecieName.Name
		case "roomaji":
			nameJP2 = pokemonSpecieName.Name
		}
	}
	p.NameJP = nameJP1 + " " + nameJP2

	for _, g := range pokemonSpecie.Genera {
		if g.Language.Name == "en" {
			p.Gender = g.Genus
			break
		}
	}

	p.Abilities, err = extractAbilities(&pokemon)
	if err != nil {
		return err
	}

	p.EggGroups, err = extractEggGroups(&pokemonSpecie)
	if err != nil {
		return err
	}

	p.EffortPoints, err = extractEffortPoints(&pokemon)
	if err != nil {
		return err
	}

	for _, level := range growthRate.Levels {
		if level.Level == 100 {
			p.Level100Experience = level.Experience
			break
		}
	}

	p.Color, err = extractColorName(&pokemonSpecie)
	if err != nil {
		return err
	}

	p.Order = pokemon.Order
	p.StepsUntilHatch = computeStepsUntilHatch(&pokemonSpecie)
	p.BaseExperience = pokemon.BaseExperience
	p.GenderRates = extractGenderRates(&pokemonSpecie)
	p.CaptureRate = pokemonSpecie.CaptureRate
	p.ArtworkURL = getArtworkURLFromID(pokemon.ID)
	p.Height = float32(pokemon.Height) / 10
	p.Weight = float32(pokemon.Weight) / 10
	p.Types = extractTypes(&pokemon)

	return nil
}
