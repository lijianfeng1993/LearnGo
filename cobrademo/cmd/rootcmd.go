package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)


var RootCmd = &cobra.Command{
	Use: "ljf",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("I am ljf.")
	},

}
