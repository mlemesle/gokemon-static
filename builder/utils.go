package builder

import (
	"fmt"
	"strings"
)

type namedResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func extractLangName(lang string, names []struct {
	Language struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"language"`
	Name string `json:"name"`
}) string {
	for _, e := range names {
		if e.Language.Name == lang {
			return e.Name
		}
	}
	return ""
}

func extractENName(names []struct {
	Language struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"language"`
	Name string `json:"name"`
}) string {
	return extractLangName("en", names)
}

func extractFRName(names []struct {
	Language struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"language"`
	Name string `json:"name"`
}) string {
	return extractLangName("fr", names)
}

func getVersionNameEN(versionName string, names []struct {
	Language struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"language"`
	Name string `json:"name"`
}) string {
	result := extractENName(names)
	if result != "" {
		return result
	}
	switch versionName {
	case "red":
		return "Red"
	case "blue":
		return "Blue"
	case "yellow":
		return "Yellow"
	case "gold":
		return "Gold"
	case "silver":
		return "Silver"
	case "crystal":
		return "Crystal"
	case "ruby":
		return "Ruby"
	case "sapphire":
		return "Sapphire"
	case "emerald":
		return "Emerald"
	case "firered":
		return "Fire Red"
	case "leafgreen":
		return "Leaf Green"
	case "diamond":
		return "Diamond"
	case "pearl":
		return "Pearl"
	case "platinum":
		return "Platinum"
	case "heartgold":
		return "HeartGold"
	case "soulsilver":
		return "SoulSilver"
	case "black":
		return "Black"
	case "white":
		return "White"
	case "colosseum":
		return "Colosseum"
	case "xd":
		return "XD"
	case "black-2":
		return "Black 2"
	case "white-2":
		return "White 2"
	case "x":
		return "X"
	case "y":
		return "Y"
	case "omega-ruby":
		return "Omega Ruby"
	case "alpha-sapphire":
		return "Alpha Sapphire"
	case "sun":
		return "Sun"
	case "moon":
		return "Moon"
	case "ultra-sun":
		return "Ultra Sun"
	case "ultra-moon":
		return "Ultra Moon"
	case "lets-go-pikachu":
		return "Let's Go, Pikachu"
	case "lets-go-eevee":
		return "Let's Go, Eevee"
	case "sword":
		return "Sword"
	case "shield":
		return "Shield"
	default:
		return "Unknow version"
	}
}

func extractNameFromInterface(x interface{}) string {
	tmp := x.(map[string]interface{})
	return tmp["name"].(string)
}

func getGenderNameFromID(id string) string {
	return strings.Title(id)
}

func getIDFromURL(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-2]
}

const officialArtworkFormat string = "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/%d.png"

func getArtworkURLFromID(id int) string {
	return fmt.Sprintf(officialArtworkFormat, id)
}

const typeAssetFormat string = "/assets/images/types/%s.png"

func getTypePictoURL(typeName string) string {
	return fmt.Sprintf(typeAssetFormat, typeName)
}

func xOr1(x int) int {
	if x != 0 {
		return x
	}
	return 1
}

/*
func GetGenerationNameFR(generationName string) string {
	switch generationName {
	case "generation-i":
		return "Gén. I"
	case "generation-ii":
		return "Gén. II"
	case "generation-iii":
		return "Gén. III"
	case "generation-iv":
		return "Gén. IV"
	case "generation-v":
		return "Gén. V"
	case "generation-vi":
		return "Gén. VI"
	case "generation-vii":
		return "Gén. VII"
	case "generation-viii":
		return "Gén. VIII"
	default:
		return "Gén. inconnue"
	}
}
*/
