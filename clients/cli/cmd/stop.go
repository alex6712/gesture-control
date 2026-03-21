package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Остановить управление с помощью жестов",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Утилита остановлена")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
