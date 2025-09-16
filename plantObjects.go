package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

//---carrot plant---

const CARROT_STAGES = 4
const CARROT_PLANT string = "carrot_plant"

func NewCarrotPlant() *Plant{
	return NewPlant(CARROT_PLANT, 5, CARROT_STAGES, LoadCarrotTextures())
}

func LoadCarrotTextures() []rl.Texture2D{
	textures := make([]rl.Texture2D, CARROT_STAGES)
	//load stage img 1-4 into textures
	textures[0] = rl.LoadTexture("assets/carrot1.png")
	textures[1] = rl.LoadTexture("assets/carrot2.png")
	textures[2] = rl.LoadTexture("assets/carrot3.png")
	textures[3] = rl.LoadTexture("assets/carrot4.png")

	return textures
}

//---test plant----
/////for testing only, dont include in final game
type TestPlant struct{
	Plant
}

func NewTestPlant() TestPlant{
	return TestPlant{
		Plant: *NewPlant("Test Plant", 5, 3, make([]rl.Texture2D, 0))}
}