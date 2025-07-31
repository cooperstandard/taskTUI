package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

type Task struct {
	id          uuid.UUID
	title       string
	description string
	project     uuid.UUID
	status      string
	createdAt   time.Time
	completedAt time.Time
}

type Project struct {
	id    uuid.UUID
	title string
}

type model struct {
	tasks    []Task
	projects []Project
	cursor   int // which task index is currently selected in the current view. When the view changes: reset to 0
	// TODO: should be some type of linked list or id for quick lookup
}

func initialModel() model {
	project := Project{
		id:    uuid.New(),
		title: "test project",
	}
	task := Task{
		id:          uuid.New(),
		title:       "test task",
		description: "this is a sample task to test the layout",
		project:     project.id,
		createdAt:   time.Now(),
	}
	m := model{
		tasks:    []Task{task, task},
		projects: []Project{project},
		cursor:   0,
	}

	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			ok := m.cursor < len(m.tasks)
			if ok {

				selectedTask := &m.tasks[m.cursor]
				selectedTask.title = "selected " + selectedTask.title

			} else {
				m.tasks = append(m.tasks, Task{
					id:          uuid.New(),
					title:       "this is a new task",
					description: "this is a sample task to test the layout",
					createdAt:   time.Now(),
				})
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "TaskTUI \n\n"

	// Iterate over our choices
	for i, choice := range m.tasks {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, choice.title)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
