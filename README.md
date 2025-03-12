# Bearthropic

Bearthropic (bthropic) is a CLI tool that connects Claude AI with Bear notes, allowing you to have AI conversations that are automatically saved and formatted in Bear.

![bearthropic demo](docs/bthropic.gif)

## Prerequisites

- [Bear](https://bear.app/) installed on your machine
- Go 1.20 or higher
- Claude API key (from Anthropic)

## Installation

To install bthropic, run the following command in your terminal:

```bash
curl -s https://raw.githubusercontent.com/mreider/bearthropic/main/install.sh | bash
```


This command downloads the latest version of the install script, detects your architecture, and installs the binary to `/usr/local/bin/bthropic`.

## How it works

Initialize bthropic with your API key:
```bash
bthropic --init
```
You'll be prompted to enter your Claude API key. The key will be stored securely in `~/.bthropic/config.json`.

Start a conversation:
```bash
bthropic --start
```

Uninstall bthropic:
```bash
bthropic --destroy
```

## Example Session

```
$ bthropic 
Starting new session with Claude.
Enter your question: How many chucks could a wood chuck chuck?
Note created. Review it; Ctrl-C or type 'end' to finish, or type changes:
Enter your question: end
```
