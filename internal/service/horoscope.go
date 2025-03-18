package service

import (
	"astrogo-cli/internal/cli/models"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// GetStarSignHoroscope returns a horoscope for the given star sign
func GetStarSignHoroscope(sign, date string) (string, error) {

	// Get zodiac sign data
	signs := models.GetZodiacSigns()
	// Normalize input
	sign = strings.ToLower(strings.TrimSpace(sign))

	// Validate the sign
	signData, valid := signs[sign]
	if !valid {
		return "", fmt.Errorf("invalid star sign: %s", sign)
	}

	// Parse the date if provided, otherwise use today
	var targetDate time.Time
	var err error

	if date == "" {
		targetDate = time.Now()
	} else {
		targetDate, err = time.Parse("2006-01-02", date)
		if err != nil {
			return "", fmt.Errorf("invalid date format (use YYYY-MM-DD): %w", err)
		}
	}

	// Generate a horoscope based on the sign and date
	return generateHoroscope(signData, targetDate), nil
}

// generateHoroscope creates a horoscope reading based on the sign and date
func generateHoroscope(sign models.ZodiacSign, date time.Time) string {
	// Seed random generator with the date to ensure consistent horoscopes for the same day
	seed := date.Year()*10000 + int(date.Month())*100 + date.Day()

	rand.Seed(int64(seed))

	// Customize horoscope based on element
	var elementTheme string
	switch sign.Element {
	case "Fire":
		elementTheme = fireElementHoroscope()
	case "Earth":
		elementTheme = earthElementHoroscope()
	case "Air":
		elementTheme = airElementHoroscope()
	case "Water":
		elementTheme = waterElementHoroscope()
	}

	// Customize horoscope based on quality
	var qualityTheme string
	switch sign.Quality {
	case "Cardinal":
		qualityTheme = cardinalQualityHoroscope()
	case "Fixed":
		qualityTheme = fixedQualityHoroscope()
	case "Mutable":
		qualityTheme = mutableQualityHoroscope()
	}

	// Planetary ruler influence
	rulerInfluence := planetaryRulerHoroscope(sign.Ruler)

	// Get a specific characteristic for this sign
	signSpecific := signSpecificHoroscope(strings.ToLower(sign.Name))

	// Build the complete horoscope
	horoscope := fmt.Sprintf("As a %s, your %s nature is highlighted today. %s %s %s",
		sign.Name, sign.Element, elementTheme, qualityTheme, signSpecific)

	// Add planetary influence if date is within 3 days of current date
	now := time.Now()
	diff := now.Sub(date).Hours() / 24
	if diff >= -3 && diff <= 3 {
		horoscope += " " + rulerInfluence
	}

	return horoscope
}

// Element-based horoscope generators
func fireElementHoroscope() string {
	options := []string{
		"Your enthusiasm and passion can help you make significant progress on creative projects.",
		"Channel your natural energy into activities that inspire you and others around you.",
		"Your confidence is contagious today - use it to motivate yourself and those you care about.",
		"Trust your intuitive flashes of insight, as they're likely to guide you in the right direction.",
	}
	return options[rand.Intn(len(options))]
}

func earthElementHoroscope() string {
	options := []string{
		"Practical matters respond well to your grounded approach today.",
		"Your steady persistence will help you overcome obstacles that might discourage others.",
		"Focus on building security and stability in areas that matter most to you.",
		"Your natural connection to the physical world enhances your ability to create tangible results.",
	}
	return options[rand.Intn(len(options))]
}

func airElementHoroscope() string {
	options := []string{
		"Your mental clarity can help you solve complex problems with innovative solutions.",
		"Communication flows easily today - express your ideas and listen attentively to others.",
		"Making connections between seemingly unrelated concepts could lead to valuable insights.",
		"Your objective perspective helps you see situations from multiple angles.",
	}
	return options[rand.Intn(len(options))]
}

func waterElementHoroscope() string {
	options := []string{
		"Your emotional intelligence guides you in navigating interpersonal dynamics with grace.",
		"Trust your intuition when it comes to understanding unspoken feelings and needs.",
		"Creative inspiration flows naturally when you allow yourself to connect with your deeper emotions.",
		"Your compassionate nature helps you support others who may be struggling today.",
	}
	return options[rand.Intn(len(options))]
}

// Quality-based horoscope generators
func cardinalQualityHoroscope() string {
	options := []string{
		"Your initiative helps you start new projects with confidence and clarity.",
		"Taking the lead on important matters will come naturally to you today.",
		"Don't hesitate to be the catalyst for positive change in your environment.",
		"Your ability to take action decisively benefits both yourself and others.",
	}
	return options[rand.Intn(len(options))]
}

func fixedQualityHoroscope() string {
	options := []string{
		"Your natural determination helps you maintain focus on long-term objectives.",
		"Stability is your strength today - use it to create a solid foundation for future growth.",
		"Your loyalty and commitment strengthen important relationships in your life.",
		"Trust your ability to persevere through challenges with unwavering resolve.",
	}
	return options[rand.Intn(len(options))]
}

func mutableQualityHoroscope() string {
	options := []string{
		"Your adaptability allows you to navigate changing circumstances with ease.",
		"Being flexible with your approach helps you find creative solutions to unexpected issues.",
		"Your versatile nature is an asset when dealing with diverse situations and people.",
		"Embracing change rather than resisting it will lead to valuable growth opportunities.",
	}
	return options[rand.Intn(len(options))]
}

// Planetary ruler influences
func planetaryRulerHoroscope(ruler string) string {
	influences := map[string][]string{
		"Sun": {
			"The Sun illuminates your authentic self, helping you shine with genuine confidence.",
			"Your creative self-expression resonates strongly with others today.",
			"Focus on core aspects of your identity that bring you vitality and purpose.",
		},
		"Moon": {
			"The Moon heightens your emotional awareness, helping you understand subtle feelings.",
			"Pay attention to your intuitive rhythms and honor your need for emotional security.",
			"Nurturing connections with family or close friends brings special fulfillment today.",
		},
		"Mercury": {
			"Mercury sharpens your mental acuity, making this an excellent day for communication.",
			"Your ability to process information and express ideas is particularly strong today.",
			"Learning new skills or gathering information flows more easily under Mercury's influence.",
		},
		"Venus": {
			"Venus enhances your appreciation for beauty and harmony in all forms.",
			"Relationships benefit from your natural charm and diplomatic approach today.",
			"Creative and artistic pursuits are especially fulfilling under Venus's influence.",
		},
		"Mars": {
			"Mars energizes your actions, helping you pursue goals with determination.",
			"Channel your assertive energy into productive activities rather than conflicts.",
			"Physical activity helps you use Mars's influence constructively and release tension.",
		},
		"Jupiter": {
			"Jupiter expands your horizons and helps you see the bigger picture.",
			"Opportunities for growth appear when you maintain an optimistic outlook.",
			"Your generosity of spirit attracts positive experiences under Jupiter's influence.",
		},
		"Saturn": {
			"Saturn reinforces your discipline and helps you build lasting structures in your life.",
			"Taking responsibility seriously leads to earned rewards and respect.",
			"Patience with long-term processes pays off under Saturn's mature influence.",
		},
		"Uranus": {
			"Uranus inspires innovative thinking that breaks you free from limiting patterns.",
			"Unexpected insights can lead to breakthrough solutions to persistent problems.",
			"Embracing your uniqueness helps you make authentic choices aligned with your true self.",
		},
		"Neptune": {
			"Neptune heightens your imagination and spiritual sensitivity.",
			"Creative inspiration flows more easily when you allow yourself quiet reflection time.",
			"Compassionate understanding helps you connect with others on a deeper level.",
		},
		"Pluto": {
			"Pluto's transformative energy helps you release what no longer serves your growth.",
			"Deep psychological insights reveal hidden aspects of yourself or situations.",
			"Embracing necessary change, though challenging, leads to powerful renewal.",
		},
	}

	// Handle composite rulers
	if strings.Contains(ruler, "/") {
		parts := strings.Split(ruler, "/")
		// Randomly select one of the rulers
		ruler = parts[rand.Intn(len(parts))]
	}

	options, exists := influences[ruler]
	if !exists {
		// Fallback if ruler not found
		return "The planetary influences today support your natural tendencies and strengths."
	}

	return options[rand.Intn(len(options))]
}

// Sign-specific horoscope elements
func signSpecificHoroscope(sign string) string {
	specifics := map[string][]string{
		"aries": {
			"Focus your pioneering spirit on breaking new ground in areas that truly matter to you.",
			"Your courage helps you face challenges that others might avoid.",
			"Balance your natural independence with meaningful collaboration today.",
		},
		"taurus": {
			"Your appreciation for life's pleasures helps you find joy in simple moments today.",
			"Your patient approach allows things to unfold at their natural pace.",
			"Your resourcefulness helps you make the most of what you already have.",
		},
		"gemini": {
			"Your natural curiosity leads you to valuable information and interesting conversations.",
			"Expressing yourself through various channels keeps your versatile mind engaged.",
			"Making connections between different areas of knowledge gives you unique insights.",
		},
		"cancer": {
			"Your protective instincts help you create safe spaces for yourself and loved ones.",
			"Honoring your emotional needs strengthens your ability to care for others.",
			"Your intuitive understanding of others' feelings enhances meaningful connections.",
		},
		"leo": {
			"Sharing your natural warmth and generosity brings joy to those around you.",
			"Creative self-expression allows your authentic self to shine brightly.",
			"Your leadership abilities inspire others to discover their own strengths.",
		},
		"virgo": {
			"Your analytical skills help you improve systems and solve complex problems.",
			"Attention to meaningful details makes a significant difference in outcomes today.",
			"Your practical approach turns abstract ideas into workable solutions.",
		},
		"libra": {
			"Your diplomatic skills help resolve tensions and create harmony in challenging situations.",
			"Seeking balance in your activities enhances your overall well-being.",
			"Your appreciation for beauty and harmony brings refinement to everything you touch.",
		},
		"scorpio": {
			"Your emotional intensity gives you access to deeper levels of understanding and connection.",
			"Trust your intuitive insights about hidden dynamics in important situations.",
			"Your resilience helps you transform challenges into opportunities for growth.",
		},
		"sagittarius": {
			"Your optimistic outlook helps you see possibilities where others might see obstacles.",
			"Expanding your horizons through learning and travel brings valuable insights.",
			"Your adventurous spirit leads you to exciting new experiences and discoveries.",
		},
		"capricorn": {
			"Your disciplined approach helps you achieve long-term goals with lasting impact.",
			"Focus on building structures that support your ambitions and aspirations.",
			"Your patient persistence leads to steady progress and eventual success.",
		},
		"aquarius": {
			"Your innovative thinking helps you find unique solutions to persistent challenges.",
			"Connecting with like-minded people amplifies your ability to create positive change.",
			"Your humanitarian vision inspires collective progress and evolution.",
		},
		"pisces": {
			"Your compassionate nature helps you understand and heal emotional wounds.",
			"Tapping into your creative imagination reveals inspiring possibilities.",
			"Your spiritual sensitivity guides you toward meaningful experiences and connections.",
		},
	}

	options, exists := specifics[sign]
	if !exists {
		// Fallback if sign not found
		return "Your unique characteristics bring special insight to today's situations."
	}

	return options[rand.Intn(len(options))]
}
