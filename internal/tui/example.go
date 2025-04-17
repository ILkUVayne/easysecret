package tui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type ExampleModel struct {
	count int
}

func (m ExampleModel) Init() tea.Cmd {
	return nil
}

func (m ExampleModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up":
			m.count++
		}
	}
	return m, nil
}

func (m ExampleModel) View() string {
	return fmt.Sprintf("Count: %d", m.count)
}
