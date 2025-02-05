
# Pokedox

Pokedox is a Command-Line Interface (CLI) tool designed to help users search for Pokémon characters efficiently. With Pokedox, you can retrieve detailed information about various Pokémon directly from your terminal. You can also retrieve the pokemon's present in a particular area and catch them.


## Features

- Explore Locations
- Get pokemon's in a particular location
- Catch a pokemon
- Get pokemon's attributes



## Installation without docker

To install Pokedox, ensure you have Go installed on your system. Then, follow these steps:

#### 1. Clone the Repository:
```bash
  git clone https://github.com/hiabhi-cpu/pokedox.git
  cd pokedox
```

#### 2. Build the Application:
```bash
  go build -o pokedox
```

#### 3. Run the Application:
```bash
  ./pokedox
```


## Installation with docker 

#### 1. Clone the project

```bash
  git clone https://github.com/hiabhi-cpu/pokedox.git
  cd pokedox
```

#### 2. Build docker image

```bash
  docker build -t pokedox .
```  
#### 3. Run the container

```bash
  docker run --rm -it pokedox
```    

## Usage
Once the application is running, you can use the following commands:

#### 1. Help :

```bash
help
```
To view all the commands present.

#### 2. View Map:

```bash
map
```
View the locations present in the pokemon wolrd. This command will give next 20 locations.

#### 3. Map back:

```bash
mapb
```
To view the previous 20 locations.

#### 4. Explore Pokémon by Type:

```bash
explore <location>
```
Replace <location> with the location that u find using map commands. This will give the list of pokemon's present in that location.

#### 5. Catch a Pokémon:

```bash
catch <pokemon_name>
```
Replace <pokemon_name> with the name of the Pokémon you want to catch. Based on the user experience and the pokemon's base experience you might catch the pokemon.

#### 6. Inspect a Pokémon:

```bash
inspect <pokemon_name>
```
Replace <pokemon_name> with the pokemon's name that you want to view its STAT.

#### 7. View all the Pokémon you caught:

```bash
pokedox
```
Gives names of all Pokémons you have caught.


## Contributing

Contributions are welcome! If you have suggestions for improvements or new features, feel free to open an issue or submit a pull request.

