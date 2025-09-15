package main

import (
	"fmt"
	// "image/draw"

	rl "github.com/gen2brain/raylib-go/raylib"
)




func DrawStore(showStore bool, textures map[string]rl.Texture2D, player *Player, buyOrSell bool) {
	// var buyOrSell bool  // true for buy, false for sell

	// Example items in the store
		items := []struct {
			Name  string
			Price int
			Image rl.Texture2D
			Buying bool
		}{
			{"carrot_seed", 5, textures["carrot_seed"], true},
			{"carrot", 20, textures["carrot"], false},
			// Add more items as needed
		}



	if showStore {
		rl.DrawRectangle(100, 100, 600, 400, rl.LightGray)
		rl.DrawText("Store - Press S to close", 120, 120, 20, rl.Black)

		
		
		
			rl.DrawText("Buy", 120, 350, 20, rl.DarkGreen)
			rl.DrawText("Sell", 220, 350, 20, rl.Red)
			rl.DrawText(fmt.Sprintf("Balance: $%d", player.money), 550, 120, 20, rl.Black)
			// setting if you're buying or selling
			

		if buyOrSell {
			for _, item := range items {

				if buyOrSell && item.Buying{
					rl.DrawTexture(item.Image, 120, int32(150+1*80), rl.White)
					rl.DrawText(fmt.Sprintf("$%d", item.Price), 120, int32(150+1*120), 20, rl.Black)
					rl.DrawText("Currently buying!", 120, 150, 20, rl.DarkGreen)

					if(rl.IsMouseButtonPressed(rl.MouseButtonLeft)){
							mouseX := rl.GetMouseX()
							mouseY := rl.GetMouseY()
							if mouseX >= 120 && mouseX <= 170 && mouseY >= 230 && mouseY <= 270 {

								if mouseX >= 120 && mouseX <= 170 && mouseY >= 230 && mouseY <= 270 {
								
									if player.inventory.AddItem("carrot_seed", 5, 1) == nil {

									
									player.money -= 5
									fmt.Println("we bought something hooray")
									fmt.Println(fmt.Sprintln("player Money: ", player.money))
								}
						}


								
						}
					}
				}
		} 



		} else if !buyOrSell {
			for i, item := range items {

				if  !item.Buying{
					rl.DrawText("Currently selling!", 120, 150, 20, rl.Red)
					rl.DrawTexture(item.Image, 120, int32(150+i*80), rl.White)
					rl.DrawText(fmt.Sprintf("$%d", item.Price), 120, int32(150+1*120), 20, rl.Black)

						if(rl.IsMouseButtonPressed(rl.MouseButtonLeft)){
							mouseX := rl.GetMouseX()
							mouseY := rl.GetMouseY()
							if mouseX >= 120 && mouseX <= 170 && mouseY >= 230 && mouseY <= 270 {
								
								if player.inventory.RemoveItem("carrot", 1) == nil {

									
									player.money += 20
									fmt.Println("we sold something hooray")
									fmt.Println(fmt.Sprintln("player Money: ", player.money))
								}
						}
					}

		}
}

	
}
	}
}