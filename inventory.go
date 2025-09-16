package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//item name constants
const CARROT_SEED string = "carrot_seed"
const CARROT string = "carrot"
const CARROT_VALUE int = 5
const HOE string = "hoe"

type Item struct {
	Name     string
	Value    int
	Quantity int
}

type Inventory struct {

    Slots         []*Item
    Limit         int
    SelectedIndex int // -1 means none selected
}

func NewInventory(limit int) *Inventory {
    return &Inventory{
        Slots:         make([]*Item, 0, limit),
        Limit:         limit,
        SelectedIndex: -1,
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
	inv.Slots = append(inv.Slots, &Item{Name: name, Value: value, Quantity: quantity})
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
    if !showInventory {
        return
    }

    // Inventory window background
    rl.DrawRectangle(
        int32(rl.GetScreenWidth())/20,
        int32(rl.GetScreenHeight())/10,
        int32(rl.GetScreenWidth()/2)-80,
        int32(rl.GetScreenHeight())/5+25*2,
        rl.LightGray,
    )
    rl.DrawText("Inventory", int32(rl.GetScreenWidth())/15, int32(rl.GetScreenHeight())/8, 20, rl.Black)

    mousePos := rl.GetMousePosition()

    // Loop through slots
    for index := 0; index < inventory.Limit; index++ {
        // Compute slot position
        x := int32(rl.GetScreenWidth())/20 + 14 + int32(index%5)*60
        y := int32(rl.GetScreenHeight())/10 + 40 + int32(index/5)*60
        slotRect := rl.NewRectangle(float32(x), float32(y), 50, 50)

        hovered := rl.CheckCollisionPointRec(mousePos, slotRect)

        // Change color if hovered or selected
        slotColor := rl.DarkGray
        if hovered {
            slotColor = rl.LightGray
        }
        rl.DrawRectangleRec(slotRect, slotColor)

        // Draw selection outline if this slot is selected
        if index == inventory.SelectedIndex {
            rl.DrawRectangleLinesEx(slotRect, 3, rl.Green)
        } else {
            rl.DrawRectangleLinesEx(slotRect, 1, rl.Black)
        }

        // Handle click
        if hovered && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
            inventory.SelectedIndex = index
            fmt.Println("Selected slot:", inventory.SelectedIndex)
        }

        // Draw item if present
        if index < len(inventory.Slots) {
            item := inventory.Slots[index]
            rl.DrawText(fmt.Sprintf("x%d", item.Quantity), x+6, y+38, 10, rl.White)

            // Draw texture if available
            if tex, ok := textures[item.Name]; ok {
                rl.DrawTexture(tex, x+6, y+5, rl.White)
            } else {
                rl.DrawText(item.Name, x+5, y+20, 10, rl.White)
            }
        }
    }
}
