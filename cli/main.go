package main

import (
	"astrogo-cli/internal/cli/commands/database"
	"astrogo-cli/internal/cli/commands/horoscope"
	"astrogo-cli/internal/cli/commands/user"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "astro",
		Short: "Astrology CLI application",
		Long:  `A command-line astrology application that provides horoscopes and birth chart analysis.`,
	}

	// Add the commands to the root command
	rootCmd.AddCommand(
		horoscope.NewStarSignCmd(),
		user.NewUserCmd(),
		database.NewDatabaseCmd(),
	)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
