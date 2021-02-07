package template

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	p "path"
	"strings"

	"github.com/markbates/pkger"
	"github.com/mlemesle/gokemon-static/builder"
)

var templates *template.Template = template.New("")

func InitTemplates(templatesDir string) error {
	return pkger.Walk(templatesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".gohtml") {
			panic(fmt.Errorf("%s is not a template file", path))
		}
		f, _ := pkger.Open(path)
		sl, _ := ioutil.ReadAll(f)
		pathParts := strings.Split(path, "/")
		filename := pathParts[len(pathParts)-1]
		templateName := strings.TrimSuffix(filename, p.Ext(filename))

		if _, err := templates.New(templateName).Parse(string(sl)); err != nil {
			return err
		}
		return nil
	})
}

func GenerateGokemon(file io.Writer, g builder.GokemonS) error {
	return templates.ExecuteTemplate(file, "gokemon", g)
}

func GeneratePokemon(file io.Writer, p builder.PokemonS) error {
	return templates.ExecuteTemplate(file, "pokemon", p)
}
