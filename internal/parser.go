package internal

import (
	"github.com/adrg/frontmatter"
	"os"
)

func ParseFile(path string) (Meta, string, error) {
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
