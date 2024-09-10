package flags

import (
    "github.com/spf13/cobra"
)

func AddDirectoryFlag(cmd *cobra.Command, directory *string) {
    cmd.Flags().StringVarP(directory, "directory", "d", "", "Directory to use")
}
