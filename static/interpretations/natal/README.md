# Natal Chart Interpretations

This directory contains interpretations for natal chart components, including signs, planets, houses, and aspects. Each category follows a consistent structure with its own schema and documentation.

## Directory Structure

```
natal/
├── README.md           # This documentation file
├── schema.json         # Natal-specific schema definitions
├── signs/             # Zodiac sign interpretations
│   ├── README.md      # Signs-specific documentation
│   ├── schema.json    # Signs schema
│   ├── index.json     # Signs index
│   └── *.json         # Individual sign files
├── planets/           # Planetary interpretations
│   ├── README.md      # Planets-specific documentation
│   ├── schema.json    # Planets schema
│   ├── index.json     # Planets index
│   └── *.json         # Individual planet files
├── houses/            # House interpretations
│   ├── README.md      # Houses-specific documentation
│   ├── schema.json    # Houses schema
│   ├── index.json     # Houses index
│   └── *.json         # Individual house files
└── aspects/           # Aspect interpretations
    ├── README.md      # Aspects-specific documentation
    ├── schema.json    # Aspects schema
    ├── index.json     # Aspects index
    └── *.json         # Individual aspect files
```

## Categories

### Signs
- Interpretations for each zodiac sign
- Includes rising, sun, and moon sign interpretations
- House-specific interpretations for each sign

### Planets
- Interpretations for each planet's influence
- Includes dignity and aspect interpretations
- House-specific interpretations for each planet

### Houses
- Interpretations for each astrological house
- Sign-specific interpretations for each house
- Planet-specific interpretations for each house

### Aspects
- Interpretations for planetary aspects
- Sign-specific aspect interpretations
- House-specific aspect interpretations

## Usage

### Loading All Natal Interpretations

```javascript
const natal = {
  signs: require('./signs'),
  planets: require('./planets'),
  houses: require('./houses'),
  aspects: require('./aspects')
};
```

### Loading Specific Categories

```javascript
const signs = require('./signs');
const planets = require('./planets');
const houses = require('./houses');
const aspects = require('./aspects');
```

## Validation

Each category includes its own schema for validation:

```javascript
const Ajv = require('ajv');
const ajv = new Ajv();

// Load category-specific schemas
const signsSchema = require('./signs/schema.json');
const planetsSchema = require('./planets/schema.json');
const housesSchema = require('./houses/schema.json');
const aspectsSchema = require('./aspects/schema.json');

// Create validators
const validateSigns = ajv.compile(signsSchema);
const validatePlanets = ajv.compile(planetsSchema);
const validateHouses = ajv.compile(housesSchema);
const validateAspects = ajv.compile(aspectsSchema);
```

## Contributing

When adding or modifying natal interpretations:
1. Follow the established schema for each category
2. Update the relevant index.json file
3. Maintain consistent formatting and style
4. Test changes with the schema validator
5. Update documentation as needed

## Additional Documentation

Each category has its own README with specific details about:
- File structure
- Data format
- Usage examples
- Category-specific guidelines 