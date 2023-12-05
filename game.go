package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const screenHeight = 720
const screenWidth = 1280
var notes = []string{"C", "D", "E", "F", "G"}

func ReadSongFromFile(path string) [][]Note{
	f, err := os.Open(fmt.Sprintf("songs/"+path))
    if err != nil {
        panic(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)

	notes := [][]Note{} 

    for scanner.Scan() {
       strNotes := strings.Split(scanner.Text(), " ")
	   
	   elements := []Note{}
	   for _, strNote := range strNotes {
			if strNote == "" {
				continue
			}
			i, err := strconv.Atoi(strNote)
			if err != nil {
				panic(err)
			}
			elements = append(elements, Note{Position: -10, Lane: i} )
	   }
	   notes = append(notes, elements)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

	return notes
}

func CenteredTextPosX(text string, fontSize int32) int32 {
	return int32(screenWidth/2-(len(text)*int(fontSize/4)))
}

func Setup() {
	rl.InitWindow(screenWidth, screenHeight, "Piano Hero!")
	rl.InitAudioDevice()
	rl.SetConfigFlags(rl.FlagMsaa4xHint) //ENABLE 4X MSAA IF AVAILABLE CONFLICT
	rl.SetTargetFPS(60)
}

func PlayingScreen(songPath string, velocity time.Duration, state *State) {
	score := 0

	camera := rl.Camera{}
	camera.Position = rl.NewVector3(0,7,12)
	camera.Up = rl.NewVector3(0, 1, 0)
	camera.Fovy = 45

	song := ReadSongFromFile(songPath)

	currentChord := 0
	ticker := time.NewTicker((velocity* time.Millisecond))
	go func() {
		for range ticker.C {
			for i := range song{
				for j := range song[i]{
					if i < currentChord{
						song[i][j].Position += 0.1
					}
				}
			}
		}
	}()

	timer := time.NewTicker(10*velocity * time.Millisecond)
	go func() {
		for range timer.C {
			currentChord += 1
		}
	}()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.BeginMode3D(camera)
		rl.ClearBackground(rl.SkyBlue)

		DrawTracks()
		DrawMarker(0, song, &score)
		DrawMarker(1, song, &score)
		DrawMarker(2, song, &score)
		DrawMarker(3, song, &score)
		DrawMarker(4, song, &score)

		for i := range song{
			if i < currentChord{
				for j := range song[i]{
					if song[i][j].Position < 10 && !song[i][j].Pressed{
						DrawDisk(song[i][j].Lane, song[i][j].Position)
					}
				}
			}
		}

		rl.EndMode3D()
		rl.DrawFPS(10,10)
		rl.DrawText(fmt.Sprintf("Score: %d", score), 10, 30, 20, rl.White)
		rl.EndDrawing()
	}

	*state = Menu
}

func BuildMenuOption(option string, index int32, optionName MenuOptions, selectedOption *MenuOptions) {
	text := option
	var fontSize int32 = 20
	color := rl.White
	if *selectedOption == optionName {
		color = rl.Green
	} 
	rl.DrawText(text, CenteredTextPosX(text, fontSize), 300+(fontSize*index*2), fontSize, color)
}

func MenuScreen(state *State) {
	selectedOption := Play

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Gray)

		title := "PIANO HERO!"
		var fontSize int32 = 40
		rl.DrawText(title, CenteredTextPosX(title, fontSize), 150, fontSize, rl.White)
		
		BuildMenuOption("Playlist", 0, Play, &selectedOption)
		BuildMenuOption("Exit", 1, Exit, &selectedOption)

		if rl.IsKeyPressed(rl.KeyDown) {
			if selectedOption < Exit {
				selectedOption++
			} else {
				selectedOption = 0
			}
		}
		if rl.IsKeyPressed(rl.KeyUp) {
			if selectedOption > 0 {
				selectedOption--
			} else {
				selectedOption = Exit
			}
		}

		if rl.IsKeyPressed(rl.KeyEnter) {
			if selectedOption == Exit {
				rl.CloseWindow()
			} else if selectedOption == Play {
				*state = Playlist
				break
			}
		}
		rl.EndDrawing()
	}
}

func SongEntry(song string, index int32, selectedSong int32) {
	text := song
	var fontSize int32 = 20
	color := rl.White
	if selectedSong == index {
		color = rl.Green
	} 
	rl.DrawText(text, CenteredTextPosX(text, fontSize), fontSize+fontSize*index*2, fontSize, color)
}

func PlaylistScreen(state *State, song *string){
	playOnce := 0

	entries, err := os.ReadDir("songs")
	if err != nil {
		fmt.Println(err)
	}

	var selectedSong int32 = 0
	soungsCount := int32(len(entries)-1)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Gray)

		for index, file := range entries {
			SongEntry(file.Name(), int32(index), selectedSong)
		}

		if rl.IsKeyPressed(rl.KeyDown) {
			if selectedSong < soungsCount {
				selectedSong++
			} else {
				selectedSong = 0
			}
		}
		if rl.IsKeyPressed(rl.KeyUp) {
			if selectedSong > 0 {
				selectedSong--
			} else {
				selectedSong = soungsCount
			}
		}

		if rl.IsKeyPressed(rl.KeyEnter) {
			if playOnce == 0 {
				playOnce++
			} else {
				*song = entries[selectedSong].Name()
				*state = Playing
				break
			}
		}
	
		rl.EndDrawing()
	}
}

func PlaySound(note string, scale string){
	fmt.Print("Starting sound\n")
	sound := rl.LoadSound(fmt.Sprintf("notes/%s%s.mp3", note, scale))
	rl.PlaySound(sound)
}