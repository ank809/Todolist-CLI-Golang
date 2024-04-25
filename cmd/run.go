package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var short string = "This is a simple todolist cli application"
var long = "You can :-  \n create \n read \n update \n delete todos using the commands"

var RootCmd = &cobra.Command{
	Use:   "todolist",
	Short: short,
	Long:  long,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(short)
			fmt.Println(long)
			fmt.Println("For more help type --help")
			return
		}
	},
}
