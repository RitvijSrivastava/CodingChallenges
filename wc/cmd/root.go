/*
Copyright © 2023 Ritvij Srivastava ritvijsrivastava99@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wc",
	Short: "A mini copy of the unix wc command line tool",
	Long: `wc counts lines, words, runes, and bytes 
	in the named files, or in the standard input if no file is named. A word is a maximal 
	string of characters delimited by spaces, tabs or newlines.`,
	Run: getCount,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("bytes", "c", false, "print the byte counts")
	rootCmd.Flags().BoolP("lines", "l", false, "print the newline counts")
	rootCmd.Flags().BoolP("words", "w", false, "print the word counts")
	rootCmd.Flags().BoolP("chars", "m", false, "print the character counts")
}

