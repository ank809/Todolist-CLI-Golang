package cmd

import (
	"fmt"

	"github.com/ank809/Todolist-CLI-Golang/database"
	"github.com/spf13/cobra"
)

var DeleteTodo = &cobra.Command{
	Use:   "delete",
	Short: "Deleting",
	Long:  "This command is used to delete a todo",
	Run:   Deletetodo,
}
var todoid string

func Deletetodo(cmd *cobra.Command, args []string) {
	res, err := database.Deletetodos(todoid)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func init() {
	DeleteTodo.Flags().StringVarP(&todoid, "ID", "i", "", "ID of todo yoou want to delete")
	DeleteTodo.MarkFlagRequired("i")
}
