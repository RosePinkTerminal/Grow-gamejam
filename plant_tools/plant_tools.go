package main

import (
	"fmt"
)

func main() {

	fmt.Println("Wait I'm so goated")
	myTool := createTool("Hoe", 3000, 5000)
	printTool(&myTool)
	setToolType(&myTool, "Axe")
	getToolType(&myTool)
	printTool(&myTool)

}

type Tool struct {
	typeTool   string
	currentLvl int
	maxLvl     int
}

func createTool(_typeTool string, _currentLvl int, _maxLvl int) Tool {

	return Tool{
		typeTool: _typeTool, currentLvl: _currentLvl, maxLvl: _maxLvl,
	}

}

func printTool(t *Tool) {

	fmt.Println("Type of tool: ", t.typeTool, " Current Level: ", t.currentLvl, " Max Level: ", t.maxLvl)

}

// now we need to create basic functions for atts, like getters, and setters
func getToolType(t *Tool) string {
	return t.typeTool
}

func setToolType(t *Tool, newType string) {
	t.typeTool = newType
}

func getCurrentLvl(t *Tool) int {
	return t.currentLvl
}

func setCurrentLvl(t *Tool, newCurrentLvl int) {
	t.currentLvl = newCurrentLvl
}

func maxLvl(t *Tool) int {
	return t.maxLvl
}

func setMaxLvl(t *Tool, newType string) {
	t.typeTool = newType
}

// function to increase level of tool
func increaseToolLvl(t *Tool, incrNum int) {
	t.currentLvl += incrNum
	if t.currentLvl >= t.maxLvl {
		t.currentLvl = t.maxLvl
	}
}
