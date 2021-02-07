package generator

import (
	"fmt"
	"log"
	"os"

	bld "github.com/mlemesle/gokemon-static/builder"
	tpl "github.com/mlemesle/gokemon-static/template"
)

const pokemonsDir string = "pokemons"
const pokemonFilepathFormat string = "/%s.html"

func GetPokemonDir(basePath string) string {
	return basePath + pokemonsDir
}

func generatePokemon(pokemonName, basePath string) (*bld.GokemonPartS, error) {
	pokemonS := bld.PokemonS{}
	pokemonS.Build(pokemonName)
	relativePokemonPath := fmt.Sprintf(pokemonsDir+pokemonFilepathFormat, pokemonName)
	pokemonPath := basePath + relativePokemonPath
	pokemonFile, err := os.Create(pokemonPath)
	if err != nil {
		log.Printf("Couldn't create file %s.\nError is : %v", pokemonPath, err)
	}
	defer pokemonFile.Close()
	if err := tpl.GeneratePokemon(pokemonFile, pokemonS); err != nil {
		log.Printf("Couldn't generate %s.\nError is : %v", pokemonPath, err)
	}
	return &bld.GokemonPartS{
		Order: pokemonS.PokemonIdentityCard.Order,
		Name:  pokemonS.Name,
		Path:  relativePokemonPath,
	}, nil
}
