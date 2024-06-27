package main

import (
	"github.com/FluffyFoxTail/boids/boids"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
	"sync"
)

// TODO replace with flag
const (
	screenWidth, screenHeight = 640, 360
	boidCount                 = 500
	viewRadius                = 13
	adjRate                   = 0.015
)

type Game struct {
	*boids.Field
	boidColor color.RGBA
}

func NewGame(field *boids.Field, boidColor color.RGBA) *Game {
	return &Game{Field: field, boidColor: boidColor}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, boid := range g.Field.BoildsList {
		screen.Set(int(boid.Position().X()+1), int(boid.Position().Y()), g.boidColor)
		screen.Set(int(boid.Position().X()-1), int(boid.Position().Y()), g.boidColor)
		screen.Set(int(boid.Position().X()), int(boid.Position().Y()-1), g.boidColor)
		screen.Set(int(boid.Position().X()), int(boid.Position().Y()+1), g.boidColor)
	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

func main() {
	boidsList := make([]*boids.Boid, boidCount)
	boidMap := make([][]int, screenHeight)

	// TODO update X and Y axis
	for i := range boidMap {
		boidMap[i] = make([]int, screenWidth)
		for j := range boidMap[i] {
			boidMap[i][j] = -1
		}
	}

	field := boids.NewField(screenWidth, screenHeight, viewRadius, adjRate, sync.RWMutex{}, boidsList, boidMap)
	boids.InitField(field)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Boids")
	if err := ebiten.RunGame(NewGame(field, color.RGBA{R: 10, G: 255, B: 50, A: 255})); err != nil {
		log.Fatal(err)
	}
}
