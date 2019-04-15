package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World from path: %s\n", r.URL.Path)
	})

	http.HandleFunc("/cards", func(w http.ResponseWriter, r *http.Request) {
		cards := newDeck()
		// hand, remainingCards := deal(cards, 5)

		// cards.saveToFile("Cards")
		// cards2 := newDeckFromFile("Cards")
		// cards2.shuffle()

		fmt.Fprintf(w, "Hello World from path: %s\n", cards)
	})

	http.ListenAndServe(":"+PORT, nil)
}
