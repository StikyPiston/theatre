package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/stikypiston/theatre/internal"
)

type model struct {
	slides       []string
	currentSlide int
	meta         internal.Meta

	width  int
	height int

	renderer *glamour.TermRenderer
}

func newModel(slides []string, meta internal.Meta) model {
	r, _ := glamour.NewTermRenderer(glamour.WithAutoStyle())

	return model{
		slides:   slides,
		meta:     meta,
		renderer: r,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "right", "l", " ":
			if m.currentSlide < len(m.slides)-1 {
				m.currentSlide++
			}

		case "left", "h":
			if m.currentSlide > 0 {
				m.currentSlide--
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.width == 0 || m.height == 0 {
		return "Loading..."
	}

	contentHeight := m.height - 1 // leave 1 row for status bar

	rendered, _ := m.renderer.Render(m.slides[m.currentSlide])

	lines := strings.Split(rendered, "\n")
	if len(lines) > contentHeight {
		lines = lines[:contentHeight]
	}
	content := strings.Join(lines, "\n")

	// pad with blank lines if slide shorter than available height
	for len(lines) < contentHeight {
		content += "\n"
		lines = append(lines, "")
	}

	return content + "\n" + m.statusBar()
}

var statusStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#11111b")). // gray
	Foreground(lipgloss.Color("#ffffff")).
	PaddingLeft(1).
	PaddingRight(1)

func (m model) statusBar() string {
	// Left part: Title | Author
	left := fmt.Sprintf("%s | %s", m.meta.Title, m.meta.Author)

	// Right part: current / total
	right := fmt.Sprintf("%d/%d", m.currentSlide+1, len(m.slides))

	// Make space between left/right
	space := m.width - lipgloss.Width(left) - lipgloss.Width(right) - 2
	if space < 0 {
		space = 0
	}

	bar := left + strings.Repeat(" ", space) + right
	return statusStyle.Render(bar)
}
