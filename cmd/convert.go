/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"path"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	to_anf bool
	to_cnf bool
	to_dnf bool
	to_nnf bool
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert the expression to various forms",
	Long: `
converts to different forms by using flags -a for anf, -c for cnf, -d for dnf, -n for nnf `,
	Run: func(cmd *cobra.Command, args []string) {

		dir_path := SCRIPT_ADDRESS

		if len(args) == 0 {
			fmt.Println(color.RedString("\n ERROR: give the path of file first \n"))
			return
		}

		fmt.Println()
		if to_anf {
			command := exec.Command("python", path.Join(dir_path, "convert.py"), "1", args[0])
			dt, err := command.CombinedOutput()
			if err != nil {
				fmt.Println(color.RedString("ERROR: bad expressions \n"))
				return
			}
			fmt.Print(color.GreenString(string(dt)))
		}
		if to_cnf {
			command := exec.Command("python", path.Join(dir_path, "convert.py"), "2", args[0])
			dt, err := command.CombinedOutput()
			if err != nil {
				fmt.Println(color.RedString("ERROR: bad expressions \n"))
				return
			}
			fmt.Print(color.GreenString(string(dt)))
		}
		if to_dnf {
			command := exec.Command("python", path.Join(dir_path, "convert.py"), "3", args[0])
			dt, err := command.CombinedOutput()
			if err != nil {
				fmt.Println(color.RedString("ERROR: bad expressions \n"))
				return
			}
			fmt.Print(color.GreenString(string(dt)))
		}
		if to_nnf {
			command := exec.Command("python", path.Join(dir_path, "convert.py"), "4", args[0])
			dt, err := command.CombinedOutput()
			if err != nil {
				fmt.Println(color.RedString("ERROR: bad expressions \n"))
				return
			}
			fmt.Print(color.GreenString(string(dt)))
		}

		fmt.Println()

	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	convertCmd.Flags().BoolVarP(&to_anf, "anf", "a", false, "convert to anf")
	convertCmd.Flags().BoolVarP(&to_cnf, "cnf", "c", false, "convert to cnf")
	convertCmd.Flags().BoolVarP(&to_dnf, "dnf", "d", false, "convert to dnf")
	convertCmd.Flags().BoolVarP(&to_nnf, "nnf", "n", false, "convert to nnf")
}
