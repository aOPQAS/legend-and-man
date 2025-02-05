// Фабричный метод

package main

import "fmt"

type Drug interface {
	setName(name string)
	setDose(dose int)
	getName() string
	getDose() int
}

type SyntheticOpioids struct {
	name string
	dose int
}

func (s *SyntheticOpioids) setName(name string) {
	s.name = name
}

func (s *SyntheticOpioids) getName() string {
	return s.name
}

func (s *SyntheticOpioids) setDose(dose int) {
	s.dose = dose
}

func (s *SyntheticOpioids) getDose() int {
	return s.dose
}

type Fentanyl struct {
	SyntheticOpioids
}

func newFentanyl() Drug {
	return &Fentanyl{
		SyntheticOpioids: SyntheticOpioids{
			name: "Fentanyl",
			dose: 1,
		},
	}
}

type Methadone struct {
	SyntheticOpioids
}

func newMethadone() Drug {
	return &Methadone{
		SyntheticOpioids: SyntheticOpioids{
			name: "Methadone",
			dose: 30,
		},
	}
}

func getSyntheticOpioids(title string) (Drug, error) {
	if title == "Fentanyl" {
		return newFentanyl(), nil
	}
	if title == "Methadone" {
		return newMethadone(), nil
	}

	return nil, fmt.Errorf("Unknown drug: %s", title)
}

func main() {
	fentanyl, _ := getSyntheticOpioids("Fentanyl")
	methadone, _ := getSyntheticOpioids("Methadone")

	printDetails(fentanyl)
	printDetails(methadone)
}

func printDetails(d Drug) {
	fmt.Printf("Drug: %s\n", d.getName())
	fmt.Printf("Dose: %d mg\n\n", d.getDose())
}
