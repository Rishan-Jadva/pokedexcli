package main

import "fmt"

func commandInspect(cfg *config, argument string) error {

	pokemon, ok := cfg.caughtPokemon[argument]

	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, poketype := range pokemon.Types {
		fmt.Printf("  - %s\n", poketype.Type.Name)
	}

	return nil
}
