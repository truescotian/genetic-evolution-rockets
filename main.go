package main

import (
	"github.com/zacg/boids"
)

const (
	iterations int = 50000
	N          int = 100
	lifetime   int = 500

	mutationRate    float64 = 0.01
	populationCount int     = 50
)

var target boids.PVector

func main() {
	target = boids.NewRandom2dPVector()
	population := NewPopulation(mutationRate, populationCount)
	lifeCounter := 0

	for i := 0; i < iterations; i++ {
		if lifeCounter < lifetime {
			population.live()
			lifeCounter++
		} else {
			lifeCounter = 0
			population.fitness()
			population.selection()
			population.reproduction()
		}
	}
}
