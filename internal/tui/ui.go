package tui

import (
	"fmt"
	"kcm-cli/internal/core"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	docStyle    = lipgloss.NewStyle().Margin(1, 2)
	titleStyle  = lipgloss.NewStyle().Background(lipgloss.Color("62")).Foreground(lipgloss.Color("230")).Padding(0, 1)
	statusStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).MarginTop(1)
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type state int

const (
	stateList state = iota
	stateNew  // Criando novo snapshot
)

type model struct {
	list      list.Model
	textInput textinput.Model
	curState  state
	statusMsg string
	err       error
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.curState == stateNew {
			switch msg.String() {
			case "enter":
				name := m.textInput.Value()
				if name != "" {
					m.statusMsg = fmt.Sprintf("Criando snapshot '%s'...", name)
					m.curState = stateList
					err := core.SaveSnapshot(name)
					if err != nil {
						m.err = err
						return m, nil
					}
					m.textInput.Reset()
					// Atualiza a lista e captura o comando
					m.list, cmd = m.updateListItems()
					return m, cmd
				}
			case "esc":
				m.curState = stateList
				m.textInput.Reset()
				return m, nil
			}
			m.textInput, cmd = m.textInput.Update(msg)
			return m, cmd
		}

		// Atalhos Globais na Lista
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "n":
			m.curState = stateNew
			m.textInput.Focus()
			return m, nil
		case "d", "x":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				err := core.DeleteProfile(i.title)
				if err != nil {
					m.err = err
					return m, nil
				}
				m.statusMsg = fmt.Sprintf("Perfil '%s' deletado.", i.title)
				m.list, cmd = m.updateListItems()
				return m, cmd
			}
		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.statusMsg = fmt.Sprintf("Aplicando '%s'...", i.title)
				err := core.ApplySnapshot(i.title)
				if err != nil {
					m.err = err
					return m, nil
				}
				m.statusMsg = fmt.Sprintf("Snapshot '%s' aplicado!", i.title)
				// Atualiza a lista para mostrar o novo rollback criado
				m.list, cmd = m.updateListItems()
				return m, cmd
			}
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	if m.curState == stateList {
		m.list, cmd = m.list.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m model) updateListItems() (list.Model, tea.Cmd) {
	profiles, _ := core.ListProfiles()
	items := []list.Item{}
	for _, p := range profiles {
		items = append(items, item{
			title: p.Name,
			desc:  fmt.Sprintf("Tema: %s | Criado: %s", p.GlobalTheme, p.CreatedAt.Format("02/01/2006 15:04")),
		})
	}
	cmd := m.list.SetItems(items)
	return m.list, cmd
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("\n  Erro: %v\n\n  Pressione 'q' para sair.", m.err)
	}

	if m.curState == stateNew {
		return docStyle.Render(fmt.Sprintf(
			"Digite o nome do novo Snapshot:\n\n%s\n\n(ENTER para salvar, ESC para cancelar)",
			m.textInput.View(),
		))
	}

	help := "n: novo | d/x: deletar | enter: aplicar | q: sair"
	return docStyle.Render(
		m.list.View() + "\n" +
			statusStyle.Render(m.statusMsg) + "\n" +
			statusStyle.Render(help),
	)
}

func StartUI() error {
	profiles, err := core.ListProfiles()
	if err != nil {
		return err
	}

	items := []list.Item{}
	for _, p := range profiles {
		items = append(items, item{
			title: p.Name,
			desc:  fmt.Sprintf("Tema: %s | Criado: %s", p.GlobalTheme, p.CreatedAt.Format("02/01/2006 15:04")),
		})
	}

	ti := textinput.New()
	ti.Placeholder = "Meu Perfil Dark"
	ti.CharLimit = 30
	ti.Width = 30

	m := model{
		list:      list.New(items, list.NewDefaultDelegate(), 0, 0),
		textInput: ti,
		curState:  stateList,
	}
	m.list.Title = "KCM-CLI: Gerenciador de Perfis KDE"

	p := tea.NewProgram(m, tea.WithAltScreen())
	_, err = p.Run()
	return err
}
