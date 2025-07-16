# Precious Metals Tracker

A simple CLI application for tracking your precious metal holdings (Gold, Silver, and Platinum) with real-time pricing data.

## Features

- **Track Multiple Metals**: Support for Gold, Silver, and Platinum holdings
- **Real-time Pricing**: Fetches current spot prices from goldapi.io
- **Local Storage**: SQLite database for secure local data storage  
- **Flexible Input**: Interactive wizard or direct CLI arguments
- **Multiple Output Formats**: Table or JSON output
- **Value Calculations**: Track total weight, current value, and purchase price vs. spot price
- **Caching**: Intelligent price caching to minimize API calls

## Prerequisites

- Go 1.23.1 or later
- A goldapi.io API key (required for price data)

## Installation

### Option 1: Build from Source

```bash
# Clone the repository
git clone https://github.com/robert430404/precious-metals-tracker.git
cd precious-metals-tracker

# Build the application
make build

# The binary will be created as 'precious-metals-tracker'
```

### Option 2: Direct Go Install

```bash
go install github.com/robert430404/precious-metals-tracker@latest
```

## Setup

### 1. Get a goldapi.io API Key

1. Visit [goldapi.io](https://www.goldapi.io/)
2. Sign up for an account
3. Get your API key from the dashboard

### 2. Initialize the Application

```bash
# Interactive setup (recommended for first-time users)
./precious-metals-tracker init

# Or provide the API key directly
./precious-metals-tracker init --api-key YOUR_API_KEY
```

This command will:
- Create the configuration directory (`~/.local/share/precious-metals-tracker`)
- Set up the SQLite database
- Store your API key securely

### 3. Configuration

The application stores data in:
- **Config Directory**: `~/.local/share/precious-metals-tracker/`
- **Database**: `~/.local/share/precious-metals-tracker/precious-metals-tracker.sqlite`
- **Cache Files**: Price data cached for 24 hours to minimize API calls

You can override the config directory by setting the `PRECIOUS_METALS_TRACKER_DATA_DIR` environment variable.

## Usage

### Adding Holdings

#### Interactive Mode (Wizard)
```bash
./precious-metals-tracker holding --add
```

The wizard will prompt you for:
- Product name
- Purchase price
- Purchase source
- Spot price at time of purchase
- Number of units
- Weight per unit (in troy ounces)
- Metal type (Gold, Silver, or Platinum)

#### CLI Mode (No Wizard)
```bash
./precious-metals-tracker holding --add \
  --name "American Silver Eagle" \
  --price "35.00" \
  --source "Local Coin Shop" \
  --spot-price "24.50" \
  --units "10" \
  --weight "1.0" \
  --type "Silver"
```

### Viewing Holdings

#### List All Holdings
```bash
# Table format (default)
./precious-metals-tracker holding --list

# JSON format
./precious-metals-tracker holding --list --format json
```

#### Get Current Values
```bash
# Shows current value, spot price, and total weight by metal type
./precious-metals-tracker holding --value

# JSON output
./precious-metals-tracker holding --value --format json
```

### Deleting Holdings

```bash
# Interactive deletion (you'll be prompted to select which holding to delete)
./precious-metals-tracker holding --delete

# Direct deletion by ID
./precious-metals-tracker holding --delete --id HOLDING_ID
```

## Command Reference

### `init`
Initialize the application and set up the database.

```bash
precious-metals-tracker init [flags]
```

**Flags:**
- `-k, --api-key string`: goldapi.io API key (bypasses interactive prompt)

### `holding`
Manage your precious metal holdings.

```bash
precious-metals-tracker holding [flags]
```

**Action Flags** (choose one):
- `-a, --add`: Add a new holding
- `-l, --list`: List all holdings  
- `-d, --delete`: Delete a holding
- `-v, --value`: Show current values and totals

**Output Flags:**
- `-f, --format string`: Output format ("table" or "json", default: "table")

**Add Holding Flags** (bypass wizard):
- `-n, --name string`: Product name
- `-p, --price string`: Purchase price
- `-s, --source string`: Purchase source
- `--spot-price string`: Spot price at time of purchase
- `-u, --units string`: Number of units
- `-w, --weight string`: Weight per unit in troy ounces
- `-t, --type string`: Metal type ("Gold", "Silver", or "Platinum")

**Delete Holding Flags:**
- `--id string`: Holding ID to delete (bypasses interactive selection)

## Examples

### Complete Workflow Example

```bash
# 1. Initialize the application
./precious-metals-tracker init --api-key your_goldapi_key

# 2. Add some holdings
./precious-metals-tracker holding --add \
  --name "1 oz Gold Buffalo" \
  --price "2100.00" \
  --source "APMEX" \
  --spot-price "2050.00" \
  --units "5" \
  --weight "1.0" \
  --type "Gold"

./precious-metals-tracker holding --add \
  --name "American Silver Eagle" \
  --price "35.00" \
  --source "Local Shop" \
  --spot-price "24.50" \
  --units "20" \
  --weight "1.0" \
  --type "Silver"

# 3. View your holdings
./precious-metals-tracker holding --list

# 4. Check current values
./precious-metals-tracker holding --value
```

### JSON Output Example

```bash
./precious-metals-tracker holding --value --format json
```

Output:
```json
[
  {
    "type": "Silver",
    "currentValue": "$520.00",
    "currentSpotPrice": "$26.00",
    "totalHoldingWeight": "20.00oz"
  },
  {
    "type": "Gold", 
    "currentValue": "$10250.00",
    "currentSpotPrice": "$2050.00",
    "totalHoldingWeight": "5.00oz"
  },
  {
    "type": "Platinum",
    "currentValue": "$0.00", 
    "currentSpotPrice": "$980.00",
    "totalHoldingWeight": "0.00oz"
  }
]
```

## Development

### Building

```bash
# Format code
make format

# Run tests
make test

# Build binary
make build
```

### Dependencies

Key dependencies include:
- **cobra**: CLI framework
- **promptui**: Interactive prompts
- **sqlite**: Local database storage
- **goldapi.io**: Real-time precious metals pricing

### Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## API Information

This application uses [goldapi.io](https://www.goldapi.io/) for real-time precious metals pricing:

- **Endpoints Used**:
  - Gold: `/api/XAU/USD`
  - Silver: `/api/XAG/USD` 
  - Platinum: `/api/XPT/USD`

- **Caching**: Prices are cached locally for 24 hours to minimize API usage
- **Rate Limits**: Respect goldapi.io's rate limits (check their documentation)

## License

This project is licensed under the terms specified in the LICENSE file.

## Support

If you encounter issues:

1. Ensure you have a valid goldapi.io API key
2. Check that your config directory has proper permissions
3. Verify your Go version meets the minimum requirement (1.23.1+)
4. Review the application logs for specific error messages

For bugs or feature requests, please open an issue on GitHub.
