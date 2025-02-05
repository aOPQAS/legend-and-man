// Стратегия

package main

import "fmt"

type MovementStrategy interface {
	move()
}

type AggressiveMoment struct{}

func (a *AggressiveMoment) move() {
	fmt.Println("An ally runs towards the enemy")
}

type RetreatMoment struct{}

func (r *RetreatMoment) move() {
	fmt.Println("An ally moves away from the enemy")
}

type Ally struct {
	strategy MovementStrategy
}

func (a *Ally) SetStrategy(strategy MovementStrategy) {
	a.strategy = strategy
}

func (a *Ally) Move() {
	a.strategy.move()
}

func main() {
	ally := &Ally{}

	ally.SetStrategy(&AggressiveMoment{})
	ally.Move()

	ally.SetStrategy(&RetreatMoment{})
	ally.Move()
}
