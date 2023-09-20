/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var givenPath string

// truthtableCmd represents the truthtable command

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

var truthtableCmd = &cobra.Command{
	Use:   "truthtable",
	Short: "generate the truthtable for any expression",
	Long: color.CyanString(`
It generates the truthtable for Boolean expressions within double quotes.`),
	Run: func(cmd *cobra.Command, args []string) {

		// if len(args) == 0 {
		// 	fmt.Println(color.RedString("\n ERROR: Atleast one argument is required \n"))
		// 	return
		// }

		// testpath := `C:\Users\rachi\OneDrive\Desktop\MyFolder\cso\scripts\test.py`
		// os.Unsetenv("TRUTHPATH")
		// print("hi:", givenPath)
		if len(givenPath) > 0 {
			if isExist, _ := Exists(givenPath); !isExist {
				fmt.Println(color.RedString("\n ERROR: no file found \n"))
				return
			}
			os.Setenv("TRUTHPATH", givenPath)
		} else {
			os.Setenv("TRUTHPATH", `C:\Users\rachi\OneDrive\Desktop\MyFolder\cso\scripts\test.py`)
		}

		path, ok := os.LookupEnv("TRUTHPATH")
		if !ok {
			fmt.Println(color.RedString("\n ERROR: give the path of file first \n"))
			return
		}

		fmt.Println()
		for _, exp := range args {

			// fmt.Println(path)

			cmd := exec.Command("python", path, exp)
			dt, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println(color.RedString("ERROR: \"%v\" is not a proper expression\n", exp))
				continue
			}
			fmt.Print(color.GreenString(string(dt)))

		}
		fmt.Println()

	},
}

func init() {
	rootCmd.AddCommand(truthtableCmd)
	truthtableCmd.PersistentFlags().StringVarP(&givenPath, "path", "p", "", "complete path to file")
}
