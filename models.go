package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawTracks() {
	// Track
	rl.DrawCube(rl.NewVector3(0, 0, 0), 5, 0.5, 20, rl.ColorFromNormalized(rl.NewVector4(0.35, 0.35, 0.35, 1)))

	// Lanes
	rl.DrawCube(rl.NewVector3(-2.5, 0.05, 0), 0.05, 0.5, 20, rl.White)
	rl.DrawCube(rl.NewVector3(-1.5, 0.05, 0), 0.05, 0.5, 20, rl.White)
	rl.DrawCube(rl.NewVector3(-0.5, 0.05, 0), 0.05, 0.5, 20, rl.White)
	rl.DrawCube(rl.NewVector3(0.5, 0.05, 0), 0.05, 0.5, 20, rl.White)
	rl.DrawCube(rl.NewVector3(1.5, 0.05, 0), 0.05, 0.5, 20, rl.White)
	rl.DrawCube(rl.NewVector3(2.5, 0.05, 0), 0.05, 0.5, 20, rl.White)
}

func DrawDisk(lane int, position float32) {
	// -10 - Disk spawn position
	// 6 - Disk perfect position
	// 10 - Disk destroy position

	colors := []color.RGBA{rl.Green, rl.Red, rl.Yellow, rl.Blue, rl.Orange}
	rl.DrawSphere(rl.NewVector3(float32(-2+lane), 0.1, position), 0.4, colors[lane])
}

func DrawMarker(lane int, song [][]Note, score *int) {
	colors := []color.RGBA{rl.Green, rl.Red, rl.Yellow, rl.Blue, rl.Orange}
	var color color.RGBA

	keys := []int32{rl.KeyQ, rl.KeyW, rl.KeyE, rl.KeyR, rl.KeyT}
	height := 0.5

	if rl.IsKeyDown(keys[lane]) {
		height = 0.6
		color = colors[lane]
		foundNote := false

		if rl.IsKeyPressed(keys[lane]) {
			for i := range song {
				for j := range song[i] {
					if song[i][j].Position > 5.5 && song[i][j].Position < 6.5 && song[i][j].Lane == lane && !song[i][j].Pressed {
						song[i][j].Pressed = true
						*score += 1
						foundNote = true
					}
				}
			}
			if !foundNote {
				*score -= 1
			}
		}

	} else {
		height = 0.5
		color = rl.DarkGray
	}

	rl.DrawCylinder(rl.NewVector3(float32(-2+lane), -0.25, 6.5), 0.4, 0.4, 0.5, 12, rl.Gray)
	rl.DrawCylinder(rl.NewVector3(float32(-2+lane), -0.2, 6.5), 0.3, 0.3, float32(height), 12, color)
}