package main

type State int

const (
	Menu State = iota
	Playing
	Results
)

func main() {
	Setup()

	state := Menu
	for {
		if state == Menu {
			MenuScreen(&state)
		}
		if state == Playing {
			PlayingScreen("little-lamb.sg", 20, &state)
		}

	}
}