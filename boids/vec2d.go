package boids

import "math"

type Vector struct {
	x float64
	y float64
}

func NewVector2D(x, y float64) Vector {
	return Vector{x, y}
}

func (v Vector) X() float64 { return v.x }

func (v Vector) Y() float64 { return v.y }

func (v Vector) Add(v2 Vector) Vector {
	return Vector{v.x + v2.x, v.y + v2.y}
}

func (v Vector) Subtract(v2 Vector) Vector {
	return Vector{v.x - v2.x, v.y - v2.y}
}

func (v Vector) Multiply(v2 Vector) Vector {
	return Vector{v.x * v2.x, v.y * v2.y}
}

func (v Vector) AddV(d float64) Vector {
	return Vector{v.x + d, v.y + d}
}

func (v Vector) MultiplyV(d float64) Vector {
	return Vector{v.x * d, v.y * d}
}

func (v Vector) DivisionV(d float64) Vector {
	return Vector{v.x / d, v.y / d}
}

func (v Vector) limit(lower, upper float64) Vector {
	return Vector{math.Min(math.Max(v.x, lower), upper),
		math.Min(math.Max(v.y, lower), upper)}
}

func (v Vector) Distance(v2 Vector) float64 {
	return math.Sqrt(math.Pow(v.x-v2.x, 2) + math.Pow(v.y-v2.y, 2))
}
