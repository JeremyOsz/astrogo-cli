package models

import "time"

// ZodiacSign represents a zodiac sign with its properties
type ZodiacSign struct {
	Name        string
	Symbol      string
	Element     string
	Quality     string
	Ruler       string
	DateStart   time.Time
	DateEnd     time.Time
	Description string
}

// Horoscope represents a horoscope reading
type Horoscope struct {
	Sign      string
	Date      time.Time
	Reading   string
	CreatedAt time.Time
}

// GetZodiacSigns returns data about all zodiac signs
func GetZodiacSigns() map[string]ZodiacSign {
	// Note: These dates are approximate and should be refined in a production app
	return map[string]ZodiacSign{
		"aries": {
			Name:        "Aries",
			Symbol:      "Ram",
			Element:     "Fire",
			Quality:     "Cardinal",
			Ruler:       "Mars",
			Description: "Aries is the first sign of the zodiac, symbolizing new beginnings, leadership, courage, and initiative.",
		},
		"taurus": {
			Name:        "Taurus",
			Symbol:      "Bull",
			Element:     "Earth",
			Quality:     "Fixed",
			Ruler:       "Venus",
			Description: "Taurus is known for stability, persistence, sensuality, and a connection to the material world.",
		},
		"gemini": {
			Name:        "Gemini",
			Symbol:      "Twins",
			Element:     "Air",
			Quality:     "Mutable",
			Ruler:       "Mercury",
			Description: "Gemini represents duality, communication, intellectual curiosity, and adaptability.",
		},
		"cancer": {
			Name:        "Cancer",
			Symbol:      "Crab",
			Element:     "Water",
			Quality:     "Cardinal",
			Ruler:       "Moon",
			Description: "Cancer embodies nurturing, emotional sensitivity, intuition, and a connection to home and family.",
		},
		"leo": {
			Name:        "Leo",
			Symbol:      "Lion",
			Element:     "Fire",
			Quality:     "Fixed",
			Ruler:       "Sun",
			Description: "Leo symbolizes creativity, self-expression, courage, leadership, and a generous spirit.",
		},
		"virgo": {
			Name:        "Virgo",
			Symbol:      "Virgin",
			Element:     "Earth",
			Quality:     "Mutable",
			Ruler:       "Mercury",
			Description: "Virgo represents analysis, practicality, service, attention to detail, and a focus on improvement.",
		},
		"libra": {
			Name:        "Libra",
			Symbol:      "Scales",
			Element:     "Air",
			Quality:     "Cardinal",
			Ruler:       "Venus",
			Description: "Libra embodies balance, harmony, justice, relationships, and aesthetic appreciation.",
		},
		"scorpio": {
			Name:        "Scorpio",
			Symbol:      "Scorpion",
			Element:     "Water",
			Quality:     "Fixed",
			Ruler:       "Pluto/Mars",
			Description: "Scorpio symbolizes intensity, transformation, depth, passion, and psychological insight.",
		},
		"sagittarius": {
			Name:        "Sagittarius",
			Symbol:      "Archer",
			Element:     "Fire",
			Quality:     "Mutable",
			Ruler:       "Jupiter",
			Description: "Sagittarius represents exploration, philosophy, optimism, freedom, and the quest for meaning.",
		},
		"capricorn": {
			Name:        "Capricorn",
			Symbol:      "Goat",
			Element:     "Earth",
			Quality:     "Cardinal",
			Ruler:       "Saturn",
			Description: "Capricorn embodies discipline, ambition, responsibility, structure, and practical achievement.",
		},
		"aquarius": {
			Name:        "Aquarius",
			Symbol:      "Water Bearer",
			Element:     "Air",
			Quality:     "Fixed",
			Ruler:       "Uranus/Saturn",
			Description: "Aquarius symbolizes innovation, originality, humanitarianism, community, and future vision.",
		},
		"pisces": {
			Name:        "Pisces",
			Symbol:      "Fish",
			Element:     "Water",
			Quality:     "Mutable",
			Ruler:       "Neptune/Jupiter",
			Description: "Pisces represents intuition, compassion, imagination, spirituality, and emotional depth.",
		},
	}
}
