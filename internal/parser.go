package internal

import (
	"github.com/adrg/frontmatter"
	"os"
	"strings"
)

func ParseMetadata(path string) (Meta, string, error) {
	var meta Meta

	f, err := os.Open(path)
	if err != nil {
		return meta, "", err
	}
	defer f.Close()

	content, err := frontmatter.Parse(f, &meta)
	if err != nil {
		return meta, "", err
	}

	return meta, string(content), nil
}

func ParseSlides(content string) []string {
	lines := strings.Split(content, "\n")

	var slides []string
	var current []string

	for _, line := range lines {
		if strings.TrimSpace(line) == "---" {
			slide := strings.TrimSpace(strings.Join(current, "\n"))
			if slide != "" {
				slides = append(slides, slide)
			}
			current = []string{}
			continue
		}
		current = append(current, line)
	}

	// Add the final slide to the array
	final := strings.TrimSpace(strings.Join(current, "\n"))
	if final != "" {
		slides = append(slides, final)
	}

	return slides
}
