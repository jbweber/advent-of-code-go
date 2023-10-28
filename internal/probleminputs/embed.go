package probleminputs

import (
	"embed"
	"fmt"
)

var (
	//go:embed resources/*
	resources embed.FS
)

func GetInput(year, day int) (string, error) {
	fb, err := resources.ReadFile(fmt.Sprintf("resources/%d/%d/input.txt", year, day))
	if err != nil {
		return "", nil
	}

	return string(fb), nil
}

func GetSample(year, day int) (string, error) {
	fb, err := resources.ReadFile(fmt.Sprintf("resources/%d/%d/sample.txt", year, day))
	if err != nil {
		return "", nil
	}

	return string(fb), nil
}
