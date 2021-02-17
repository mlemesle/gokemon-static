package builder

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mtslzr/pokeapi-go"
)

type PokemonEvolutionGridS struct {
	Grid [][]*PokemonEvolutionCellS
}

type PokemonEvolutionCellS struct {
	Name            string
	ArtworkURL      string
	EvolutionMethod string
	RowSpan         int
}

func NewPokemonEvolutionGridS() *PokemonEvolutionGridS {
	return &PokemonEvolutionGridS{
		Grid: nil,
	}
}

func NewPokemonEvolutionGridSWithSize(nbRow, nbCol int) *PokemonEvolutionGridS {
	grid := make([][]*PokemonEvolutionCellS, nbRow)
	for i := range grid {
		grid[i] = make([]*PokemonEvolutionCellS, nbCol)
	}
	return &PokemonEvolutionGridS{
		Grid: grid,
	}
}

func NewPokemonEvolutionCellS() *PokemonEvolutionCellS {
	return &PokemonEvolutionCellS{
		Name:            "",
		ArtworkURL:      "",
		EvolutionMethod: "",
		RowSpan:         1,
	}
}

func (p *PokemonEvolutionCellS) Build(pokemonName string, evolutionDetails []struct {
	Gender                interface{} `json:"gender"`
	HeldItem              interface{} `json:"held_item"`
	Item                  interface{} `json:"item"`
	KnownMove             interface{} `json:"known_move"`
	KnownMoveType         interface{} `json:"known_move_type"`
	Location              interface{} `json:"location"`
	MinAffection          interface{} `json:"min_affection"`
	MinBeauty             interface{} `json:"min_beauty"`
	MinHappiness          interface{} `json:"min_happiness"`
	MinLevel              int         `json:"min_level"`
	NeedsOverworldRain    bool        `json:"needs_overworld_rain"`
	PartySpecies          interface{} `json:"party_species"`
	PartyType             interface{} `json:"party_type"`
	RelativePhysicalStats interface{} `json:"relative_physical_stats"`
	TimeOfDay             string      `json:"time_of_day"`
	TradeSpecies          interface{} `json:"trade_species"`
	Trigger               struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"trigger"`
	TurnUpsideDown bool `json:"turn_upside_down"`
}) error {
	pokemon, err := pokeapi.Pokemon(pokemonName)
	if err != nil {
		return err
	}
	pokemonSpecie, err := pokeapi.PokemonSpecies(pokemonName)
	if err != nil {
		return err
	}
	var evolutionTriggerParts []string
	for _, evolutionDetail := range evolutionDetails {
		if evolutionDetail.Gender != nil {
			genderID := strconv.Itoa(int(evolutionDetail.Gender.(float64)))
			evolutionTriggerParts = append(evolutionTriggerParts, getGenderNameFromID(genderID))
		}
		if evolutionDetail.HeldItem != nil {
			heldItemName := extractNameFromInterface(evolutionDetail.HeldItem)
			item, err := pokeapi.Item(heldItemName)
			if err != nil {
				return err
			}
			evolutionTriggerParts = append(evolutionTriggerParts, fmt.Sprintf("Hold %s", extractENName(item.Names)))
		}
		if evolutionDetail.Item != nil {
			itemName := extractNameFromInterface(evolutionDetail.Item)
			item, err := pokeapi.Item(itemName)
			if err != nil {
				return err
			}
			evolutionTriggerParts = append(evolutionTriggerParts, fmt.Sprintf("Use %s", extractENName(item.Names)))
		}
		if evolutionDetail.KnownMove != nil {
			knownMoveName := extractNameFromInterface(evolutionDetail.KnownMove)
			move, err := pokeapi.Move(knownMoveName)
			if err != nil {
				return err
			}
			evolutionTriggerParts = append(evolutionTriggerParts, fmt.Sprintf("Knowing %s", extractENName(move.Names)))
		}
		if evolutionDetail.KnownMoveType != nil {
			knownMoveTypeName := extractNameFromInterface(evolutionDetail.KnownMoveType)
			pType, err := pokeapi.Type(knownMoveTypeName)
			if err != nil {
				return err
			}
			evolutionTriggerParts = append(evolutionTriggerParts, fmt.Sprintf("Knowing attack of type %s", extractENName(pType.Names)))
		}
		if evolutionDetail.Location != nil {
			locationName := extractNameFromInterface(evolutionDetail.Location)
			location, err := pokeapi.Location(locationName)
			if err != nil {
				return nil
			}
			evolutionTriggerParts = append(evolutionTriggerParts, fmt.Sprintf("At %s", extractENName(location.Names)))
		}
		if evolutionDetail.MinAffection != nil {
			affection := int(evolutionDetail.MinAffection.(float64))
			evolutionTriggerParts = append(evolutionTriggerParts, fmt.Sprintf("Affection level %d", affection))
		}
		if evolutionDetail.MinBeauty != nil {
			beauty := int(evolutionDetail.MinBeauty.(float64))
			evolutionTriggerParts = append(evolutionTriggerParts, fmt.Sprintf("Beauty level %d", beauty))
		}
		if evolutionDetail.MinHappiness != nil {
			happiness := int(evolutionDetail.MinHappiness.(float64))
			evolutionTriggerParts = append(evolutionTriggerParts, fmt.Sprintf("Happiness level %d", happiness))
		}
		if evolutionDetail.MinLevel > 0 {
			evolutionTriggerParts = append(evolutionTriggerParts, fmt.Sprintf("Level %d", evolutionDetail.MinLevel))
		}
		if evolutionDetail.NeedsOverworldRain {
			evolutionTriggerParts = append(evolutionTriggerParts, "Need rain in Overworld")
		}
		if evolutionDetail.PartySpecies != nil {
			pokemonName := extractNameFromInterface(evolutionDetail.PartySpecies)
			specie, err := pokeapi.PokemonSpecies(pokemonName)
			if err != nil {
				return err
			}
			evolutionTriggerParts = append(evolutionTriggerParts, fmt.Sprintf("Need %s in team", extractENName(specie.Names)))
		}
		if evolutionDetail.PartyType != nil {
			typeName := extractNameFromInterface(evolutionDetail.PartyType)
			pType, err := pokeapi.Type(typeName)
			if err != nil {
				return err
			}
			evolutionTriggerParts = append(evolutionTriggerParts, fmt.Sprintf("Need pokemon of type %s in team", extractENName(pType.Names)))
		}
		if evolutionDetail.RelativePhysicalStats != nil {
			relativePhysicalStats := int(evolutionDetail.RelativePhysicalStats.(float64))
			if relativePhysicalStats > 0 {
				evolutionTriggerParts = append(evolutionTriggerParts, "Attack > Defense")
			} else if relativePhysicalStats < 0 {
				evolutionTriggerParts = append(evolutionTriggerParts, "Attack < Defense")
			} else {
				evolutionTriggerParts = append(evolutionTriggerParts, "Attack = Defense")
			}
		}
		if evolutionDetail.TimeOfDay != "" {
			evolutionTriggerParts = append(evolutionTriggerParts, strings.Title(evolutionDetail.TimeOfDay))
		}
		if evolutionDetail.TradeSpecies != nil {
			tradeSpecie := extractNameFromInterface(evolutionDetail.TradeSpecies)
			specie, err := pokeapi.PokemonSpecies(tradeSpecie)
			if err != nil {
				return err
			}
			evolutionTriggerParts = append(evolutionTriggerParts, fmt.Sprintf("Trade with %s", extractENName(specie.Names)))
		}
		if evolutionDetail.TurnUpsideDown {
			evolutionTriggerParts = append(evolutionTriggerParts, "Device upside down")
		}
	}
	p.Name = extractENName(pokemonSpecie.Names)
	p.ArtworkURL = getArtworkURLFromID(pokemon.ID)
	p.EvolutionMethod = strings.Join(evolutionTriggerParts, " / ")
	return nil
}

func (p *PokemonEvolutionGridS) Build(pokemonName string) error {
	pokemonSpecie, err := pokeapi.PokemonSpecies(pokemonName)
	if err != nil {
		return err
	}
	evolutionChain, err := pokeapi.EvolutionChain(getIDFromURL(pokemonSpecie.EvolutionChain.URL))
	if err != nil {
		return err
	}

	totalSubEvolutions := 0
	for _, d1 := range evolutionChain.Chain.EvolvesTo {
		totalSubEvolutions += xOr1(len(d1.EvolvesTo))
	}
	totalSubEvolutions = xOr1(totalSubEvolutions)

	*p = *NewPokemonEvolutionGridSWithSize(totalSubEvolutions, 3)

	currentDepth1, currentDepth2 := 0, 0
	for _, evolveDepth1 := range evolutionChain.Chain.EvolvesTo {
		targetDepth1 := NewPokemonEvolutionCellS()
		if err := targetDepth1.Build(evolveDepth1.Species.Name, evolveDepth1.EvolutionDetails); err != nil {
			return err
		}
		targetDepth1.RowSpan = xOr1(len(evolveDepth1.EvolvesTo))
		p.Grid[currentDepth1][1] = targetDepth1
		currentDepth1++

		for _, evolveDepth2 := range evolveDepth1.EvolvesTo {
			targetDepth2 := NewPokemonEvolutionCellS()
			if err = targetDepth2.Build(evolveDepth2.Species.Name, evolveDepth2.EvolutionDetails); err != nil {
				return err
			}
			p.Grid[currentDepth2][2] = targetDepth2
			currentDepth2++
		}
	}
	baseEvolution := NewPokemonEvolutionCellS()
	baseEvolution.Build(pokemonName, nil)
	baseEvolution.RowSpan = totalSubEvolutions
	p.Grid[0][0] = baseEvolution
	return nil
}
