package main

import (
	"fmt"
	"os"
)

func ParseArgs() (string, string, error) {
	if len(os.Args) != 5 {
		return "", "", fmt.Errorf("failed to parse args, must be 3 values")
	}

	if os.Args[1] != "--old" {
		return "", "", fmt.Errorf("missed option --old")
	}

	if os.Args[3] != "--new" {
		return "", "", fmt.Errorf("missed option --new")
	}

	return os.Args[2], os.Args[4], nil
}
