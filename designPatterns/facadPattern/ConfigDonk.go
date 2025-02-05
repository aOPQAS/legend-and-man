// фасад
package main

import "fmt"

type Computer struct {
	fps        int
	resolution string
	graphics   string
}

func (c *Computer) Setup() {
	fmt.Println("Setting up your computer settings:")
	fmt.Println("FPS:", c.fps)
	fmt.Println("Permission:", c.resolution)
	fmt.Println("Graphics:", c.graphics)
}

type Volume struct {
	SoundInTheLobby int
	FullSound       int
	SoundInGame     int
}

func (v *Volume) Setup() {
	fmt.Println("Setting Sound Settings:")
	fmt.Println("Sound in the lobby:", v.SoundInTheLobby)
	fmt.Println("General sound:", v.FullSound)
	fmt.Println("Sound in the game", v.SoundInGame)
}

type Sensa struct {
	DPI                  int
	Sensitivity          int
	Hz                   int
	ProximitySensitivity int
	SensitivityInLinux   int
	MouseAcceleration    int
	RawInput             int
	eDPI                 int
}

func (s *Sensa) Setup() {
	fmt.Println("Setting mouse sensitivity:")
	fmt.Println("DPI:", s.DPI)
	fmt.Println("Sensitivity:", s.Sensitivity)
	fmt.Println("Hz:", s.Hz)
	fmt.Println("ProximitySensitivity:", s.ProximitySensitivity)
	fmt.Println("SensitivityInLinux:", s.SensitivityInLinux)
	fmt.Println("MouseAcceleration:", s.MouseAcceleration)
	fmt.Println("RawInput:", s.RawInput)
	fmt.Println("eDPI:", s.eDPI)
}

type GamerSetup struct {
	computer *Computer
	volume   *Volume
	sensa    *Sensa
}

func (gs *GamerSetup) SetupAll() {
	gs.computer.Setup()
	gs.volume.Setup()
	gs.sensa.Setup()
	fmt.Println("Setup complete!")
}

func main() {
	computer := &Computer{fps: 144, resolution: "1280x960", graphics: "High"}
	volume := &Volume{SoundInTheLobby: 15, FullSound: 65, SoundInGame: 100}
	sensa := &Sensa{DPI: 800, Sensitivity: 2, Hz: 1000, ProximitySensitivity: 1, SensitivityInLinux: 6, MouseAcceleration: 0, RawInput: 1, eDPI: 1000}

	gamerSetup := &GamerSetup{computer: computer, volume: volume, sensa: sensa}

	gamerSetup.SetupAll()
}
