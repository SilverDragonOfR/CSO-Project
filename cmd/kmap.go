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
	sop      bool
	pos      bool
	num      string
	minterms string
	dontcare string
)

// kmapCmd represents the kmap command
var kmapCmd = &cobra.Command{
	Use:   "kmap",
	Short: "give the SOP or POS form",
	Long: `
gives the SOP or POS form of minterms and dontcare by specifying
-m as "1,3,4,5" and -d "2,8,9"`,
	Run: func(cmd *cobra.Command, args []string) {
		dir_path := SCRIPT_ADDRESS

		if !sop && !pos {
			fmt.Println(color.RedString("\n ERROR: give atleast one of sop or pos flag \n"))
			return
		}
		fmt.Println()
		if sop {
			command := exec.Command("python", path.Join(dir_path, "kmap.py"), "1", num, minterms, dontcare)
			dt, err := command.CombinedOutput()
			if err != nil {
				fmt.Println(color.RedString("ERROR: bad expressions \n"))
				return
			}
			fmt.Print(color.GreenString(string(dt)))
		}
		if pos {
			command := exec.Command("python", path.Join(dir_path, "kmap.py"), "2", num, minterms, dontcare)
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
	rootCmd.AddCommand(kmapCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kmapCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	kmapCmd.Flags().BoolVarP(&sop, "sop", "s", false, "gives the SOP form")
	kmapCmd.Flags().BoolVarP(&pos, "pos", "p", false, "gives the POS form")
	kmapCmd.Flags().StringVarP(&num, "num", "n", "4", "number of variables (< 27)")
	kmapCmd.Flags().StringVarP(&minterms, "minterms", "m", "1", "comma separated minterms")
	kmapCmd.Flags().StringVarP(&dontcare, "dontcare", "d", "2", "comma separated dontcare")
}
