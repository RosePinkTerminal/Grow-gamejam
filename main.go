package main

import(
	rl "github.com/gen2brain/raylib-go/raylib"
	"fmt"
)

type Player struct {
	inventory *Inventory
	money     int
}

func main(){
	rl.InitWindow(800,450,"Game")
	defer rl.CloseWindow()
	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	rl.SetTargetFPS(60);

	var player Player = Player{
		inventory: NewInventory(10),
		money:     25,
	}

	player.inventory.AddItem("carrot", 20, 10)

	var showInventory bool = false
	var showStore bool = false
	var buying bool = false

	// Load all textures into a map
	textures := make(map[string]rl.Texture2D)
	textures["carrot"] = rl.LoadTexture("assets/carrot..png")
	textures["carrot_seed"] = rl.LoadTexture("assets/carrot_seed.png")
	textures["shopkeeper"] = rl.LoadTexture("assets/shopkeeper.png")

	// Add more textures as needed
	defer rl.UnloadTexture(textures["carrot"])

	gm := NewGameManager()

	for !rl.WindowShouldClose(){
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)


		gm.UpdatePlot()

		if rl.IsKeyPressed(rl.KeyE) {
			showInventory = !showInventory
		}

		if showInventory {
			DrawInventory(showInventory, textures, player.inventory)
		}

		if rl.IsKeyPressed(rl.KeyS) {
			showStore = !showStore
		}

		if showStore {
			DrawStore(showStore, textures, &player, buying)
		}

		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && showStore {
			mouseX := rl.GetMouseX()
			mouseY := rl.GetMouseY()
			if mouseX >= 120 && mouseX <= 200 && mouseY >= 330 && mouseY <= 380 {
				buying = true
				fmt.Println("buying: ", buying)
			} else if mouseX >= 240 && mouseX <= 320 && mouseY >= 330 && mouseY <= 380 {
				buying = false
				fmt.Println("buying: ", buying)
			}
		}

		if(rl.IsKeyPressed(rl.KeyI)){
			fmt.Println("selected index: ", gm.SelectedIndex);
		}

		rl.EndDrawing()
	}
}