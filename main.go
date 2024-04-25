package main

import (
	"fmt"
	"os"

	"github.com/ank809/Todolist-CLI-Golang/cmd"
)

func main() {

	cmd.RootCmd.AddCommand(cmd.CreateTodo)
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
}
