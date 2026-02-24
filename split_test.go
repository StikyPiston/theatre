package main

import (
	"github.com/stikypiston/theatre/internal"
	"testing"
)

func TestSplitSlides(t *testing.T) {
	content := `
# Slide One

Hello

---

# Slide Two

World
`

	slides := internal.ParseSlides(content)

	if len(slides) != 2 {
		t.Fatalf("expected 2 slides, got %d", len(slides))
	}

	if slides[0] != "# Slide One\n\nHello" {
		t.Errorf("unexpected first slide:\n%s", slides[0])
	}

	if slides[1] != "# Slide Two\n\nWorld" {
		t.Errorf("unexpected second slide:\n%s", slides[1])
	}
}
