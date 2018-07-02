package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init(){
	RootCmd.AddCommand(VersionCmd)
}

var VersionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("ljf is version 1.0")
	},
}