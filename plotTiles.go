package main

import(
	rl "github.com/gen2brain/raylib-go/raylib"
	"fmt"
)

type PlotTiles struct{
	plots []*Plot
	pivot rl.Vector2
}

const HALF_TILE float32 = 16;
const FULL_TILE float32 = 32;

//generates tiles in a box grid
func GenerateTilesFromGrid(rows int, columns int, _pivot rl.Vector2) *PlotTiles{
	_plots := make([]*Plot, rows * columns);
	k := 0;
	for i:=0; i<rows; i++{
		for j:=0; j<columns; j++{
			_plots[k] = NewPlot(rl.NewVector2(_pivot.X + HALF_TILE + float32(i)*FULL_TILE, _pivot.Y + HALF_TILE + float32(j)*FULL_TILE));
			k++;
		}
	}

	return &PlotTiles{plots: _plots, pivot: _pivot};
}

//generates tiles from a list of vector locations
func GenerateTilesFromVectorArray(locations []rl.Vector2) *PlotTiles{
	_plots := make([]*Plot, len(locations))
	for i, locat := range locations{
		_plots[i] = NewPlot(locat);
	}

	return &PlotTiles{plots: _plots, pivot: rl.NewVector2(0,0)};
}

func (tiles PlotTiles)DrawTiles(){
	for i:=0; i<len(tiles.plots); i++{
		tiles.plots[i].DrawPlot();
	}
}

func (gm *GameManager) PlantPlot(plot *Plot){
	if(gm.GetSelectedItem != nil){
		gm.GetSelectedItem().Quantity--;
	switch(gm.GetSelectedItem().Name){
	case CARROT_SEED:
		plot.AddPlant(NewCarrotPlant());
		fmt.Println("Planted Carrot");
	}
	}
	
}

func (gm *GameManager)UpdatePlot(){
	DrawDirt();

	mouseLocation := rl.GetMousePosition();
	for _, plot := range gm.plots{
		plot.DrawPlot();

		if((mouseLocation.X >= plot.location.X + gm.pivot.X - 16 && mouseLocation.X <= plot.location.X + 16 + gm.pivot.X)&&(mouseLocation.Y >= plot.location.Y + gm.pivot.Y - 16 && mouseLocation.Y <= plot.location.Y + 16 + gm.pivot.Y)){
			plot.isMouseOver = true
		}else{
			plot.isMouseOver = false
		}
		if(rl.IsMouseButtonPressed(rl.MouseButtonLeft) && plot.isMouseOver){
			if(plot.HasPlant()){
				gm.Harvest(plot);
			}else{
				gm.PlantPlot(plot)
			}
		}
	}
}

func DrawDirt(){
	rl.DrawRectangle(64, 64, 672, 320, rl.Brown)
}