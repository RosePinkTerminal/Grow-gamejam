package main

// import (
// 	rl "github.com/gen2brain/raylib-go/raylib"
// )

type TestPlant struct{
	Plant
}

func NewTestPlant() TestPlant{
	return TestPlant{
		Plant: NewPlant("Test Plant", 5, 3)}
}