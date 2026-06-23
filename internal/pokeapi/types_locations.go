package pokeapi

type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type NamedAPIResourceList struct {
	Count    int                `json:"count"`
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

type LocationArea struct {
	Id                   int                   `json:"id"`
	Name                 string                `json:"name"`
	GameIndex            int                   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             NamedAPIResource      `json:"location"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource `json:"encounter_method"`
	VersionDetails  EncounterVersionDetails
}

type Name struct {
	Name     string
	Language NamedAPIResource
}

type PokemonEncounter struct {
	Pokemon        NamedAPIResource         `json:"pokemon"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

type EncounterVersionDetails struct {
	Rate    int              `json:"rate"`
	Version NamedAPIResource `json:"version"`
}

type VersionEncounterDetail struct {
	Version          NamedAPIResource
	MaxChance        int
	EncounterDetails []Encounter
}

type Encounter struct {
	MinLevel        int              `json:"min_level"`
	MaxLevel        int              `json:"max_level"`
	ConditionValues NamedAPIResource `json:"condition_values"`
	Chance          int              `json:"chance"`
	Method          NamedAPIResource `json:"method"`
}
