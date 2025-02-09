package main

import "fmt"

type Computer struct {
	CPU       string
	RAM       string
	Storage   string
	VideoCard string
}

type ComputerBuilder interface {
	SetCPU(cpu string)
	SetRAM(ram string)
	SetStorage(storage string)
	SetVideoCard(videocard string)
	GetComputer() *Computer
}

type ConcreteComputerBuilder struct {
	computer *Computer
}

func (c *ConcreteComputerBuilder) SetCPU(cpu string) {
	c.computer.CPU = cpu
}

func (c *ConcreteComputerBuilder) SetRAM(ram string) {
	c.computer.RAM = ram
}

func (c *ConcreteComputerBuilder) SetStorage(storage string) {
	c.computer.Storage = storage
}

func (c *ConcreteComputerBuilder) SetVideoCard(videocard string) {
	c.computer.VideoCard = videocard
}

func (c *ConcreteComputerBuilder) GetComputer() *Computer {
	return c.computer
}

type ComputerDirector struct {
	builder ComputerBuilder
}

func (d *ComputerDirector) Create() {
	d.builder.SetCPU("9 9900X3D")
	d.builder.SetRAM("128GB")
	d.builder.SetStorage("16TB SSD")
	d.builder.SetVideoCard("NVIDIA хуикс 9090")
}

func main() {
	builder := &ConcreteComputerBuilder{computer: &Computer{}}
	director := &ComputerDirector{builder: builder}

	director.Create()

	computer := builder.GetComputer()
	fmt.Println("The created computer:")
	fmt.Println("CPU:", computer.CPU)
	fmt.Println("RAM:", computer.RAM)
	fmt.Println("Storage:", computer.Storage)
	fmt.Println("Graphics:", computer.VideoCard)
}
