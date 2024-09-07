/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"path/filepath"
	"strings"
    "os"
	"github.com/spf13/cobra"
)

// rmspaceCmd represents the rmspace command
var rmspaceCmd = &cobra.Command{
	Use:   "rmspace",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if directory == "" {
			fmt.Println("Directory and generic name are required")
			return
		}

		files, err := getFilesInDirectory(directory)
		if err != nil {
			fmt.Println("Error retrieving files :", err)
			return
		}

		for i, file := range files {
		
			newfilename := strings.ReplaceAll(filepath.Base(file), " ", "_")
			newpath := filepath.Join(filepath.Dir(file), newfilename)
			err := os.Rename(file, newpath)
            if err != nil {
                fmt.Println("Error renaming file :", err)
            } else {
                fmt.Printf("File renamed to : %s\n", newpath)
            }

			i++
		}
	},
}

func init() {
	rootCmd.AddCommand(rmspaceCmd)
	rmspaceCmd.Flags().StringVarP(&directory, "directory", "d", "", "Directory containing the files")

}
