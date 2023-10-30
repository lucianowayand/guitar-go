package game

import (
	"bufio"
	"fmt"
	"image/color"
	"os"
	"strconv"
	"strings"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func ReadSongFromFile(path string) [][]Note{
	f, err := os.Open(fmt.Sprintf("songs/"+path))
    if err != nil {
        panic(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)

	notes := [][]Note{} 

    for scanner.Scan() {
       strNotes := strings.Split(scanner.Text(), "")
	   
	   elements := []Note{}
	   for _, strNote := range strNotes {
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

func DrawTracks(){
	// Track
	rl.DrawCube(rl.NewVector3(0,0,0), 5, 0.5, 20, rl.DarkGray)

	// Lanes
	rl.DrawCube(rl.NewVector3(-2.5,0.05,0), 0.05, 0.5, 20, rl.White)
	rl.DrawCube(rl.NewVector3(-1.5,0.05,0), 0.05, 0.5, 20, rl.White)
	rl.DrawCube(rl.NewVector3(-0.5,0.05,0), 0.05, 0.5, 20, rl.White)
	rl.DrawCube(rl.NewVector3(0.5,0.05,0), 0.05, 0.5, 20, rl.White)
	rl.DrawCube(rl.NewVector3(1.5,0.05,0), 0.05, 0.5, 20, rl.White)
	rl.DrawCube(rl.NewVector3(2.5,0.05,0), 0.05, 0.5, 20, rl.White)
}

func DrawDisk(lane int, position float32){
	// -10 - Disk spawn position
	// 6 - Disk perfect position
	// 10 - Disk destroy position

	colors := []color.RGBA{rl.Green, rl.Red, rl.Yellow, rl.Blue, rl.Orange}
	rl.DrawSphere(rl.NewVector3(float32(-2+lane),0.1,position), 0.4, colors[lane])
}

type Note struct {
	Position float32
	Lane int
}

func Play(songPath string, velocity time.Duration) {
	screenWidth := int32(1280)
	screenHeight := int32(720)

	rl.InitWindow(screenWidth, screenHeight, "Guitar Go!")
	rl.SetConfigFlags(rl.FlagMsaa4xHint) //ENABLE 4X MSAA IF AVAILABLE
	rl.SetTargetFPS(60)

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

		for i := range song{
			if i < currentChord{
				for j := range song[i]{
					if(song[i][j].Position < 10){
						DrawDisk(song[i][j].Lane, song[i][j].Position)
					}
				}
			}
		}
		rl.EndMode3D()

		rl.DrawFPS(10,10)
		if len(song) < currentChord {
			rl.DrawText("GAME OVER", (screenWidth/2) - 120 , 20, 40, rl.White)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}