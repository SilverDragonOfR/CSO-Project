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
	is_anf bool
	is_cnf bool
	is_dnf bool
	is_nnf bool
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check whether the expression is in various forms",
	Long: `
checks the different forms by using flags -a for anf, -c for cnf, -d for dnf, -n for nnf `,
	Run: func(cmd *cobra.Command, args []string) {

		dir_path := SCRIPT_ADDRESS

		if len(args) == 0 {
			fmt.Println(color.RedString("\n ERROR: give the path of file first \n"))
			return
		}

		fmt.Println()
		if is_anf {
			command := exec.Command("python", path.Join(dir_path, "check.py"), "1", args[0])
			dt, err := command.CombinedOutput()
			if err != nil {
				fmt.Println(color.RedString("ERROR: bad expressions \n"))
				return
			}
			fmt.Print(color.GreenString(string(dt)))
		}
		if is_cnf {
			command := exec.Command("python", path.Join(dir_path, "check.py"), "2", args[0])
			dt, err := command.CombinedOutput()
			if err != nil {
				fmt.Println(color.RedString("ERROR: bad expressions \n"))
				return
			}
			fmt.Print(color.GreenString(string(dt)))
		}
		if is_dnf {
			command := exec.Command("python", path.Join(dir_path, "check.py"), "3", args[0])
			dt, err := command.CombinedOutput()
			if err != nil {
				fmt.Println(color.RedString("ERROR: bad expressions \n"))
				return
			}
			fmt.Print(color.GreenString(string(dt)))
		}
		if is_nnf {
			command := exec.Command("python", path.Join(dir_path, "check.py"), "4", args[0])
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
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	checkCmd.Flags().BoolVarP(&is_cnf, "cnf", "c", false, "convert to cnf")
	checkCmd.Flags().BoolVarP(&is_anf, "anf", "a", false, "convert to anf")
	checkCmd.Flags().BoolVarP(&is_dnf, "dnf", "d", false, "convert to dnf")
	checkCmd.Flags().BoolVarP(&is_nnf, "nnf", "n", false, "convert to nnf")
}
