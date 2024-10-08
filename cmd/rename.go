package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"drip/cli/flags"
	"drip/utils"

	"github.com/spf13/cobra"
)


var (
	directory   string
	genericName string
	space       bool
)


var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename all files in a directory with a generic name and custom suffix",
	Run: func(cmd *cobra.Command, args []string) {

		if directory == "" {
			fmt.Println("Directory is required")
			return
		}

		if !space && genericName == "" {
			fmt.Println("Generic name is required when --space is not used")
			return
		}

		if space {
			rmspace(directory)
		} else {
			genericrename(directory, genericName)
		}
	},
}


func init() {

	rootCmd.AddCommand(renameCmd)
	flags.AddDirectoryFlag(renameCmd, &directory)


	renameCmd.Flags().StringVarP(&genericName, "generic", "g", "", "Generic name for renaming files")
	renameCmd.Flags().BoolVarP(&space, "space", "s", false, "Remove spaces from filenames")

}



func rmspace(directory string) {
	files, err := utils.GetFilesInDirectory(directory)
	if err != nil {
		fmt.Println("Error retrieving files:", err)
		return
	}

	for _, file := range files {
		newFilename := strings.ReplaceAll(filepath.Base(file), " ", "_")
		newPath := filepath.Join(filepath.Dir(file), newFilename)
		err := os.Rename(file, newPath)
		if err != nil {
			fmt.Println("Error renaming file:", err)
		} else {
			fmt.Printf("File renamed to: %s\n", newPath)
		}
	}
}


func genericrename(directory, genericName string) {
	files, err := utils.GetFilesInDirectory(directory)
	if err != nil {
		fmt.Println("Error retrieving files:", err)
		return
	}

	for _, file := range files {
	
		baseName := strings.TrimSuffix(genericName, filepath.Ext(file))
		newName := fmt.Sprintf("%s", baseName) 
		fmt.Printf("files: %s -> new name: %s\n", filepath.Base(file), newName)
	
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a custom suffix or press Enter to validate: ")
		suffix, _ := reader.ReadString('\n')
		suffix = strings.TrimSpace(suffix)
	
	
		if suffix != "" {
			newName += "_" + suffix
		}
	
		newPath := filepath.Join(filepath.Dir(file), newName+filepath.Ext(file))
		err := os.Rename(file, newPath)
		if err != nil {
			fmt.Println("Error renaming file:", err)
		} else {
			fmt.Printf("File renamed to: %s\n", newName+filepath.Ext(file))
		}
	}
	
}


