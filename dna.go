package main

import (
	"github.com/zacg/boids"
	"math"
	"math/rand"
	"time"
)

const (
	maxForce float64 = 0.1 // strength of thruster
)

type DNA struct {
	genes []boids.PVector
}

func NewDNAFromChild(dna []boids.PVector) DNA {
	d := DNA{
		genes: dna,
	}
	return d
}

func NewDNA() (dna DNA) {

	dna = DNA{
		genes: make([]boids.PVector, lifetime),
	}

	for i := 0; i < len(dna.genes); i++ {
		dna.genes[i] = boids.NewRandom2dPVector()
		rand.Seed(time.Now().UnixNano())
		dna.genes[i].Mult(rand.Float64() * maxForce)
	}

	return dna
}

func (dna *DNA) crossover(partner DNA) DNA {
	child := make([]boids.PVector, len(dna.genes))

	// Pick a midpoint
	crossover := rand.Intn(len(dna.genes))

	// Take half from one and half from the otehr
	for i := 0; i < len(dna.genes); i++ {
		if i > crossover {
			child[i] = dna.genes[i]
		} else {
			child[i] = partner.genes[i]
		}
	}
	newGenes := NewDNAFromChild(child)
	return newGenes
}

func (dna *DNA) mutate() {
	for i := 0; i < len(dna.genes); i++ {
		if rand.Float64() < mutationRate {
			rand.Seed(time.Now().UnixNano())
			angle := rand.Float64() * (math.Pi * 2)
			dna.genes[i] = boids.NewPVector2D(math.Cos(angle), math.Sin(angle))
			dna.genes[i].Mult(rand.Float64() * maxForce)
		}
	}
}
