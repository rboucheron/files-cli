/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"path/filepath"
)

var (
	cssFiles   bool
	fonts      bool
	separator  string

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
	addDirectoryFlag(genCmd, &directory)
	genCmd.Flags().BoolVarP(&cssFiles, "cssFiles", "c", false, "Create CSS file")
	genCmd.Flags().BoolVarP(&fonts, "fonts", "f", false, "Declare font in file")
	genCmd.Flags().StringVarP(&separator, "separator", "s", "", "Separator used in font file names")
}

func genCssFontsFile(directory, separator string) {
	files, err := getFilesInDirectory(directory)
	if err != nil {
		fmt.Println("Error retrieving files:", err)
		return
	}
	cssPath := filepath.Join(directory, "fonts.css")
	cssFile, err := os.Create(cssPath)
	if err != nil {
		fmt.Println("Error creating file:", err)
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
		fontType := findWeight(fileInfo[1])

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

	fmt.Println("CSS file generated successfully")
}

func findWeight(fontType string) [2]string {
	switch fontType {
	case "Thin":
		return [2]string{"100", "normal"}
	case "ExtraLight":
		return [2]string{"200", "normal"}
	case "Light":
		return [2]string{"300", "normal"}
	case "Normal":
		return [2]string{"400", "normal"}
	case "Medium":
		return [2]string{"500", "normal"}
	case "SemiBold":
		return [2]string{"600", "normal"}
	case "Bold":
		return [2]string{"700", "normal"}
	case "ExtraBold":
		return [2]string{"800", "normal"}
	case "Black":
		return [2]string{"900", "normal"}
	case "ThinItalic":
		return [2]string{"100", "italic"}
	case "ExtraLightItalic":
		return [2]string{"200", "italic"}
	case "LightItalic":
		return [2]string{"300", "italic"}
	case "NormalItalic":
		return [2]string{"400", "italic"}
	case "MediumItalic":
		return [2]string{"500", "italic"}
	case "SemiBoldItalic":
		return [2]string{"600", "italic"}
	case "BoldItalic":
		return [2]string{"700", "italic"}
	case "ExtraBoldItalic":
		return [2]string{"800", "italic"}
	case "BlackItalic":
		return [2]string{"900", "italic"}
	default:
		return [2]string{"unknown", "unknown"}
	}
}
