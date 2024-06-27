package boids

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	position   Vector
	velocity   Vector
	viewRadius float64
	id         int
}

func (b *Boid) Position() Vector {
	return b.position
}

func (b *Boid) Velocity() Vector {
	return b.velocity
}

func (b *Boid) calcAcceleration(field *Field) Vector {
	upper, lower := b.Position().AddV(b.viewRadius), b.Position().AddV(-b.viewRadius)
	avgPosition, avgVelocity, separation := Vector{0, 0}, Vector{0, 0}, Vector{0, 0}
	count := 0.0
	field.rWlock.RLock()
	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, field.screenWidth); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, field.screenHeight); j++ {
			if otherBoidId := field.BoidsMap[int(j)][int(i)]; otherBoidId != -1 && otherBoidId != b.id {
				if dist := field.BoildsList[otherBoidId].Position().Distance(b.position); dist < b.viewRadius {
					count++
					avgVelocity = avgVelocity.Add(field.BoildsList[otherBoidId].Velocity())
					avgPosition = avgPosition.Add(field.BoildsList[otherBoidId].Position())
					separation = separation.Add(b.Position().Subtract(field.BoildsList[otherBoidId].Position()).DivisionV(dist))
				}
			}
		}
	}
	field.rWlock.RUnlock()

	accel := Vector{b.borderBounce(b.Position().X(), field.screenWidth), b.borderBounce(b.Position().Y(), field.screenHeight)}
	if count > 0 {
		avgPosition, avgVelocity = avgPosition.DivisionV(count), avgVelocity.DivisionV(count)
		accelAlignment := avgVelocity.Subtract(b.velocity).MultiplyV(field.adjRate)
		accelCohesion := avgPosition.Subtract(b.position).MultiplyV(field.adjRate)
		accelSeparation := separation.MultiplyV(field.adjRate)
		accel = accel.Add(accelAlignment).Add(accelCohesion).Add(accelSeparation)
	}
	return accel
}

func (b *Boid) borderBounce(pos, maxBorderPos float64) float64 {
	if pos < b.viewRadius {
		return 1 / pos
	} else if pos > maxBorderPos-b.viewRadius {
		return 1 / (pos - maxBorderPos)
	}
	return 0
}

func (b *Boid) moveOne(field *Field) {
	acceleration := b.calcAcceleration(field)
	field.rWlock.Lock()
	b.velocity = b.velocity.Add(acceleration).limit(-1, 1)
	field.BoidsMap[int(b.Position().Y())][int(b.Position().X())] = -1
	b.position = b.position.Add(b.velocity)
	field.BoidsMap[int(b.Position().Y())][int(b.Position().X())] = b.id
	field.rWlock.Unlock()
}

func (b *Boid) start(field *Field) {
	for {
		b.moveOne(field)
		time.Sleep(5 * time.Millisecond)
	}
}

func createBoid(bid int, field *Field) {
	b := Boid{
		position:   Vector{rand.Float64() * field.screenWidth, rand.Float64() * field.screenHeight},
		velocity:   Vector{(rand.Float64() * 2) - 1.0, (rand.Float64() * 2) - 1.0},
		viewRadius: field.viewRadius,
		id:         bid,
	}

	field.BoildsList[bid] = &b
	field.BoidsMap[int(b.Position().Y())][int(b.Position().X())] = b.id
	go b.start(field)
}
