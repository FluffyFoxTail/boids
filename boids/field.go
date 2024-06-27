package boids

import (
	"sync"
)

type Field struct {
	screenWidth, screenHeight float64
	viewRadius                float64
	adjRate                   float64
	rWlock                    sync.RWMutex
	BoildsList                []*Boid
	BoidsMap                  [][]int
}

func NewField(screenWidth float64,
	screenHeight float64,
	viewRadius float64,
	adjRate float64,
	rWlock sync.RWMutex,
	boildsList []*Boid,
	boidsMap [][]int) *Field {
	return &Field{screenWidth: screenWidth, screenHeight: screenHeight, viewRadius: viewRadius, adjRate: adjRate, rWlock: rWlock, BoildsList: boildsList, BoidsMap: boidsMap}
}

func InitField(field *Field) {
	for i := 0; i < len(field.BoildsList); i++ {
		createBoid(i, field)
	}
}
