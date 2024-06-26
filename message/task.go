package message

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/qualidafial/pomo"
)

func NewTask(status pomo.Status) tea.Cmd {
	return func() tea.Msg {
		return NewTaskMsg{
			Status: status,
		}
	}
}

type NewTaskMsg struct {
	Status pomo.Status
}

func EditTask(task pomo.Task) tea.Cmd {
	return func() tea.Msg {
		return EditTaskMsg{
			Task: task,
		}
	}
}

type EditTaskMsg struct {
	pomo.Task
}

func PromptDeleteTask(task pomo.Task) tea.Cmd {
	return func() tea.Msg {
		return PromptDeleteTaskMsg{
			Task: task,
		}
	}
}

type PromptDeleteTaskMsg struct {
	pomo.Task
}

func SaveTask(task pomo.Task) tea.Cmd {
	return func() tea.Msg {
		return SaveTaskMsg{
			Task: task,
		}
	}
}

type SaveTaskMsg struct {
	Task pomo.Task
}

func CancelEdit() tea.Msg {
	return CancelEditMsg{}
}

type CancelEditMsg struct{}
