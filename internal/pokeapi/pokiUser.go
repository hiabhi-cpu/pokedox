package pokeapi

type PokiUser struct {
	Pokemons  map[string]Pokemon
	Experince int
}

func NewUser() PokiUser {
	return PokiUser{
		Pokemons:  map[string]Pokemon{},
		Experince: 0,
	}
}
