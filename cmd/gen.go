/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"drip/cli/flags"
	"drip/utils"
	"fmt"
	"os"
	"strings"
	"drip/utils/stringutils"
	"drip/cli/colors"

	"github.com/spf13/cobra"
	"path/filepath"
)

var (
	cssFiles  bool
	fonts     bool
	separator string
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a CSS file for fonts",
	Long:  "This command generates a CSS file defining font-face rules based on font files found in the specified directory.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen called")
		if directory == "" {
			fmt.Println("Directory is required")
			return
		}

		if cssFiles && fonts {
			genCssFontsFile(directory, separator)
		}
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	flags.AddDirectoryFlag(genCmd, &directory)

	genCmd.Flags().BoolVarP(&cssFiles, "cssFiles", "c", false, "Create CSS file")
	genCmd.Flags().BoolVarP(&fonts, "fonts", "f", false, "Declare font in file")
	genCmd.Flags().StringVarP(&separator, "separator", "s", "", "Separator used in font file names")
}

func genCssFontsFile(directory, separator string) {
	files, err := utils.GetFilesInDirectory(directory)
	if err != nil {
		fmt.Println("Error retrieving files:", err)
		return
	}
	cssPath := filepath.Join(directory, "fonts.css")
	cssFile, err := os.Create(cssPath)
	if err != nil {
		fmt.Println(colors.ErrorColor("Error creating file:"), err)
		return
	}
	defer cssFile.Close()

	for _, file := range files {

		ext := filepath.Ext(file)
		fileSuffix := strings.TrimSuffix(filepath.Base(file), ext)
		fileInfo := strings.Split(fileSuffix, separator)

		if len(fileInfo) < 2 {
			fmt.Println("Invalid file name format:", fileSuffix)
			continue
		}

		fileName := fileInfo[0]
		fontType := stringutils.FindWeight(fileInfo[1])

		fileDirectory := filepath.Base(file)

		var format string

		if ext == ".ttf" {
			format = "truetype"
			_, err = fmt.Fprintf(cssFile, "@font-face {\n  font-family: '%s';\n  src: url('%s') format('%s');\n  font-weight: %s;\n  font-style: %s;\n}\n\n", fileName, fileDirectory, format, fontType[0], fontType[1])
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}
	}

	fmt.Println(colors.SuccessColor("CSS file generated successfully")) 

	
}


