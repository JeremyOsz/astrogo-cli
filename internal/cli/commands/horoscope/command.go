package horoscope

import (
	"astrogo-cli/internal/service"
	"fmt"

	"github.com/spf13/cobra"
)

// NewStarSignCmd creates the star sign command
func NewStarSignCmd() *cobra.Command {
	var sign string
	var date string

	starSignCmd := &cobra.Command{
		Use:   "starsign",
		Short: "Get your horoscope based on star sign",
		Run: func(cmd *cobra.Command, args []string) {
			if sign == "" {
				fmt.Println("Error: Star sign is required")
				sign = SelectSign()
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

	return starSignCmd
}
