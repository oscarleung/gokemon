package gokemon

//Pokedex -
type Pokedex struct {
	Created  string
	Modified string
	Name     string
	Pokemon  []Pokemon
}

//Pokemon -
type Pokemon struct {
	Name         string
	Resource_URI string
}
