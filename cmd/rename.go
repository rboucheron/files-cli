package cmd

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
    "strings"

    "github.com/spf13/cobra"
)

var directory string
var genericName string


var renameCmd = &cobra.Command{
    Use:   "rename",
    Short: "Rename all files in a directory with a generic name and custom suffix",
    Run: func(cmd *cobra.Command, args []string) {
        if directory == "" || genericName == "" {
            fmt.Println("Directory and generic name are required")
            return
        }

        files, err := getFilesInDirectory(directory)
        if err != nil {
            fmt.Println("Error retrieving files :", err)
            return
        }

        for i, file := range files {
            newName := fmt.Sprintf("%s_%d", genericName, i+1)
            fmt.Printf("files: %s -> new name : %s\n", filepath.Base(file), newName)

            reader := bufio.NewReader(os.Stdin)
            fmt.Print("Enter a custom suffix or press Enter to validate.: ")
            suffix, _ := reader.ReadString('\n')
            suffix = strings.TrimSpace(suffix)

            if suffix != "" {
                newName += "_" + suffix
            }

            newPath := filepath.Join(filepath.Dir(file), newName+filepath.Ext(file))
            err := os.Rename(file, newPath)
            if err != nil {
                fmt.Println("Error renaming file :", err)
            } else {
                fmt.Printf("File renamed to : %s\n", newName+filepath.Ext(file))
            }
        }
    },
}

func init() {
    rootCmd.AddCommand(renameCmd)


    renameCmd.Flags().StringVarP(&directory, "directory", "d", "", "Directory containing the files")
    renameCmd.Flags().StringVarP(&genericName, "name", "n", "", "Generic name for files")
}


func getFilesInDirectory(directory string) ([]string, error) {
    var files []string
    err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            files = append(files, path)
        }
        return nil
    })
    if err != nil {
        return nil, err
    }
    return files, nil
}
