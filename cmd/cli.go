/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/gabrielteiga/golang-cli-cobra/app"
	"github.com/spf13/cobra"
)

var num1, num2 float64

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Sum two numbers",
	Long:  `Long description`,
	Run: func(cmd *cobra.Command, args []string) {
		calc := app.NewCalc()
		calc.A = num1
		calc.B = num2

		fmt.Println(calc.Sum())
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)

	cliCmd.Flags().Float64VarP(&num1, "num1", "a", 0, "First number")
	cliCmd.Flags().Float64VarP(&num2, "num2", "b", 0, "Second number")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cliCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cliCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
