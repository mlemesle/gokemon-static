package builder

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

func extractFRName(names []struct {
	Language struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"language"`
	Name string `json:"name"`
}) string {
	return extractLangName("fr", names)
}
