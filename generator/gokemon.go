package generator

import (
	"fmt"
	"os"

	bld "github.com/mlemesle/gokemon-static/builder"
	tpl "github.com/mlemesle/gokemon-static/template"
	"github.com/mtslzr/pokeapi-go"
)

const gokemonFilepath string = "/gokemon.html"

func getAllPokemonNames() ([]string, error) {
	var pokemonNames []string
	limit, offset := 250, 0
	for {
		resource, err := pokeapi.Resource("pokemon", offset, limit)
		if err != nil {
			return nil, err
		}
		for _, pokemon := range resource.Results {
			pokemonNames = append(pokemonNames, pokemon.Name)
		}
		if resource.Count == len(pokemonNames) {
			break
		}
		offset += limit
	}
	return pokemonNames, nil
}

func GetGokemonDir(basePath string) string {
	return basePath
}

func GenerateGokemon(basePath string) error {
	pokemonNames, err := getAllPokemonNames()
	if err != nil {
		return err
	}
	var pokemons []bld.GokemonPartS
	nbPokemons := len(pokemonNames)
	fmt.Println(fmt.Sprintf("Preparing to export %d pokemons", nbPokemons))
	for i, pokemonName := range pokemonNames {
		fmt.Println(fmt.Sprintf("%d/%d", i+1, nbPokemons))
		gokemonPartS, err := generatePokemon(pokemonName, basePath)
		if err != nil {
			return err
		}
		pokemons = append(pokemons, *gokemonPartS)
	}
	gokemonFile, err := os.Create(GetGokemonDir(basePath) + gokemonFilepath)
	if err != nil {
		return err
	}
	defer gokemonFile.Close()
	gokemon := &bld.GokemonS{}
	gokemon.Build(pokemons)
	return tpl.GenerateGokemon(gokemonFile, *gokemon)
}
