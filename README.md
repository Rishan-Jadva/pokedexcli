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
