package main

type Note struct {
	Position float32
	Lane     int
	Pressed  bool
}

type State int

const (
	Menu State = iota
	Playlist
	Playing
	Results
	GameOver
)

type MenuOptions int

const (
	Play MenuOptions = iota
	// Options
	Exit
)