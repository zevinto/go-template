/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	//_ "go-template/app/event"
	"go-template/app/route"
	"go-template/internal/server"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start serve",
	Long:  `start serve`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func run() {
	http := server.New()
	http.GetRouter(route.New())
	http.Run()
}
