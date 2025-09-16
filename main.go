package main

import(
	rl "github.com/gen2brain/raylib-go/raylib"
	"fmt"
)

func main(){
	rl.InitWindow(800,450,"Game")
	defer rl.CloseWindow()
	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	rl.SetTargetFPS(60);

	gm := NewGameManager()

	for !rl.WindowShouldClose(){
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)



		gm.UpdatePlot()

		if(rl.IsKeyPressed(rl.KeyS)){
			fmt.Println("selected item: ", gm.GetSelectedItem().Name)
		}

		rl.EndDrawing()
	}
}