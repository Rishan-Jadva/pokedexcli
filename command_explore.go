package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, argument string) error {

	pokemonResp, err := cfg.pokeapiClient.ListPokemon(&argument)
	if err != nil {
		return errors.New("unknown location")
	}
	fmt.Printf("Exploring %s...\n", argument)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonResp.PokemonEncounters {
		fmt.Print(" - ")
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
