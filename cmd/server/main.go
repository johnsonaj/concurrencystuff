package main

import (
	"log"
	"sync"

	cs "github.com/johnsonaj/concurrencystuff"
)

func main() {
	js := cs.New("https://api.yomomma.info/", "https://foaas.com/asshole")

	response := make(chan interface{})
	errChan := make(chan error)
	var wg sync.WaitGroup

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(index int) {
			defer wg.Done()

			log.Printf("thread %d started", index)
			joke, err := js.GetJoke()
			if err != nil {
				// you can always log your erros in a channel also. This way
				// they are recorded but the program keeps running
				errChan <- err
			} else {
				response <- joke
				log.Printf("thread %d finished", index)
			}
		}(i)
	}

	go func() {
		for v := range response {
			log.Println(v)
		}
	}()

	wg.Wait()
}
