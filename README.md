# Bearthropic

Bearthropic (bthropic) is a CLI tool that connects Claude AI with Bear notes, allowing you to have AI conversations that are automatically saved and formatted in Bear.

## Prerequisites

- [Bear](https://bear.app/) installed on your machine
- Go 1.20 or higher
- Claude API key (from Anthropic)

## Installation

### Option 1: Download Binary (Recommended)

```bash
# Download the latest release for your platform
wget https://github.com/mreider/bearthropic/releases/latest/download/bearthropic-$(uname -s)-$(uname -m) -O bthropic

# Make it executable
chmod +x bthropic

# Move to your PATH
sudo mv bthropic /usr/local/bin/
```

### Option 2: Build from Source

```bash
# Clone the repository
git clone https://github.com/matthewreider/bthropic
cd bthropic

# Build the binary
go build -o bthropic

# Optional: Move to your PATH
sudo mv bthropic /usr/local/bin/
```

## Getting Started

1. Initialize bthropic with your API key:
```bash
bthropic --init
```
You'll be prompted to enter your Claude API key. The key will be stored securely in `~/.bthropic/config.json`.

2. Start a conversation:
```bash
bthropic --start
```

## Usage

During an interactive session:

1. Type your question when prompted
2. Claude will respond and create a new Bear note with both your question and the response
3. You'll be asked if you want to clarify or modify the response
   - Answer 'y' to continue the conversation (the same note will be updated)
   - Answer 'n' to end the session

## Example Session

```
$ bthropic --start
Starting new session with Claude. Type your question:
> What are the main differences between Go and Rust?

[Claude responds and creates a Bear note]

Would you like to clarify or modify the response? (y/n): y
> Can you elaborate more on memory management?

[Claude updates the existing Bear note]

Would you like to clarify or modify the response? (y/n): n
```

## Building from Source

1. Ensure you have Go 1.20+ installed:
```bash
go version
```

2. Clone the repository:
```bash
git clone https://github.com/matthewreider/bthropic
cd bthropic
```

3. Install dependencies:
```bash
go mod download
```

4. Build the binary:
```bash
go build
```

5. Run the binary:
```bash
./bthropic --help
```

## License

MIT
