package main

import (
	"log"

	cs "github.com/johnsonaj/concurrencystuff"
)

func main() {
	joke, err := cs.GetJoke()
	if err != nil {
		panic(err)
	}

	log.Println(joke)
}
