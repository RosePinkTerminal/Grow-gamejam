package main

import(
	rl "github.com/gen2brain/raylib-go/raylib"
	"fmt"
)

type PlotTiles struct{
	plots []*Plot
}

const HALF_TILE float32 = 16;
const FULL_TILE float32 = 32;

//generates tiles in a box grid
func GenerateTilesFromGrid(rows int, columns int, pivot rl.Vector2) PlotTiles{
	_plots := make([]*Plot, rows * columns);
	k := 0;
	for i:=0; i<rows; i++{
		for j:=0; j<columns; j++{
			_plots[k] = NewPlot(rl.NewVector2(pivot.X + HALF_TILE + float32(i)*FULL_TILE, pivot.Y + HALF_TILE + float32(j)*FULL_TILE));
			fmt.Println("plot: ", _plots[i].location.Y, ", ", _plots[i].location.Y);
			k++;
		}
	}

	return PlotTiles{plots: _plots};
}

//generates tiles from a list of vector locations
func GenerateTilesFromVectorArray(locations []rl.Vector2) PlotTiles{
	_plots := make([]*Plot, len(locations))
	for i, locat := range locations{
		_plots[i] = NewPlot(locat);
	}

	return PlotTiles{plots: _plots};
}

func (tiles PlotTiles)DrawTiles(){
	for i:=0; i<len(tiles.plots); i++{
		tiles.plots[i].DrawPlot();
	}
}