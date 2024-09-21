package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/adityaRawat14/cli-todo/todo"
)

const (
	todoFile = ".todo.json"
)

func main() {

	add := flag.Bool("add", false, "add a new todo")
	done := flag.Int("done", 0, "complete todo")
	delete := flag.Int("delete", 0, "delete todo")
	list := flag.Bool("list", false, "list todo")
	flag.Parse()

	myTodos := &todo.Todos{}
	if err := myTodos.Load(todoFile); err != nil {

		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	switch {
	case *add:

		inputTask, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Println("Invalid commands")
			os.Exit(1)
		}

		myTodos.Add(inputTask)
		err = myTodos.Store(todoFile)
		if err != nil {
			os.Exit(1)
		}

	case *done > 0:
		err := myTodos.Done(*done-1)
		if err != nil {
			fmt.Println("done error",err)
			os.Exit(1)
		}
		err = myTodos.Store(todoFile)
		if err != nil {
			fmt.Println("error in updating ")
			os.Exit(1)
		}

	case *delete > 0:
		err := myTodos.Delete(*delete-1)
		if err != nil {
			fmt.Println("delete failed:",err)
			os.Exit(1)
		}
		err = myTodos.Store(todoFile)
		if err != nil {
			fmt.Println("error in updating ")
			os.Exit(1)
		}

	case *list:
		myTodos.Print()

	default:
		fmt.Println("Invalid Command")

	}

}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}
	text := scanner.Text()
	if len(text) == 0 {
		return "", errors.New("command is required !")
	}
	return text, nil

}
