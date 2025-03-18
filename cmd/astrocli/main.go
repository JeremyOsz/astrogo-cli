package main

import (
	"astrogo-cli/internal/cli/commands"
	"astrogo-cli/internal/service"
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

	// Star sign command
	var sign string
	var date string
	starSignCmd := &cobra.Command{
		Use:   "starsign",
		Short: "Get your horoscope based on star sign",
		Run: func(cmd *cobra.Command, args []string) {
			if sign == "" {
				fmt.Println("Error: Star sign is required")
				sign = commands.SelectSign()
			}

			// Format: if no date specified, use today
			displayDate := "today"
			if date != "" {
				displayDate = date
			}

			fmt.Printf("Sign: %v\n", sign)

			horoscope, err := service.GetStarSignHoroscope(sign, date)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}

			fmt.Printf("🌟 Horoscope for %s (%s):\n\n%s\n", sign, displayDate, horoscope)
		},
	}

	// Add flags to the star sign command
	starSignCmd.Flags().StringVarP(&sign, "sign", "s", "", "Your zodiac sign (e.g., aries, taurus, gemini)")
	starSignCmd.Flags().StringVarP(&date, "date", "d", "", "Date for horoscope (YYYY-MM-DD), defaults to today")

	// Add the command to the root command
	rootCmd.AddCommand(starSignCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
