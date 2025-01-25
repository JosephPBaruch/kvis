#!/bin/bash

# Define the path to the kvis.sh script
SCRIPT_PATH="./kvis.sh"

# Ensure the script has execute permissions
chmod +x "$SCRIPT_PATH"

# Create a symbolic link to the script in /usr/local/bin
sudo ln -s "$SCRIPT_PATH" /usr/local/bin/kvis

echo "Symbolic link created. You can now run the script by typing 'kvis'."