#!/usr/bin/env bash

# This script downloads the latest release binary from your GitHub repo,
# detects if the architecture is arm64 or x86_64, renames the binary to "bthropic",
# makes it executable, and moves it to /usr/local/bin.

# Only work on Darwin (MacOS)
OS=$(uname -s)
if [ "$OS" != "Darwin" ]; then
    echo "This installer is only for macOS."
    exit 1
fi

# Determine architecture
ARCH=$(uname -m)
if [ "$ARCH" == "arm64" ]; then
    BINARY_NAME="bthropic-darwin-arm64"
elif [ "$ARCH" == "x86_64" ]; then
    BINARY_NAME="bthropic-darwin-amd64"
else
    echo "Unsupported architecture: $ARCH"
    exit 1
fi

# Set the base URL for your GitHub release downloads.
# Update this URL to match your repository's download URL.
REPO_URL="https://github.com/mreider/bearthropic/releases/latest/download"
DOWNLOAD_URL="${REPO_URL}/${BINARY_NAME}"

echo "Downloading $DOWNLOAD_URL ..."
curl -L -o bthropic "$DOWNLOAD_URL"
if [ $? -ne 0 ]; then
    echo "Download failed."
    exit 1
fi

chmod +x bthropic

echo "Installing bthropic to /usr/local/bin ..."
sudo mv bthropic /usr/local/bin/bthropic
if [ $? -eq 0 ]; then
    echo "Installation successful!"
else
    echo "Installation failed. Please check your permissions."
    exit 1
fi
