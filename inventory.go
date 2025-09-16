package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//item name constants
const CARROT_SEED string = "carrot_seed"
const CARROT string = "carrot"
const HOE string = "hoe"

type Item struct {
	Name     string
	Value    int
	Quantity int
}

type Inventory struct {
	Slots []Item
	Limit int
}

func NewInventory(limit int) *Inventory {
	return &Inventory{
		Slots: make([]Item, 0, limit),
		Limit: limit,
	}
}

func (inv *Inventory) AddItem(name string, value int, quantity int) error {
	if len(inv.Slots) >= inv.Limit {
		return fmt.Errorf("inventory full")
	}
	for i, item := range inv.Slots {
		if item.Name == name {
			inv.Slots[i].Quantity += quantity
			return nil
		}
	}
	inv.Slots = append(inv.Slots, Item{Name: name, Value: value, Quantity: quantity})
	return nil
}

func (inv *Inventory) RemoveItem(name string, quantity int) error {
	for i, item := range inv.Slots {
		if item.Name == name {
			if item.Quantity < quantity {
				return fmt.Errorf("not enough quantity")
			}
			inv.Slots[i].Quantity -= quantity
			if inv.Slots[i].Quantity == 0 {
				inv.Slots = append(inv.Slots[:i], inv.Slots[i+1:]...)
			}
			return nil
		}
	}
	return fmt.Errorf("item not found")
}

func (inv *Inventory) ListItems() {
	if len(inv.Slots) == 0 {
		fmt.Println("Inventory is empty")
		return
	}
	for _, item := range inv.Slots {
		fmt.Printf("%s: %d\n", item.Name, item.Quantity)
	}
}

func DrawInventory(showInventory bool, textures map[string]rl.Texture2D, inventory *Inventory) {


	inventory.ListItems()

	// Draw inventory UI once per call
	rl.DrawRectangle(int32(rl.GetScreenWidth())/20, int32(rl.GetScreenHeight())/10, int32(rl.GetScreenWidth()/2)-80, int32(rl.GetScreenHeight())/5+25*2, rl.LightGray)
	rl.DrawText("Inventory", int32(rl.GetScreenWidth())/15, int32(rl.GetScreenHeight())/8, 20, rl.Black)

	for index := 0; index < inventory.Limit; index++ {
		rl.DrawRectangle(int32(rl.GetScreenWidth())/20+14+int32(index%5)*60, int32(rl.GetScreenHeight())/10+40+int32(index/5)*60, 50, 50, rl.DarkGray)
		if index < len(inventory.Slots) {
    		item := inventory.Slots[index]
    		rl.DrawText(fmt.Sprintf("x%d", item.Quantity), int32(rl.GetScreenWidth())/20+20+int32(index%5)*60, int32(rl.GetScreenHeight())/10+78+int32(index/5)*60, 10, rl.White)

	// use item name to load texture dynamically
    		if tex, ok := textures[item.Name]; ok {
        		rl.DrawTexture(tex, int32(rl.GetScreenWidth())/20+20+int32(index%5)*60, int32(rl.GetScreenHeight())/10+45+int32(index/5)*60, rl.White)
    		} 
		}
	}
}
