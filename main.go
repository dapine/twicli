package main

import (
	"fmt"
	"log"
)

const (
	rootURL  = "https://api.twitch.tv/helix"
	clientId = "bhm2qd7pfnpy0l60y6rzlol6itbby7"
)

func main() {
	client := NewClient(clientId)

	// games, err := client.GetTopGames()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, g := range games.Data {
	// 	fmt.Println(g.Name)
	// }

	streams, err := client.GetTopStreams(40)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range streams.Data {
		fmt.Println(s)
	}
}
