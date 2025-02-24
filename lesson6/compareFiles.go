package main

import (
	"fmt"
)

func compareFiles(oldFile, newFile string) error {
	Oldsnapshot, err := ReadFile(oldFile)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	Newsnapshot, err := ReadFile(newFile)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	for _, Newsnapshot := range Newsnapshot {
		isAdded := true
		for _, Oldsnapshot := range Oldsnapshot {
			if Newsnapshot == Oldsnapshot {
				isAdded = false
				break
			}
		}
		if isAdded {
			fmt.Println("ADDED", Newsnapshot)
		}
	}

	for _, Oldsnapshot := range Oldsnapshot {
		isAdded := true
		for _, Newsnapshot := range Newsnapshot {
			if Oldsnapshot == Newsnapshot {
				isAdded = false
				break
			}
		}
		if isAdded {
			fmt.Println("REMOVED", Oldsnapshot)
		}
	}
	return nil
}
