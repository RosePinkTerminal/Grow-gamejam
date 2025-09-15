package main

import (
	"fmt"
	"time"
// 	rl "github.com/gen2brain/raylib-go/raylib"
)

type Plant struct{
	name string
	currentStage int
	maxStage int
	timerTime time.Duration
	watered bool
	fertilized bool
	readyToHarvest bool
	//animation object
	//physics object
	// stageTimer time.Timer
}

func NewPlant(_name string, _timerTime time.Duration, max int) Plant{
	plantObj := Plant{
		name: _name,
		currentStage: 0,
		maxStage: max,
		timerTime: _timerTime,
		watered: false,
		fertilized: false,
		readyToHarvest: false}

	//new animation object
	//new physics object

	plantObj.SetPlantTimer(plantObj.timerTime)

	return plantObj
}

func (p Plant) PrintPlant(){
	fmt.Println("---\nName: ", p.name, "\nStage: ", p.currentStage, "\nMax: ", p.maxStage, "\nWatered: ", p.watered, "\nFertilized: ", p.fertilized, "\nHarvestable: ", p.readyToHarvest, "\nTimer: "/*, p.stageTimer.GetDuration(), "\n---"*/)
}

func (p *Plant) UpdatePlant(){
	//update sprite
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
	p.SetPlantTimer(p.timerTime)
}

func (p *Plant) UpdateStage(){
	fmt.Println("Stage updated")
	if(p.GetStage() >= p.GetMaxStage()){ 
		p.SetReadyToHarvest()
	}else{p.IncrementStage()}
}

func (p *Plant) IsWatered() bool{
	return p.watered
}

func (p *Plant) SetWatered(){
	p.watered = true
}

func (p *Plant) IsFertilized() bool{
	return p.fertilized
}

func (p *Plant) SetFertilized(){
	p.fertilized = true
} 

func (p *Plant) IsReadyToHarvest() bool{
	return p.readyToHarvest
}

func (p *Plant) SetReadyToHarvest(){
	p.readyToHarvest = true
	fmt.Println("Ready to harvest!")
}

func (p *Plant) SetPlantTimer(endTime time.Duration){
	stageTimer := *time.NewTimer(time.Second * endTime)
	go func(p *Plant){
		<-stageTimer.C
		fmt.Println("timer done")
		p.UpdateStage()
	}(p)
}