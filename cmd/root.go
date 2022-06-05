/*
	Copyright © 2022 Purvesh Patel <ind.purvesh@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/joho/godotenv"
)

var rootCmd = &cobra.Command{
	Use:   "avored",
	Short: "AvoRed Go is an e commerce for Go Language",
	Long: `AvoRed Go is an e commerce for Go Language`,
	Run: handle,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func handle (cmd *cobra.Command, args []string) {
	godotenv.Load()
}
