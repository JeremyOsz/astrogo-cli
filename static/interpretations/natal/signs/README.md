# Zodiac Sign Interpretations

This directory contains detailed astrological interpretations for all twelve zodiac signs. Each sign is stored in its own JSON file for easy maintenance and selective loading.

## Directory Structure

```
signs/
├── README.md           # This documentation file
├── schema.json         # JSON schema for validating sign files
├── index.json          # Index of all available signs with metadata
├── aries.json          # Aries sign interpretations
├── taurus.json         # Taurus sign interpretations
├── gemini.json         # Gemini sign interpretations
├── cancer.json         # Cancer sign interpretations
├── leo.json           # Leo sign interpretations
├── virgo.json         # Virgo sign interpretations
├── libra.json         # Libra sign interpretations
├── scorpio.json       # Scorpio sign interpretations
├── sagittarius.json   # Sagittarius sign interpretations
├── capricorn.json     # Capricorn sign interpretations
├── aquarius.json      # Aquarius sign interpretations
└── pisces.json        # Pisces sign interpretations
```

## File Structure

Each sign file follows the schema defined in `schema.json` and contains:

1. General description of the sign
2. Key strengths
3. Key weaknesses
4. Rising sign interpretation
5. Sun sign interpretation with house placements
6. Moon sign interpretation

## Usage

### Loading a Single Sign

```javascript
const signData = require('./signs/aries.json');
```

### Loading All Signs

```javascript
const index = require('./signs/index.json');
const signs = {};
for (const [name, metadata] of Object.entries(index.signs)) {
  signs[name] = require(`./signs/${metadata.file}`);
}
```

### Validating Sign Files

Use the schema in `schema.json` to validate sign files:

```javascript
const Ajv = require('ajv');
const ajv = new Ajv();
const schema = require('./signs/schema.json');
const validate = ajv.compile(schema);
```

## Metadata

The `index.json` file contains additional metadata for each sign:
- Element (fire, earth, air, water)
- Modality (cardinal, fixed, mutable)
- Ruling planet
- Degree range
- Order in the zodiac

## Contributing

When adding or modifying sign interpretations:
1. Ensure the file follows the schema defined in `schema.json`
2. Update the `last_updated` field in `index.json`
3. Maintain consistent formatting and style
4. Test the changes with the schema validator

## Version Control

Each sign file is tracked individually, making it easy to:
- Review changes to specific signs
- Roll back changes if needed
- Collaborate on different signs simultaneously 