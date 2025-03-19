# Astrogo CLI

A command-line interface tool for astrology calculations and analysis, built with Go. This project serves as a sandbox for exploring Go CLI development and astrology calculations.

## Features

- Command-line interface for astrology calculations
- Interactive terminal UI using Bubble Tea
- Modular architecture for easy extension
- Built with modern Go practices and best practices

## Prerequisites

- Go 1.23.0 or higher
- Git
- GCC compiler and build tools
- Swiss Ephemeris library and data files

### Installing Swiss Ephemeris (Ubuntu/Debian)

```bash
# Install build tools
sudo apt-get update
sudo apt-get install -y build-essential gcc

# Install Swiss Ephemeris library and data
sudo apt-get install -y libswe-dev swe-standard-data
```

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/astrogo-cli.git
cd astrogo-cli
```

2. Install dependencies:
```bash
go mod download
```

3. Build the project:
```bash
go build
```

## Usage

```bash
astrogo-cli [command]
astrogo-cli starsign // Get your horoscope based on star sign
astrogo-cli db start // Start the database
```

## Project Structure

```
astrogo-cli/
├── cli/            # Command-line interface implementation
├── configs/        # Configuration files
├── data/          # Data files and resources
├── internal/      # Internal packages
├── pkg/           # Public packages
├── static/        # Static assets
└── tests/         # Test files
```

## Dependencies

Major dependencies include:
- `github.com/charmbracelet/bubbletea` - For interactive terminal UI
- `github.com/spf13/cobra` - For CLI command structure
- `github.com/pocketbase/pocketbase` - For data management

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[Add your license information here]

## Author

[Add your name/contact information here]