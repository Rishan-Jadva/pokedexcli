# PokedexCLI
A simple command-line Pokedex written in Go. It interacts with https://pokeapi.co/ to let you explore locations, find Pokemon and attempt to catch them - all from your terminal.


## Features
- Explore the Pokemon world by navigating using the paginated location.
- View the Pokemon available in a specific location.
- Attempt to catch wild Pokemon -- catch rate dependent on the Pokemon's base stats.
- Inspect the details of caught Pokemon.
- Maintain a Pokedex of caught Pokemon.
- Built-in help system.
- Caching system to reduce API calls.


## Example
```
📘 Pokedex > map
        canalave-city-area
        eterna-city-area
        pastoria-city-area
        sunyshore-city-area
        sinnoh-pokemon-league-area
        oreburgh-mine-1f
        oreburgh-mine-b1f
        valley-windworks-area
        eterna-forest-area
        fuego-ironworks-area
        mt-coronet-1f-route-207
        mt-coronet-2f
        mt-coronet-3f
        mt-coronet-exterior-snowfall
        mt-coronet-exterior-blizzard
        mt-coronet-4f
        mt-coronet-4f-small-room
        mt-coronet-5f
        mt-coronet-6f
        mt-coronet-1f-from-exterior
📘 Pokedex > map
        mt-coronet-1f-route-216
        mt-coronet-1f-route-211
        mt-coronet-b1f
        great-marsh-area-1
        great-marsh-area-2
        great-marsh-area-3
        great-marsh-area-4
        great-marsh-area-5
        great-marsh-area-6
        solaceon-ruins-2f
        solaceon-ruins-1f
        solaceon-ruins-b1f-a
        solaceon-ruins-b1f-b
        solaceon-ruins-b1f-c
        solaceon-ruins-b2f-a
        solaceon-ruins-b2f-b
        solaceon-ruins-b2f-c
        solaceon-ruins-b3f-a
        solaceon-ruins-b3f-b
        solaceon-ruins-b3f-c
📘 Pokedex > mapb
        canalave-city-area
        eterna-city-area
        pastoria-city-area
        sunyshore-city-area
        sinnoh-pokemon-league-area
        oreburgh-mine-1f
        oreburgh-mine-b1f
        valley-windworks-area
        eterna-forest-area
        fuego-ironworks-area
        mt-coronet-1f-route-207
        mt-coronet-2f
        mt-coronet-3f
        mt-coronet-exterior-snowfall
        mt-coronet-exterior-blizzard
        mt-coronet-4f
        mt-coronet-4f-small-room
        mt-coronet-5f
        mt-coronet-6f
        mt-coronet-1f-from-exterior
📘 Pokedex > explore mt-coronet-4f
        Exploring mt-coronet-4f...
        Found Pokemon:
        - clefairy
        - zubat
        - golbat
        - machoke
        - graveler
        - magikarp
        - gyarados
        - dratini
        - dragonair
        - nosepass
        - medicham
        - lunatone
        - solrock
        - barboach
        - whiscash
        - chingling
        - bronzong
📘 Pokedex > catch zubat
        Throwing a Pokeball at zubat...
        zubat was caught!
📘 Pokedex > inspect zubat
        Name: zubat
        Height: 8
        Weight: 75
        Stats:
         - hp: 40
         - attack: 45
         - defense: 35
         - special-attack: 30
         - special-defense: 40
         - speed: 55
        Types:
         - poison
         - flying
📘 Pokedex > pokedex
        Your Pokedex:
        - zubat
📘 Pokedex > exit
        Closing the Pokedex... Goodbye!
```


## Getting Started
### Prerequisites
* Go 1.20+
* Internet connection (for communicating with the API)

### Installation
```bash
git clone https://github.com/Rishan-Jadva/pokedexcli.git
cd pokedexcli
go build -o pokedexcli
./pokedexcli
```
