package command

import "fmt"

// CommandInvoker — исполнитель команд
type CommandInvoker struct {
	history *CommandHistory
}

// NewCommandInvoker — конструктор
func NewCommandInvoker() *CommandInvoker {
	return &CommandInvoker{
		history: NewCommandHistory(),
	}
}

// Execute — выполняет команду и сохраняет в историю
func (i *CommandInvoker) Execute(cmd Command) error {
	fmt.Printf("\nВыполнение: %s\n", cmd.GetName())

	err := cmd.Execute()
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return err
	}

	i.history.Add(cmd)
	fmt.Printf("Команда сохранена в истории\n")
	return nil
}

// Undo — отменяет последнюю команду
func (i *CommandInvoker) Undo() error {
	fmt.Println("\nОТМЕНА ПОСЛЕДНЕЙ КОМАНДЫ")
	return i.history.Undo()
}

// Redo — повторяет последнюю отмененную команду
func (i *CommandInvoker) Redo() error {
	fmt.Println("\nПОВТОР ПОСЛЕДНЕЙ ОТМЕНЕННОЙ КОМАНДЫ")
	return i.history.Redo()
}

// ShowHistory — показывает историю
func (i *CommandInvoker) ShowHistory() {
	i.history.Show()
}
