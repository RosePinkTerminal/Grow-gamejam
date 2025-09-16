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

func (gm GameManager) GetSelectedItem() Item{
	return gm.Slots[gm.selectedItem];
}