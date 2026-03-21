package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Запустить управление с помощью жестов",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Утилита запущена...")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
