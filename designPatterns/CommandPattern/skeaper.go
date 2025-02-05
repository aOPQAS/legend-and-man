package main

import "fmt"

type Speaker struct {
	run bool
}

func (s *Speaker) on() {
	s.run = true
	fmt.Println("Music is playing")
}

func (s *Speaker) off() {
	s.run = false
	fmt.Println("Music is not playing")
}
