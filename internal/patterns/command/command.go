package command

import "fmt"

// Command — интерфейс для всех команд
type Command interface {
	Execute() error
	Undo() error
	GetName() string
}

// CommandHistory — история выполненных команд (для Undo/Redo)
type CommandHistory struct {
	commands []Command
	position int
}

// NewCommandHistory — конструктор истории
func NewCommandHistory() *CommandHistory {
	return &CommandHistory{
		commands: []Command{},
		position: -1,
	}
}

// Add — добавляет команду в историю
func (h *CommandHistory) Add(cmd Command) {
	// Удаляем всё после текущей позиции
	h.commands = h.commands[:h.position+1]
	h.commands = append(h.commands, cmd)
	h.position++
}

// Undo — отменяет последнюю команду
func (h *CommandHistory) Undo() error {
	if h.position < 0 {
		return fmt.Errorf("нечего отменять")
	}
	cmd := h.commands[h.position]
	err := cmd.Undo()
	if err == nil {
		h.position--
	}
	return err
}

// Redo — повторяет отмененную команду
func (h *CommandHistory) Redo() error {
	if h.position+1 >= len(h.commands) {
		return fmt.Errorf("нечего повторять")
	}
	h.position++
	cmd := h.commands[h.position]
	return cmd.Execute()
}

// Show — показывает историю команд
func (h *CommandHistory) Show() {
	if len(h.commands) == 0 {
		fmt.Println("История пуста")
		return
	}
	fmt.Println("\nИСТОРИЯ КОМАНД:")
	for i, cmd := range h.commands {
		marker := ""
		if i == h.position {
			marker = " ← текущая"
		}
		fmt.Printf("  %d. %s%s\n", i+1, cmd.GetName(), marker)
	}
}
