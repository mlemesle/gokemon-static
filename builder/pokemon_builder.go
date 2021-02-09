package builder

import (
	"fmt"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

const officialArtworkFormat string = "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/%d.png"
const typeAssetFormat string = "/assets/images/types/%s.png"
const genderRateRatio float32 = 12.5

type GenderRates struct {
	FemaleRate float32
	MaleRate   float32
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
	GenderRates        *GenderRates
	Color              string
	CaptureRate        int
	ArtworkURL         string
}

type pokemonDescriptionKeyS struct {
	GenerationName   string
	GenerationNameEN string
}

type PokemonDescriptionsS struct {
	DescriptionAndGameNameByGeneration map[pokemonDescriptionKeyS]map[string]string
}

type PokemonS struct {
	Name                string
	PokemonCard         *PokemonCardS
	PokemonDescriptions *PokemonDescriptionsS
}

func (p *PokemonS) Build(pokemonName string) error {
	if err := p.fillCardFields(pokemonName); err != nil {
		return err
	}
	if err := p.fillDescriptionsFields(pokemonName); err != nil {
		return err
	}
	return nil
}

// PokemonCardS

func extractTypes(p *structs.Pokemon) []string {
	result := make([]string, len(p.Types))
	for i, t := range p.Types {
		result[i] = fmt.Sprintf(typeAssetFormat, t.Type.Name)
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

func extractGenderRates(p *structs.PokemonSpecies) *GenderRates {
	if p.GenderRate < -1 {
		return nil
	}
	femaleGenderRate := float32(p.GenderRate) * genderRateRatio
	return &GenderRates{
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

func (p *PokemonS) fillCardFields(pokemonName string) error {
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

	nameFR, nameEN, nameJP1, nameJP2 := "", "", "", ""
	for _, pokemonSpecieName := range pokemonSpecie.Names {
		switch pokemonSpecieName.Language.Name {
		case "fr":
			nameFR = pokemonSpecieName.Name
		case "en":
			nameEN = pokemonSpecieName.Name
		case "ja":
			nameJP1 = pokemonSpecieName.Name
		case "roomaji":
			nameJP2 = pokemonSpecieName.Name
		}
	}

	gender := ""
	for _, g := range pokemonSpecie.Genera {
		if g.Language.Name == "en" {
			gender = g.Genus
			break
		}
	}

	abilities, err := extractAbilities(&pokemon)
	if err != nil {
		return err
	}

	eggGroups, err := extractEggGroups(&pokemonSpecie)
	if err != nil {
		return err
	}

	effortPoints, err := extractEffortPoints(&pokemon)
	if err != nil {
		return err
	}

	level100Experience := -1
	for _, level := range growthRate.Levels {
		if level.Level == 100 {
			level100Experience = level.Experience
			break
		}
	}

	pokemonColor, err := extractColorName(&pokemonSpecie)
	if err != nil {
		return err
	}

	p.Name = nameEN
	p.PokemonCard = &PokemonCardS{
		Order:              pokemon.Order,
		NameFR:             nameFR,
		NameEN:             nameEN,
		NameJP:             nameJP1 + " " + nameJP2,
		Gender:             gender,
		Abilities:          abilities,
		EggGroups:          eggGroups,
		EffortPoints:       effortPoints,
		StepsUntilHatch:    computeStepsUntilHatch(&pokemonSpecie),
		BaseExperience:     pokemon.BaseExperience,
		Level100Experience: level100Experience,
		GenderRates:        extractGenderRates(&pokemonSpecie),
		Color:              pokemonColor,
		CaptureRate:        pokemonSpecie.CaptureRate,
		ArtworkURL:         fmt.Sprintf(officialArtworkFormat, pokemon.ID),
		Height:             float32(pokemon.Height) / 10,
		Weight:             float32(pokemon.Weight) / 10,
		Types:              extractTypes(&pokemon),
	}
	return nil
}

// PokemonDescriptions
func (p *PokemonS) fillDescriptionsFields(pokemonName string) error {
	pokemon, err := pokeapi.Pokemon(pokemonName)
	if err != nil {
		return err
	}
	pokemonSpecie, err := pokeapi.PokemonSpecies(pokemon.Species.Name)
	if err != nil {
		return err
	}

	descriptionAndGameNameByGeneration := make(map[pokemonDescriptionKeyS]map[string]string)

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
			if _, ok := descriptionAndGameNameByGeneration[key]; !ok {
				descriptionAndGameNameByGeneration[key] = make(map[string]string)
			}
			descriptionAndGameNameByGeneration[key][getVersionNameEN(version.Name, version.Names)] += " " + flavorTextEntry.FlavorText
		}
	}

	p.PokemonDescriptions = &PokemonDescriptionsS{
		DescriptionAndGameNameByGeneration: descriptionAndGameNameByGeneration,
	}
	return nil
}
