package main

import (
	"github.com/zacg/boids"

	"math"
	"math/rand"
	"time"
)

const (
	width  float64 = 50.00
	height float64 = 50.00
)

type Rocket struct {
	location     boids.PVector
	velocity     boids.PVector
	acceleration boids.PVector

	// Size
	r float64

	fitness     float64
	dna         DNA
	geneCounter int
}

func NewRocket() Rocket {

	return Rocket{
		dna:         NewDNA(),
		geneCounter: 0,
		location:    boids.NewPVector2D(width/2, height+20),
	}
}

func NewRocketFromChild(l boids.PVector, child DNA) (rocket Rocket) {
	rocket.acceleration = boids.NewPVector2D(0, 0)
	rocket.velocity = boids.NewRandom2dPVector()
	rocket.location = l
	rocket.r = 4
	rocket.dna = child
	return rocket
}

// Fitness is our fitness function
// INITIALLY WE USED:
// by using 1 divided by distance, large distances become small
// numbers and small distances become large.
// So the further away the rocket is from the target, the lower the fitness,
// and the closer the rocket is from the target, the higher the fitness
func (r *Rocket) setFitness() {
	d := r.location.Dist(target) // how close did we get?
	r.fitness = math.Pow(1/d, 2) // Squaring 1 divided by distance
}

func (r *Rocket) run() {
	r.applyForce(r.dna.genes[r.geneCounter])
	r.geneCounter = (r.geneCounter + 1) % len(r.dna.genes)
	r.update()
}

func (r *Rocket) applyForce(v boids.PVector) {
	r.acceleration.Add(v)
}

// Update is our physical model (Euler integration)
func (r *Rocket) update() {
	r.velocity.Add(r.acceleration) // velocity changes according to acceleration
	r.location.Add(r.velocity)     // location changes according to velocity
	r.acceleration.Mult(0)
}

// crossover Creates a child using random midpoint method of crossover
// which is the first section of genes taken from parent A, and the second
// section from parent B.
func (r *Rocket) crossover(partner Rocket) (child Rocket) {
	child = NewRocket()

	// get a random index from genes as a midpoint
	midPoint := rand.Intn(len(r.dna.genes))

	for i := 0; i < len(r.dna.genes); i++ {
		// before midpoint copy genes from A, after midpoint copy from B
		if i > midPoint {
			child.dna.genes[i] = r.dna.genes[i]
		} else {
			child.dna.genes[i] = partner.dna.genes[i]
		}
	}

	return child
}

func (r *Rocket) mutate() {
	for i := 0; i < len(r.dna.genes); i++ {
		rand.Seed(time.Now().UnixNano())
		if rand.Float64() < mutationRate {
			r.dna.genes[i] = boids.NewRandom2dPVector()
		}
	}
}

func (r *Rocket) getDNA() DNA {
	return r.dna
}
