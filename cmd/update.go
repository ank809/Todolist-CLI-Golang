package cmd

import (
	"fmt"

	"github.com/ank809/Todolist-CLI-Golang/database"
	"github.com/spf13/cobra"
)

var UpdateTodo = &cobra.Command{
	Use:   "update",
	Short: "Updating",
	Long:  "This command is used to upadte the todo",
	Run:   UpdateTodos,
}

var Id string

// var (title string ; description string)

func UpdateTodos(cmd *cobra.Command, args []string) {

	res, err := database.UpdateTodo(Id, title, description)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(res)

}

func init() {
	UpdateTodo.Flags().StringVarP(&Id, "ID", "i", "", "Id is used to find and update the document")
	UpdateTodo.Flags().StringVarP(&title, "title", "t", "", "Title of the Todo ")
	UpdateTodo.Flags().StringVarP(&description, "description", "d", "", "Description")
	UpdateTodo.MarkFlagRequired("i")
	UpdateTodo.MarkFlagRequired("title")
}
