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
    Short: "Renommer tous les fichiers dans un répertoire avec un nom générique et suffixe personnalisé",
    Run: func(cmd *cobra.Command, args []string) {
        if directory == "" || genericName == "" {
            fmt.Println("Le répertoire et le nom générique sont requis.")
            return
        }

        files, err := getFilesInDirectory(directory)
        if err != nil {
            fmt.Println("Erreur lors de la récupération des fichiers :", err)
            return
        }

        for i, file := range files {
            newName := fmt.Sprintf("%s_%d", genericName, i+1)
            fmt.Printf("Fichier: %s -> Nouveau nom proposé: %s\n", filepath.Base(file), newName)

            reader := bufio.NewReader(os.Stdin)
            fmt.Print("Entrez un suffixe personnalisé ou appuyez sur Entrée pour valider: ")
            suffix, _ := reader.ReadString('\n')
            suffix = strings.TrimSpace(suffix)

            if suffix != "" {
                newName += "_" + suffix
            }

            newPath := filepath.Join(filepath.Dir(file), newName+filepath.Ext(file))
            err := os.Rename(file, newPath)
            if err != nil {
                fmt.Println("Erreur lors du renommage du fichier :", err)
            } else {
                fmt.Printf("Fichier renommé en : %s\n", newName+filepath.Ext(file))
            }
        }
    },
}

func init() {
    rootCmd.AddCommand(renameCmd)


    renameCmd.Flags().StringVarP(&directory, "directory", "d", "", "Répertoire contenant les fichiers")
    renameCmd.Flags().StringVarP(&genericName, "name", "n", "", "Nom générique pour les fichiers")
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
