package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stikypiston/theatre/internal"
)

func main() {
	// Ensure a file path was provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: theatre <file.md>")
		os.Exit(1)
	}

	filepath := os.Args[1]

	// Load slides + metadata
	meta, slides, err := internal.LoadSlide(filepath)
	if err != nil {
		log.Fatalf("failed to load slides: %v", err)
	}

	if len(slides) == 0 {
		log.Fatal("no slides found in file")
	}

	// Create model
	m := newModel(slides, meta)

	// Start fullscreen Bubble Tea program
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatalf("error running theatre: %v", err)
	}
}
