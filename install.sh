#!/bin/bash

set -e  # Exit on error

# Define variables
APP_NAME="md"
INSTALL_DIR="/usr/local/bin"
GITHUB_REPO="zivoxRoot/markdown-reader"
LATEST_RELEASE="0.0.0"
DOWNLOAD_URL="https://github.com/$GITHUB_REPO/raw/refs/heads/main/bin/linux/md"

# Download and install the executable
echo "Downloading $APP_NAME and copying it to $INSTALL_DIR..."
sudo curl -L -o "$INSTALL_DIR/$APP_NAME" "$DOWNLOAD_URL"

# Set permissions
sudo chmod +x "$INSTALL_DIR/$APP_NAME"

# Ensure the directory is in PATH
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo "Adding $INSTALL_DIR to PATH..."
    echo 'export PATH="$INSTALL_DIR:$PATH"' >> "$HOME/.bashrc"
    echo 'export PATH="$INSTALL_DIR:$PATH"' >> "$HOME/.zshrc"
    export PATH="$INSTALL_DIR:$PATH"
fi

echo "$APP_NAME installed successfully in $INSTALL_DIR"
echo "Run '$APP_NAME' to start using it!"
