package internal

import (
	"os"
	"path/filepath"
)

func LoadSlide(path string) (Meta, []string, error) {
	meta, _, err := ParseMetadata(path)
	if err != nil {
		return meta, []string{}, err
	}

	slidespath, err := filepath.Abs(path)
	if err != nil {
		return meta, []string{}, err
	}

	filecontent, err := os.ReadFile(slidespath)
	if err != nil {
		return meta, []string{}, err
	}

	slidescontent := ParseSlides(string(filecontent))

	return meta, slidescontent, err
}
