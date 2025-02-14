package pokeapi

type LocationAreaListItem struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaList struct {
	Count    int                    `json:"count"`
	Next     string                 `json:"next"`
	Previous *string                `json:"previous"`
	Results  []LocationAreaListItem `json:"results"`
}

type LocationArea struct {
	Name    string             `json:"name"`
	Pokemon []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name           string `json:"name"`
	URL            string `json:"url"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []Stat `json:"stats"`
	Types          []Type `json:"types"`
}

type Stat struct {
	BaseStat int `json:"base_stat"`
	Stat     struct {
		Name string `json:"name"`
	} `json:"stat`
}

type Type struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}
