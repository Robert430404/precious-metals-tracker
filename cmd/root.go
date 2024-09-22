package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "precious-metals-tracker",
	Short: "Precious Metals Tracker is a simple cli tool that lets you track your holdings",
	Long:  `Precious Metals Tracker is a simple cli application for tracking precious metal holdings.

It's intended to give you the ability to track common things like current value by weight, total weight in your holdings, and the price you paid in comparison to the spot price at the time of purchase.`,
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


