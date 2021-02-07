package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/markbates/pkger"
	"github.com/mlemesle/gokemon-static/generator"
	"github.com/mlemesle/gokemon-static/template"
)

const templatesDir string = "/template/files"
const assetsDir string = "/assets"

const basePath string = "gokemon/"

func copyDir(src string, dst string) error {
	return pkger.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		filepath := strings.Split(path, ":")[1]
		if info.IsDir() {
			if err := os.MkdirAll("gokemon"+filepath, 0755); err != nil {
				return err
			}
		} else {
			srcFile, err := pkger.Open(path)
			if err != nil {
				return err
			}
			defer srcFile.Close()
			dstFile, err := os.Create("gokemon" + filepath)
			if err != nil {
				return err
			}
			defer dstFile.Close()
			if _, err := io.Copy(dstFile, srcFile); err != nil {
				return err
			}
		}
		return nil
	})
}

func createGokemonDirectories() {
	fmt.Println("Creating all necessary directories...")
	gokemonDir := generator.GetGokemonDir(basePath)
	if _, err := os.Stat(gokemonDir); os.IsNotExist(err) {
		os.MkdirAll(gokemonDir, 0755)
	}
	pokemonDir := generator.GetPokemonDir(basePath)
	if _, err := os.Stat(pokemonDir); os.IsNotExist(err) {
		os.MkdirAll(pokemonDir, 0755)
	}
	if err := copyDir(assetsDir, gokemonDir+"assets"); err != nil {
		panic(err)
	}
	fmt.Println("Directory tree created !")
}

func main() {
	pkger.Include(templatesDir)
	pkger.Include(assetsDir)
	createGokemonDirectories()
	if err := template.InitTemplates(templatesDir); err != nil {
		panic(err)
	}
	if err := generator.GenerateGokemon(basePath); err != nil {
		panic(err)
	}
	fs := http.FileServer(http.Dir("./gokemon"))
	http.Handle("/", fs)

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
