package main

// import (
// 	rl "github.com/gen2brain/raylib-go/raylib"
// )

type Plant struct{
	name string
	currentStage int
	maxStage int
	watered bool
	fertilized bool
	readyToHarvest bool
	//animation object
	//physics object
	//timer object
}

func NewPlant(_name string, max int) Plant{
	//new animation object
	//new physics object
	//new timer

	return Plant{
		name: _name,
		currentStage: 0,
		maxStage: max,
		watered: false,
		fertilized: false,
		readyToHarvest: false}
}

func (p *Plant) GetName() string{
	return p.name
}

func (p *Plant) GetStage() int{
	return p.currentStage
}

func (p *Plant) GetMaxStage() int{
	return p.maxStage
}

func (p *Plant) IncrementStage(){
	p.currentStage++
}

func (p *Plant) IsWatered() bool{
	return p.watered
}

func (p *Plant) IsFertilized() bool{
	return p.fertilized
}

func (p *Plant) IsReadyToHarvest() bool{
	return p.readyToHarvest
}