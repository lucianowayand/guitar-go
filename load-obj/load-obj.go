package main

import (
	"fmt"

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
	screenWidth := int32(1280)
	screenHeight := int32(720)

	rl.InitWindow(screenWidth, screenHeight, "Garrafa")
	rl.SetTargetFPS(60)

	camera := rl.Camera{}
	camera.Position = rl.NewVector3(15,7,17)
	camera.Up = rl.NewVector3(0, 1, 0)
	camera.Fovy = 45

	// model := rl.LoadModel("../models/garrafa1.obj")
	// logo := rl.LoadTexture("../models/madeira.png")
	// model.Materials.Maps.Texture = logo

	terra := rl.LoadModel("../models/bola.obj")
	terra_text := rl.LoadTexture("../models/MAPA-DO-MUNDO.jpg")
	terra.Materials.Maps.Texture = terra_text

	lua := rl.LoadModel("../models/bola-peq.obj")
	lua_text := rl.LoadTexture("../models/lua.jpg")
	lua.Materials.Maps.Texture = lua_text
	
	for !rl.WindowShouldClose() {
		CameraMovement(&camera, 0.2)

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode3D(camera)

		rl.DrawModel(terra, rl.NewVector3(0, 0, 0), 1, rl.White)
		rl.DrawModel(lua, rl.NewVector3(0, 0, 0), 1, rl.White)

		// rl.DrawGrid(10, 1.0)
		rl.EndMode3D()
		rl.DrawText(fmt.Sprintf("[%.2f, %.2f, %.2f]", camera.Position.X, camera.Position.Y, camera.Position.Z), 10, 20, 20, rl.Black)
		rl.EndDrawing()
	}

	rl.UnloadModel(terra)
	rl.UnloadModel(lua)
	rl.CloseWindow()
}