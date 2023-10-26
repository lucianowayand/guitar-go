package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func CameraMovement(camera *rl.Camera3D, velocity float32){
	if(rl.IsKeyDown(rl.KeyUp)){
		rl.UpdateCameraPro(camera, rl.NewVector3(velocity,0,0), rl.Vector3Zero(), 0)
	} 
	if(rl.IsKeyDown(rl.KeyDown)){
		rl.UpdateCameraPro(camera, rl.NewVector3(-velocity,0,0), rl.Vector3Zero(), 0)
	} 
	if(rl.IsKeyDown(rl.KeyLeft)){
		rl.UpdateCameraPro(camera, rl.NewVector3(0,-velocity,0), rl.Vector3Zero(), 0)
	} 
	if(rl.IsKeyDown(rl.KeyRight)){
		rl.UpdateCameraPro(camera, rl.NewVector3(0,velocity,0), rl.Vector3Zero(), 0)
	} 

	if(rl.IsKeyDown(rl.KeyX)){
		rl.UpdateCameraPro(camera, rl.NewVector3(0,0,velocity),rl.Vector3Zero(), 0)
	}
	if(rl.IsKeyDown(rl.KeyZ)){
		rl.UpdateCameraPro(camera, rl.NewVector3(0,0,-velocity),rl.Vector3Zero(), 0)
	}

	if(rl.IsKeyDown(rl.KeyA)){
		rl.UpdateCameraPro(camera, rl.Vector3Zero(), rl.NewVector3(-velocity*10,0,0), 0)
	}
	if(rl.IsKeyDown(rl.KeyD)){
		rl.UpdateCameraPro(camera, rl.Vector3Zero(), rl.NewVector3(velocity*10,0,0), 0)
	}
	if(rl.IsKeyDown(rl.KeyW)){
		rl.UpdateCameraPro(camera, rl.Vector3Zero(), rl.NewVector3(0,-velocity*10,0), 0)
	}
	if(rl.IsKeyDown(rl.KeyS)){
		rl.UpdateCameraPro(camera, rl.Vector3Zero(), rl.NewVector3(0,velocity*10,0), 0)
	}

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

func DrawDisk(lane int){
	colors := []color.RGBA{rl.Green, rl.Red, rl.Yellow, rl.Blue, rl.Orange}
	rl.DrawSphere(rl.NewVector3(float32(-2+lane),0.1,0), 0.4, colors[lane])
}

func main() {
	screenWidth := int32(1280)
	screenHeight := int32(720)

	rl.InitWindow(screenWidth, screenHeight, "raylib [shaders] example - basic lighting")
	rl.SetConfigFlags(rl.FlagMsaa4xHint) //ENABLE 4X MSAA IF AVAILABLE
	rl.SetTargetFPS(60)

	camera := rl.Camera{}
	camera.Position = rl.NewVector3(0,7,12)
	camera.Up = rl.NewVector3(0, 1, 0)
	camera.Fovy = 45

	for !rl.WindowShouldClose() {
		CameraMovement(&camera, 0.2)

		rl.BeginDrawing()
		rl.BeginMode3D(camera)
		rl.ClearBackground(rl.Gray)

		DrawTracks()
		DrawDisk(0)
		DrawDisk(1)
		DrawDisk(2)
		DrawDisk(3)
		DrawDisk(4)

		rl.EndMode3D()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}