package cmd

import (
	"fmt"

	"github.com/ank809/Todolist-CLI-Golang/database"
	"github.com/spf13/cobra"
)

var CreateTodo = &cobra.Command{
	Use:   "create",
	Short: "Creating",
	Long:  "This command is used to create todo",
	Run:   CreateTodos,
}

var (
	title       string
	description string
)

func init() {
	CreateTodo.Flags().StringVarP(&title, "title", "t", "", "Title of the Todo ")
	CreateTodo.Flags().StringVarP(&description, "description", "d", "", "Description")
	CreateTodo.MarkFlagRequired("title")
}

func CreateTodos(cmd *cobra.Command, args []string) {
	res, err := database.AddTodo(title, description)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(res)

}
