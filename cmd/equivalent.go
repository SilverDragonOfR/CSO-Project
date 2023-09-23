/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"path"
	"strings"

	"github.com/fatih/color"

	"github.com/spf13/cobra"
)

// equivalentCmd represents the equivalent command
var equivalentCmd = &cobra.Command{
	Use:   "equivalent",
	Short: "check whether statements are equivalent",
	Long: `
It outputs true or False depending on whether all the statements are equivalent or not`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 1 {
			fmt.Println(color.RedString("\n ERROR: Atleast 2 staements need to be given"))
			return
		}
		fmt.Println()

		dir_path := SCRIPT_ADDRESS

		command := exec.Command("python", path.Join(dir_path, "equivalent.py"), strings.Join(args, " "))
		dt, err := command.CombinedOutput()
		if err != nil {
			fmt.Println(color.RedString("ERROR: bad expressions \n"))
			return
		}
		fmt.Print(color.GreenString(string(dt)))

		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(equivalentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// equivalentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// equivalentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
