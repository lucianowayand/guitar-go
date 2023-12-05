package main

func main() {
	Setup()

	state := Menu
	var song string
	for {
		if state == Menu {
			MenuScreen(&state)
		}
		if state == Playlist {
			PlaylistScreen(&state, &song)
		}
		if state == Playing {
			PlayingScreen(song, 30, &state)
		}

		if state == GameOver {
			GameOverScreen(&state)
		}
	}
}