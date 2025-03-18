package horoscope

import (
	"astrogo-cli/internal/cli/ui/result"

	tea "github.com/charmbracelet/bubbletea"
)

func SelectSign() string {
	// TODO: get options from signs

	choices := []string{
		// Star signs in order
		"Aries",
		"Taurus",
		"Gemini",
		"Cancer",
		"Leo",
		"Virgo",
		"Libra",
		"Scorpio",
		"Sagittarius",
		"Capricorn",
		"Aquarius",
		"Pisces",
	}

	// List possible star signs
	selectProgram := tea.NewProgram(result.InitialResultModel(choices))
	m, err := selectProgram.Run()
	if err != nil {
		panic(err)
	}

	// Assert the type of m to your custom Model type
	if model, ok := m.(result.Model); ok {
		// Set the choice from the asserted model
		choice := model.Choice // Access the choice from the asserted model

		// Return the selected sign
		return choice
	}

	// Handle the case where the assertion fails
	return "" // or handle the error as needed
}
