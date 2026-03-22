package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"cli/internal/ffmpeg"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Запустить управление с помощью жестов",
	RunE: func(cmd *cobra.Command, args []string) error {
		found, err := ffmpeg.Find()
		if err != nil {
			return err
		}

		fmt.Printf("FFmpeg найден: %s\n", found.Path)
		fmt.Println("Утилита запущена...")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
