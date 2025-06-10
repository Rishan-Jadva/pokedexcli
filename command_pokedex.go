package main

import "fmt"

func commandPokedex(cfg *config, argument string) error {

	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("you have no pokemon in your pokedex")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
