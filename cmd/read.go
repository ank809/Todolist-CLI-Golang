package cmd

import (
	"fmt"

	"github.com/ank809/Todolist-CLI-Golang/database"
	"github.com/spf13/cobra"
)

var ReadTodo = &cobra.Command{
	Use:   "read",
	Short: "Reading",
	Long:  "This command is used to read all todos",
	Run:   ReadTodos,
}

func ReadTodos(cmd *cobra.Command, args []string) {
	res, err := database.GetAllTodos()
	if err != nil {
		fmt.Println(err)
		return

	}
	for _, todo := range res {
		fmt.Println(todo)
	}
}
