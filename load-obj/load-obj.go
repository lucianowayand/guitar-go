package main

import (
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

func main() {
	screenSize := rl.NewVector2(1280,720)

	rl.InitWindow(int32(screenSize.X), int32(screenSize.Y), "Guitar Go!")

	camera := rl.Camera{}
	camera.Position = rl.NewVector3(0,10,10)
	camera.Up = rl.NewVector3(0, 1, 0)
	camera.Fovy = 45

	rl.SetTargetFPS(60)
	model := rl.LoadModel("./models/garrafa.obj")
	

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Gray)

		CameraMovement(&camera, 0.2)
		rl.BeginMode3D(camera)

		rl.DrawGrid(16,5)
		rl.DrawModel(model, rl.NewVector3(0,0,0), 1, rl.White)

		rl.EndMode3D()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}