package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, argument string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", argument)

	pokemonResp, err := cfg.pokeapiClient.CatchPokemon(&argument)
	if err != nil {
		return err
	}

	baseChance := 0.8
	divisor := 608

	catchChance := baseChance - (float64(pokemonResp.BaseExperience) / float64(divisor))

	if catchChance < 0 {
		catchChance = 0
	}
	if catchChance > 1 {
		catchChance = 1
	}

	randomNumber := rand.Float64()

	if randomNumber < catchChance {
		fmt.Printf("%s was caught!\n", pokemonResp.Name)
		cfg.caughtPokemon[pokemonResp.Name] = pokemonResp
	} else {
		fmt.Printf("%s escaped!\n", pokemonResp.Name)
	}

	return nil
}
