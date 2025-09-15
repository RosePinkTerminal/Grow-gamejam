package main

import()

type Plot struct{
	*Plant
	tilled bool
	watered bool
	fertilized bool
	hasPlant bool
	//location
}

func NewPlot() *Plot{
	return &Plot{
		Plant: nil,
		tilled:false,
		watered:false,
		fertilized:false,
		hasPlant:false}
}

func (p Plot)DrawPlot(){
	if(p.HasPlant()){
			p.GetPlant().DrawPlant();
	}
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

func (p Plot)GetPlant() Plant{
	return *p.Plant;
}