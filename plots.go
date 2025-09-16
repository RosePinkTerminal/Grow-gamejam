package main

import(
	rl "github.com/gen2brain/raylib-go/raylib"
	"fmt"
)

type Plot struct{
	*Plant
	tilled bool
	watered bool
	fertilized bool
	hasPlant bool
	isMouseOver bool
	location rl.Vector2
}

func NewPlot(_location rl.Vector2) *Plot{
	return &Plot{
		Plant: nil,
		tilled:false,
		watered:false,
		fertilized:false,
		hasPlant:false,
		isMouseOver: false,
		location: _location}
}

func (gm *GameManager)Select(p *Plot){
	switch(gm.Slots[gm.selectedItem].Name){
	case CARROT_SEED:
			p.AddPlant(NewCarrotPlant());
			gm.Slots[gm.selectedItem].Quantity--;
	}
}

func (p Plot)DrawPlot(){
	if(!p.HasPlant()){
		p.DrawEmptyPlot()
	}else{
		p.GetPlant().DrawPlant(p.location);
	}
}

func (p Plot)DrawEmptyPlot(){
	//draw rectangle at plot
	rl.DrawRectangle(int32(p.location.X+16), int32(p.location.Y+16), 32, 32, rl.Brown)
}

//fully resets plot to empty state
func (p *Plot)ResetPlot(){
	p.RemovePlant();
	p.tilled = false;
	p.watered = false;
	p.fertilized = false;
	p.hasPlant = false;
}

//only resets tilled, watered, and fertilized state
func (p *Plot)ResetState(){
	p.tilled = false;
	p.watered = false;
	p.fertilized = false;
}

func (p *Plot)AddPlant(plant *Plant){
	p.Plant = plant;
	p.hasPlant = true;
}

func (p *Plot)RemovePlant(){
	p.Plant = nil;
	p.hasPlant = false;
}

func (p Plot)IsTilled() bool{
	return p.tilled;
}

func (p *Plot)Till(){
	p.tilled = true;
}

func (p Plot)IsWatered() bool{
	return p.watered;
}

func (p *Plot)Water(){
	p.watered = true;
}

func (p Plot)isFertilized() bool{
	return p.fertilized;
}

func (p *Plot)Fertilize(){
	p.fertilized = true;
}

func (p Plot)HasPlant() bool{
	return p.hasPlant;
}

func (p Plot)GetPlant() *Plant{
	if(!p.hasPlant){
		return nil
	}
	return p.Plant;
}

func (gm *GameManager)Harvest (p *Plot){
	if(p.GetStage() >= p.GetMaxStage()){
		fmt.Println("yahoo!");

		switch(p.GetName()){
		case CARROT_PLANT:
			gm.AddItem(CARROT, CARROT_VALUE, 1);
			p.RemovePlant();
			break;
		}
	}
}