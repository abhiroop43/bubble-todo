package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func initialModel() model {
	return model{
		choices:  []string{"Practice Coding", "Exercise", "Read a Book"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "j", "down":
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}
		case "k", "up":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}
		case "enter":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	viewModel := "\nWhat do you plan to do today?\n\n"

	for i, choice := range m.choices {
		cursor := " "

		if m.cursor == i {
			cursor = ">"
		}

		selected := " "
		if _, ok := m.selected[i]; ok {
			selected = "x"
		}

		viewModel += fmt.Sprintf("%s [%s] %s\n", cursor, selected, choice)
	}

	viewModel += "\nPress q to quit.\n"

	return viewModel
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("An error has occurred: %v", err)
		os.Exit(1)
	}
}
