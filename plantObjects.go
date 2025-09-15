package main

// import (
// 	rl "github.com/gen2brain/raylib-go/raylib"
// )

//---carrot plant---
type CarrotPlant struct{
	Plant
}

func NewCarrotPlant() CarrotPlant{
	return CarrotPlant{Plant: *NewPlant("Carrot Plant", 5, 4)}
}

//---test plant----
/////for testing only, dont include in final game
type TestPlant struct{
	Plant
}

func NewTestPlant() TestPlant{
	return TestPlant{
		Plant: *NewPlant("Test Plant", 5, 3)}
}