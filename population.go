package main

import (
	"fmt"
	"github.com/zacg/boids"

	"math/rand"
)

type Population struct {
	mutationRate float64
	population   []Rocket
	matingPool   []Rocket
	generations  int
}

func NewPopulation(mutationRate float64, populationCount int) *Population {

	p := &Population{
		mutationRate: mutationRate,
		population:   make([]Rocket, populationCount),
		generations:  0,
	}
	for i := 0; i < populationCount; i++ {
		p.population[i] = NewRocket()
	}

	return p
}

func (p *Population) live() {
	for i := 0; i < len(p.population); i++ {
		p.population[i].run()
	}
}

func (p *Population) fitness() {
	for i := 0; i < len(p.population); i++ {
		p.population[i].setFitness()
	}
}

func (p *Population) selection() {
	// Mating pool
	// we want to make it so there is a percent score
	// is proportional to the amount they would get picked
	// from a basket of potential parents.
	// So if 'a' has a 5% of being picked, we want
	// 'a' to be in the pool 5% of the time, which is N=fitness*100
	p.matingPool = nil
	p.matingPool = make([]Rocket, 0)

	var maxFitness = p.getMaxFitness()
	var minFitness = p.getMinFitness()

	for i := 0; i < len(p.population); i++ {
		fitnessNormal := normalize(p.population[i].fitness, maxFitness, minFitness)
		n := int(fitnessNormal * 100)
		fmt.Println(n)
		for j := 0; j < n; j++ {
			p.matingPool = append(p.matingPool, p.population[i])
		}
	}
}

// Find highest fintess for the population
func (p *Population) getMaxFitness() float64 {
	record := p.population[0].fitness
	for i := 0; i < len(p.population); i++ {
		if p.population[i].fitness > record {
			record = p.population[i].fitness
		}
	}
	return record
}

// Find highest fintess for the population
func (p *Population) getMinFitness() float64 {
	record := p.population[0].fitness
	for i := 0; i < len(p.population); i++ {
		if p.population[i].fitness < record {
			record = p.population[i].fitness
		}
	}
	return record
}

func normalize(currValue, maxValue, minValue float64) float64 {
	return (currValue - minValue) / (maxValue - minValue)
}

func (p *Population) reproduction() {
	for i := 0; i < len(p.population); i++ {
		a := rand.Intn(len(p.matingPool))
		b := rand.Intn(len(p.matingPool))

		mom := p.matingPool[a]
		dad := p.matingPool[b]

		momGenes := mom.getDNA()
		dadGenes := dad.getDNA()

		child := momGenes.crossover(dadGenes)

		child.mutate()

		// overwrite the population with new child
		location := boids.NewPVector2D(width/2, height+20)
		p.population[i] = NewRocketFromChild(location, child)
	}
	p.generations++
}
