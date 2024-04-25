package main

import (
	"fmt"
	"os"

	"github.com/ank809/Todolist-CLI-Golang/cmd"
)

func main() {

	cmd.RootCmd.AddCommand(cmd.CreateTodo)
	cmd.RootCmd.AddCommand(cmd.ReadTodo)
	cmd.RootCmd.AddCommand(cmd.UpdateTodo)
	cmd.RootCmd.AddCommand(cmd.DeleteTodo)
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
}
