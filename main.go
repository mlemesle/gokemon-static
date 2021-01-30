package main

import (
	"fmt"
	"github.com/markbates/pkger"
	"github.com/mlemesle/gokemon-static/generator"
	"github.com/mlemesle/gokemon-static/template"
	"os"
)

const templatesDir string = "/template/files"

const basePath string = "gokemon/"

func createGokemonDirectories() {
	gokemonDir := generator.GetGokemonDir(basePath)
	if _, err := os.Stat(gokemonDir); os.IsNotExist(err) {
		os.MkdirAll(gokemonDir, 0755)
	}
	pokemonDir := generator.GetPokemonDir(basePath)
	if _, err := os.Stat(pokemonDir); os.IsNotExist(err) {
		os.MkdirAll(pokemonDir, 0755)
	}
}

func main() {
	createGokemonDirectories()
	pkger.Include(templatesDir)
	template.InitTemplates(templatesDir)
	err := generator.GenerateGokemon(basePath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done")
}
