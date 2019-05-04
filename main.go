package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/rivo/tview"
)

const (
	rootURL  = "https://api.twitch.tv/helix"
	clientId = "bhm2qd7pfnpy0l60y6rzlol6itbby7"
	player   = "mpv"
)

func main() {
	client := NewClient(clientId)

	games, err := client.GetTopGames()
	if err != nil {
		log.Fatal(err)
	}

	app := tview.NewApplication()

	gameList := tview.NewList().
		ShowSecondaryText(false).
		AddItem("Quit", "", 'q', func() {
			app.Stop()
		})

	streamList := tview.NewList()

	for _, g := range games.Data {
		gameList.AddItem(g.Name, "", 0, func() {
			id := games.Data[gameList.GetCurrentItem()-1].Id
			selectGame(client, id, streamList, gameList, app)
		})
	}

	grid := tview.NewGrid().
		SetRows(0).
		SetColumns(50, 50).
		SetBorders(true).
		AddItem(gameList, 0, 0, 1, 1, 0, 0, false).
		AddItem(streamList, 0, 1, 1, 2, 0, 0, false)

	if err := app.SetRoot(grid, true).SetFocus(gameList).Run(); err != nil {
		panic(err)
	}
}

func selectGame(client Client, id string, streamList *tview.List, gameList *tview.List, app *tview.Application) {
	streams, err := client.FromGameId(id)
	if err != nil {
		log.Fatal(err)
	}

	streamList.Clear()

	streamList.AddItem("Back", "Back to game list", 'b', func() { app.SetFocus(gameList) })

	for _, s := range streams.Data {
		unf := fmt.Sprintf("\t%s - %d viewers", s.UserName, int(s.ViewerCount))
		streamList.AddItem(s.Title, unf, 0, func() {
			username := streams.Data[streamList.GetCurrentItem()-1].UserName
			launchPlayer(player, "https://twitch.tv/"+username)
		})
	}

	app.SetFocus(streamList)
}

func launchPlayer(videoPlayer string, streamLink string) {
	cmd := exec.Command(videoPlayer, streamLink)
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
}
