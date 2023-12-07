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

func DrawMeter(barrinha float32){
	/* Version made with cubes (edges are smooths)*/
	rl.PushMatrix()
		rl.Rotatef(60, 1, 0, 0)
		rl.DrawCube(rl.NewVector3(-6, 0.99, 1), 4, 0.01, 3, rl.ColorFromNormalized(rl.NewVector4(0.96, 0.95, 0.64, 1)))
		rl.DrawCube(rl.NewVector3(-7, 1, 1,), 1, 0.2, 0.5, rl.ColorFromNormalized(rl.NewVector4(1, 0, 0, 1)))
		rl.DrawCube(rl.NewVector3(-6, 1, 1,), 1, 0.2, 0.5, rl.ColorFromNormalized(rl.NewVector4(1, 1, 0, 1)))
		rl.DrawCube(rl.NewVector3(-5, 1, 1,), 1, 0.2, 0.5, rl.ColorFromNormalized(rl.NewVector4(0, 1, 0, 1)))
		rl.DrawCube(rl.NewVector3(barrinha, 1.001, 1,), 0.2,0.2, 1, rl.Black)
 	rl.PopMatrix()
	
	/* Version made with planes*/
	
	// rl.PushMatrix()
	// 	rl.Rotatef(65, 1, 0, 0)
	// 	rl.DrawPlane(rl.NewVector3(-6, 2, 1), rl.NewVector2(4, 2), rl.ColorFromNormalized(rl.NewVector4(0.96, 0.95, 0.64, 1)))
	// 	rl.DrawPlane(rl.NewVector3(-7, 2.1, 1), rl.NewVector2(1, 0.5), rl.ColorFromNormalized(rl.NewVector4(1, 0, 0, 1)))
	// 	rl.DrawPlane(rl.NewVector3(-6, 2.1, 1), rl.NewVector2(1, 0.5), rl.ColorFromNormalized(rl.NewVector4(1, 1, 0, 1)))
	// 	rl.DrawPlane(rl.NewVector3(-5, 2.1, 1), rl.NewVector2(1, 0.5), rl.ColorFromNormalized(rl.NewVector4(0, 1, 0, 1)))
	// 	rl.DrawPlane(rl.NewVector3(barrinha, 2.2, 1), rl.NewVector2(0.2, 1.5), rl.ColorFromNormalized(rl.NewVector4(0, 0, 0, 1)))
	// rl.PopMatrix()

   
}

func DrawDisk(lane int, position float32) {
	// -10 - Disk spawn position
	// 6 - Disk perfect position
	// 10 - Disk destroy position

	colors := []color.RGBA{rl.Green, rl.Red, rl.Yellow, rl.Blue, rl.Orange}
	rl.DrawSphere(rl.NewVector3(float32(-2+lane), 0.1, position), 0.4, colors[lane])
}

func DrawMarker(lane int, song [][]Note, score *int, barrinha *float32, paused *bool) {
	colors := []color.RGBA{rl.Green, rl.Red, rl.Yellow, rl.Blue, rl.Orange}
	var color color.RGBA

	keys := []int32{rl.KeyA, rl.KeyS, rl.KeyJ, rl.KeyK, rl.KeyL}
	height := 0.5

	if rl.IsKeyDown(keys[lane]) && !*paused {
		height = 0.6
		color = colors[lane]
		foundNote := false
		

		if rl.IsKeyPressed(keys[lane]){
			for i := range song {
				for j := range song[i] {
					if song[i][j].Position > 5.5 && song[i][j].Position < 6.5 && song[i][j].Lane == lane && !song[i][j].Pressed {
						song[i][j].Pressed = true
						*score += 1
						if *barrinha < -4.6 {
							*barrinha += 0.1
						}
						foundNote = true
						PlaySound(notes[lane], "4")
					}
				}
			}
			if !foundNote {
				*score -= 1
				if *barrinha > -7.3 {
					*barrinha -= 0.1
				}
				PlayWrongNote()
			}
		}

	} else {
		height = 0.5
		color = rl.DarkGray
	}

	rl.DrawCylinder(rl.NewVector3(float32(-2+lane), -0.25, 6.5), 0.4, 0.4, 0.5, 12, rl.Gray)
	rl.DrawCylinder(rl.NewVector3(float32(-2+lane), -0.2, 6.5), 0.3, 0.3, float32(height), 12, color)
}