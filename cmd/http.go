/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/gabrielteiga/golang-cli-cobra/web"
	"github.com/spf13/cobra"
)

var port string

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Run HTTP Server",
	Long: `This command will start a HTTP Server on the port informed. 
	The default port is 4040 and you can change it using the flag --port or -p. 

	Example of usage: 
		go run main.go http -p=:4040
		
	After that, you can access the server on localhost:4040/?a=x&b=y (change x and y for numbers that you want to sum)`,
	Run: func(cmd *cobra.Command, args []string) {
		server := web.Server{Port: port}
		server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
	httpCmd.Flags().StringVarP(&port, "port", "p", ":4040", "Port to be used on HTTP Server")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
