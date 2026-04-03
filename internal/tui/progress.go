package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type ProgressMsg float64

type ProgressModel struct {
	pw       int
	progress progress.Model
	percent  float64
	label    string
}

func NewProgressModel(label string) ProgressModel {
	return ProgressModel{
		progress: progress.New(progress.WithDefaultGradient()),
		label:    label,
	}
}

func (m ProgressModel) Init() tea.Cmd { return nil }

func (m ProgressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	case progress.FrameMsg:
		newModel, cmd := m.progress.Update(msg)
		m.progress = newModel.(progress.Model)
		return m, cmd
	case ProgressMsg:
		m.percent = float64(msg)
		if m.percent >= 1.0 {
			return m, tea.Quit
		}
		return m, nil
	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - 10
		return m, nil
	default:
		return m, nil
	}
}

func (m ProgressModel) View() string {
	pad := strings.Repeat(" ", 2)
	return "\n" +
		pad + m.label + "\n\n" +
		pad + m.progress.ViewAs(m.percent) + "\n\n"
}
