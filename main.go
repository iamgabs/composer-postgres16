package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list         list.Model
	currentState string
	selectedItem string
	width        int
	height       int
}

func initialModel() model {
	dbs := []string{"PostgreSQL", "MySQL", "MariaDB", "MongoDB", "Oracle", "SQL Server"}
	m := model{
		currentState: "databases",
		width:        80,
		height:       20,
	}

	items := make([]list.Item, len(dbs))
	for i, db := range dbs {
		items[i] = item{title: db, desc: "Database: " + db}
	}

	// Criar a lista com dimensões apropriadas
	delegate := list.NewDefaultDelegate()
	m.list = list.New(items, delegate, m.width, m.height)
	m.list.Title = "Selecione um Banco de Dados"
	m.list.SetShowStatusBar(false)
	m.list.SetFilteringEnabled(false)
	m.list.Styles.Title = lipgloss.NewStyle().MarginLeft(2).Bold(true)

	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

type versionSelectionMsg struct {
	versions []list.Item
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		m.list.SetSize(msg.Width-4, msg.Height-4)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if m.currentState == "databases" {
				if len(m.list.Items()) > 0 {
					selected := m.list.SelectedItem().(item)
					m.selectedItem = selected.title
					return m, loadVersions(m)
				}
			} else if m.currentState == "versions" {
				if len(m.list.Items()) > 0 {
					selected := m.list.SelectedItem().(item)
					return m, startDockerCompose(selected.title)
				}
			}
		case "esc":
			if m.currentState == "versions" {
				m.currentState = "databases"

				dbs := []string{"PostgreSQL", "MySQL", "MariaDB", "MongoDB", "Oracle", "SQL Server"}
				items := make([]list.Item, len(dbs))
				for i, db := range dbs {
					items[i] = item{title: db, desc: "Descrição do " + db}
				}

				delegate := list.NewDefaultDelegate()
				m.list = list.New(items, delegate, m.width, m.height)
				m.list.Title = "Select a Database"
				m.list.SetShowStatusBar(false)
				m.list.SetFilteringEnabled(false)
				m.list.Styles.Title = lipgloss.NewStyle().MarginLeft(2).Bold(true)

				return m, nil
			}
		}
	case versionSelectionMsg:
		m.currentState = "versions"

		delegate := list.NewDefaultDelegate()
		m.list = list.New(msg.versions, delegate, m.width, m.height)
		m.list.Title = "Select a version"
		m.list.SetShowStatusBar(false)
		m.list.SetFilteringEnabled(false)
		m.list.Styles.Title = lipgloss.NewStyle().MarginLeft(2).Bold(true)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func loadVersions(m model) tea.Cmd {
	return func() tea.Msg {
		dir := filepath.Join("docker-configs", m.selectedItem)
		files, err := os.ReadDir(dir)
		if err != nil {
			fmt.Printf("Error while reading directory %s: %v\n", dir, err)
			return versionSelectionMsg{
				versions: []list.Item{
					item{
						title: "Erro: directory not found",
						desc:  fmt.Sprintf("Cound'nt find %s", dir),
					},
				},
			}
		}

		var items []list.Item
		for _, file := range files {
			if file.IsDir() {
				items = append(items, item{title: file.Name(), desc: "Version " + file.Name()})
			}
		}

		if len(items) == 0 {
			items = append(items, item{title: "Any available version", desc: ""})
		}

		return versionSelectionMsg{versions: items}
	}
}

func startDockerCompose(version string) tea.Cmd {
	return func() tea.Msg {
		fmt.Println("Rodando Docker Compose para", version)
		cmd := exec.Command("echo", "Rodando Docker Compose para", version)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		_ = cmd.Run()
		return tea.Quit
	}
}

func main() {
	m := initialModel()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println("Erro ao iniciar a TUI:", err)
		os.Exit(1)
	}
}