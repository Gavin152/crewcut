package cmd

import (
	"database/sql"
	"fmt"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"os"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open an existing Crew",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := getCrews()
		if err != nil {
			fmt.Printf("Error getting crews: %v\n", err)
			os.Exit(1)
		}

		var crews []crew
		for result.Next() {
			var id int
			var name string
			_ = result.Scan(&id, &name)
			crews = append(crews, crew{id, name})
		}

		var vpHeight int
		if len(crews) > 10 {
			vpHeight = len(crews) + 4
		} else {
			vpHeight = 14
		}

		vp := viewport.New(60, vpHeight)
		vp.Style = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#008896")).
			PaddingLeft(2)

		app := tea.NewProgram(model{
			cursor:   0,
			crews:    crews,
			viewport: vp,
		})
		_, err = app.Run()
		if err != nil {
			fmt.Println("Whoops, something went wrong here:\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}

func getCrews() (*sql.Rows, error) {
	db, err := sql.Open("sqlite", "data.db")
	defer db.Close()

	stmt, _ := db.Prepare("SELECT id, name FROM crews")
	res, err := stmt.Query()

	return res, err
}

// ==========================================
//				BUBBLETEA SETUP
// ==========================================

type crew struct {
	id   int
	name string
}

type model struct {
	cursor   int
	crews    []crew
	viewport viewport.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.crews)-1 {
				m.cursor++
			}
		case "enter", "l":

		}
	}
	return m, nil
}

func (m model) View() string {
	tuiString := "Which crew do you want to open?\n\n"

	for i, _ := range m.crews {
		if i == m.cursor {
			tuiString += "> "
		} else {
			tuiString += "  "
		}

		tuiString += fmt.Sprintf("%d ", m.crews[i].id)
		tuiString += m.crews[i].name
		tuiString += "\n"
	}

	for range m.viewport.Height - len(m.crews) - 4 {
		tuiString += "\n"
	}

	tuiString += fmt.Sprintf("\nPress 'q' to quit or 'Enter' to select")
	m.viewport.SetContent(tuiString)
	return m.viewport.View()
}
