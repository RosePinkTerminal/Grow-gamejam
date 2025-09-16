package main

import(
	rl "github.com/gen2brain/raylib-go/raylib"
)

//better draw texture function
func DrawTextureEz(texture rl.Texture2D, pos rl.Vector2, angle float32, scale float32, color rl.Color) {
    sourceRect := rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height))
    destRect := rl.NewRectangle(pos.X, pos.Y, float32(texture.Width)*scale, float32(texture.Height)*scale)
    origin := rl.Vector2Scale(rl.NewVector2(float32(texture.Width)/2, float32(texture.Height)/2), scale)
    rl.DrawTexturePro(texture, sourceRect,
        destRect,
        origin, angle, color)
}