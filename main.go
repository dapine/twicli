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
		}).SetSelectedFocusOnly(true)

	streamList := tview.NewList().SetSelectedFocusOnly(true)

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

	streamList.AddItem("Back", "", 'b', func() { app.SetFocus(gameList) })

	updatePageStreamList(streams, streamList)

	lastPage := streams.Pagination.Cursor

	streamList.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		fmt.Println("", shortcut)

		if index+1 >= streamList.GetItemCount() {
			strs, err := client.PaginateStreams(id, "", lastPage)
			if err != nil {
				log.Fatal(err)
			}

			streams.Data = append(streams.Data, strs.Data...)
			// streams.Data = strs.Data
			lastPage = strs.Pagination.Cursor

			updatePageStreamList(streams, streamList)
		}
	})

	app.SetFocus(streamList)
}

func updatePageStreamList(streams *Streams, streamList *tview.List) {
	// By default, twitch API loads 20 streams on a page, so add just the last 20 streams to the UI list
	for _, s := range streams.Data[len(streams.Data)-20:] {
		title := fmt.Sprintf("[%d viewers] %s", int(s.ViewerCount), s.Title)
		streamList.AddItem(title, s.UserName, 0, func() {
			for _, d := range streams.Data {
				fmt.Println(d.Title)
			}

			username := streams.Data[streamList.GetCurrentItem()-1].UserName
			launchPlayer(player, "https://twitch.tv/"+username)
		})
	}
}

func launchPlayer(videoPlayer string, streamLink string) {
	cmd := exec.Command(videoPlayer, streamLink)
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
}
