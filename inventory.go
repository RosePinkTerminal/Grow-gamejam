package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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

func (inv *Inventory) AddItem(name string, quantity int) error {
	if len(inv.Slots) >= inv.Limit {
		return fmt.Errorf("inventory full")
	}
	for i, item := range inv.Slots {
		if item.Name == name {
			inv.Slots[i].Quantity += quantity
			return nil
		}
	}
	inv.Slots = append(inv.Slots, Item{Name: name, Quantity: quantity})
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

func InventoryExample() {
	rl.InitWindow(800, 600, "Inventory Example")
	defer rl.CloseWindow()

	inventory := NewInventory(10)
	inventory.AddItem("Potion", 5)
	inventory.AddItem("Sword", 1)
	inventory.AddItem("Sword", 1)
	inventory.AddItem("Sword", 1)
	inventory.AddItem("Sword", 1)
	inventory.AddItem("Sword", 1)
	inventory.ListItems()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawRectangle(int32(rl.GetScreenWidth())/20, int32(rl.GetScreenHeight())/10, int32(rl.GetScreenWidth())/2 + 50 * 5, int32(rl.GetScreenHeight())/2 + 25 * 2 , rl.LightGray)
		rl.DrawText("Inventory", int32(rl.GetScreenWidth())/10, int32(rl.GetScreenHeight())/10, 20, rl.Black)

		for i, item := range inventory.Slots {
			rl.DrawText(fmt.Sprintf("%s: %d", item.Name, item.Quantity), int32(10), int32(10+i*20), 20, rl.Black)
		}

		rl.EndDrawing()
	}
}


