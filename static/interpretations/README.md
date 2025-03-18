# Astrological Interpretations

This directory contains comprehensive astrological interpretations organized by type. Each category is stored in its own directory with a consistent structure.

## Directory Structure

```
interpretations/
├── README.md           # This documentation file
├── schema.json         # Common schema definitions
├── natal/             # Natal chart interpretations
│   ├── README.md      # Natal-specific documentation
│   ├── schema.json    # Natal-specific schema
│   ├── index.json     # Index of natal interpretations
│   ├── signs/         # Zodiac sign interpretations
│   │   ├── README.md  # Signs-specific documentation
│   │   ├── schema.json
│   │   ├── index.json
│   │   └── *.json     # Individual sign files
│   ├── planets/       # Planetary interpretations
│   │   ├── README.md
│   │   ├── schema.json
│   │   ├── index.json
│   │   └── *.json     # Individual planet files
│   ├── houses/        # House interpretations
│   │   ├── README.md
│   │   ├── schema.json
│   │   ├── index.json
│   │   └── *.json     # Individual house files
│   └── aspects/       # Aspect interpretations
│       ├── README.md
│       ├── schema.json
│       ├── index.json
│       └── *.json     # Individual aspect files
├── synastry/          # Synastry interpretations
│   ├── README.md
│   ├── schema.json
│   ├── index.json
│   └── *.json         # Relationship aspect files
├── transits/          # Transit interpretations
│   ├── README.md
│   ├── schema.json
│   ├── index.json
│   └── *.json         # Transit aspect files
└── progressions/      # Progression interpretations
    ├── README.md
    ├── schema.json
    ├── index.json
    └── *.json         # Progression aspect files
```

## Categories

### Natal Chart Interpretations
- **Signs**: Interpretations for each zodiac sign in different positions
- **Planets**: Interpretations for each planet's influence
- **Houses**: Interpretations for each astrological house
- **Aspects**: Interpretations for planetary aspects

### Synastry Interpretations
- Relationship compatibility
- Planetary interactions between charts
- Composite chart interpretations

### Transit Interpretations
- Current planetary movements
- Timing of events
- Period influences

### Progression Interpretations
- Secondary progressions
- Solar arc directions
- Age-related developments

## Common Structure

Each category follows a consistent structure:
1. `schema.json`: Defines the data structure
2. `index.json`: Lists available interpretations with metadata
3. Individual JSON files for each interpretation
4. Category-specific README with detailed documentation

## Usage

### Loading Interpretations

```javascript
// Load a specific category
const signs = require('./natal/signs');
const planets = require('./natal/planets');
const houses = require('./natal/houses');
const aspects = require('./natal/aspects');

// Load all natal interpretations
const natal = {
  signs,
  planets,
  houses,
  aspects
};
```

### Validation

Each category includes a schema for validation:

```javascript
const Ajv = require('ajv');
const ajv = new Ajv();
const schema = require('./natal/signs/schema.json');
const validate = ajv.compile(schema);
```

## Contributing

When adding or modifying interpretations:
1. Follow the established schema for each category
2. Update the relevant index.json file
3. Maintain consistent formatting and style
4. Test changes with the schema validator
5. Update documentation as needed

## Version Control

Each interpretation file is tracked individually, making it easy to:
- Review changes to specific interpretations
- Roll back changes if needed
- Collaborate on different categories simultaneously
- Maintain a clear history of updates

## Additional Documentation

Each category has its own README with specific details about:
- File structure
- Data format
- Usage examples
- Category-specific guidelines 