package main

import (
	"log"

	cs "github.com/johnsonaj/concurrencystuff"
)

func main() {
	js := cs.New("https://api.yomomma.info/", "https://foaas.com/asshole")

	joke, err := js.GetJoke()
	if err != nil {
		panic(err)
	}

	fo, err := js.FuckOffAsshole("something")
	if err != nil {
		panic(err)
	}

	log.Println(joke)
	log.Println(fo)
}
