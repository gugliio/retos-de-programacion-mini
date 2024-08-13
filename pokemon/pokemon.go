package pokemon

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

/*
LISTA POKÉMON
Lista los 151 primeros Pokémon
utilizando la PokéAPI.
*/

type data struct {
	Results []Pokemon `json:"results"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func Execute() {
	url := "https://pokeapi.co/api/v2/pokemon?limit=151"

	res, getErr := http.Get(url)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var parsedResponse data
	err := json.Unmarshal(body, &parsedResponse)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range parsedResponse.Results {
		format, _ := fmt.Printf("Pokemon: %s, URL: %s", v.Name, v.URL)
		fmt.Println(format)
	}
}
