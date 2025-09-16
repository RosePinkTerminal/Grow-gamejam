package main

import (
	"fmt"
	"time"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Plant struct{
	name string
	currentStage int
	maxStage int
	timerTime time.Duration
	harvestable bool
	stageText []rl.Texture2D
	//physics object
	// stageTimer time.Timer
}

func NewPlant(_name string, _timerTime time.Duration, max int, textures []rl.Texture2D) *Plant{
	plantObj := Plant{
		name: _name,
		currentStage: 1,
		maxStage: max,
		timerTime: _timerTime,
		harvestable: false,
		stageText: textures}

	plantObj.SetPlantTimer(plantObj.timerTime)

	return &plantObj
}

func (p Plant) PrintPlant(){
	fmt.Println("---\nName: ", p.name, "\nStage: ", p.currentStage, "\nMax: ", p.maxStage, "\nHarvestable: ", p.harvestable, "\nTimer: "/*, p.stageTimer.GetDuration(), "\n---"*/)
}

func (p Plant) DrawPlant(location rl.Vector2){
	//check if texture at stage-1 exists
	if(len(p.stageText) >= p.currentStage-1 && &p.stageText[p.currentStage-1] != nil){
		//draw texture at stage-1
		DrawTextureEz(p.stageText[p.currentStage-1], rl.NewVector2(location.X+32,location.Y+32), 0, 1, rl.White)
	}
}

func (p *Plant) GetName() string{
	return p.name
}

func (p Plant) GetStage() int{
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
	fmt.Println("Stage updated, ", p.currentStage)
	if(p.currentStage < p.maxStage){
		p.IncrementStage()
	}
	
	if(p.currentStage >= p.maxStage){ 
		p.SetReadyToHarvest()
	}
} 

func (p *Plant) IsReadyToHarvest() bool{
	return p.harvestable
}

func (p *Plant) SetReadyToHarvest(){
	p.harvestable = true
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