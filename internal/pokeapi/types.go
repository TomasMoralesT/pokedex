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
	Name string `json:"name"`
	URL  string `json:"url"`
}
