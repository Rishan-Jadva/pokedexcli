package main

import (
	"fmt"
)

func commandExplore(cfg *config, argument string) error {

	pokemonResp, err := cfg.pokeapiClient.ListPokemon(&argument)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", argument)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonResp.PokemonEncounters {
		fmt.Print(" - ")
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
