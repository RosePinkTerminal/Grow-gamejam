package main

import(
	rl "github.com/gen2brain/raylib-go/raylib"
	// "fmt"
)

type GameManager struct{
	*PlotTiles
	selectedItem int
	*Inventory
}

func NewGameManager() GameManager{
	pt := GenerateTilesFromGrid(5, 5, rl.NewVector2(32, 32));
	inv := NewInventory(5);
	inv.AddItem(CARROT_SEED, 5, 1);

	return GameManager{PlotTiles: pt, selectedItem: 0, Inventory: inv};
}

func (gm GameManager) GetSelectedItem() *Item{
	return gm.Slots[gm.selectedItem];
}

//better draw texture function
func DrawTextureEz(texture rl.Texture2D, pos rl.Vector2, angle float32, scale float32, color rl.Color) {
    sourceRect := rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height))
    destRect := rl.NewRectangle(pos.X, pos.Y, float32(texture.Width)*scale, float32(texture.Height)*scale)
    origin := rl.Vector2Scale(rl.NewVector2(float32(texture.Width)/2, float32(texture.Height)/2), scale)
    rl.DrawTexturePro(texture, sourceRect,
        destRect,
        origin, angle, color)
}