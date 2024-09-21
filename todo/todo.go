package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
)

const (
	todoFile = ".todo.json"
)

type todoItem struct {
	Task        string
	Done        bool
	CreatedAt   string
	CompletedAt string
}

type Todos []todoItem

func (t *Todos) Add(task string) error {
	curTime := time.Now()
	formattedTime := curTime.Format("15:04")

	todo := todoItem{
		Task:        task,
		Done:        false,
		CreatedAt:   formattedTime,
		CompletedAt: "-",
	}

	*t = append(*t, todo)
	return nil
}

func (t *Todos) Done(index int) error {

	ls := *t
	if index < 0 || index >= len(ls) {
		return errors.New("Invalid input")
	}

	curTime := time.Now()
	formattedTime := curTime.Format("15:04")
	ls[index].CompletedAt = formattedTime
	ls[index].Done = true

	return nil

}

func (t *Todos) Delete(index int) error {
	ls := *t
	if index < 0 || index >= len(ls) {
		return errors.New("Invalid input")
	}
	*t = append(ls[:index-1], ls[index:]...)
	return nil

}

func (t *Todos) Load(fileName string) error {
	file, err := os.ReadFile(todoFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todos) Store(fileName string) error {

	data, err := json.Marshal(t)

	if err != nil {
		return err
	}
	return os.WriteFile(todoFile, data, 0644)

}
func (t *Todos) Print() {

	fmt.Println("")

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Status"},
		},
	}

	// subtotal := float64(0)
	for index, todo := range *t {
		var tick string
		if todo.CompletedAt == "-" {
			tick = fmt.Sprintf("\033[31m%s\033[0m", "\u2718")
		} else {
			tick = fmt.Sprintf("\033[32m%s\033[0m", "\u2714")
		}
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", index+1)},
			{Text: todo.Task},
			{Align: simpletable.AlignCenter, Text: tick},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleDefault)
	fmt.Println(table.String())

}
