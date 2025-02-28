#!/bin/bash
set -euo pipefail

# Global OS and ARCH variables
OS=$(uname -s)
ARCH=$(uname -m)

# Check for required commands
for cmd in wget sudo; do
  if ! command -v "$cmd" &> /dev/null; then
    echo "Error: '$cmd' is required but not installed."
    exit 1
  fi
done

# Verify that a checksum utility exists
if ! command -v sha256sum &> /dev/null && ! command -v shasum &> /dev/null; then
  echo "Error: Neither 'sha256sum' nor 'shasum' is available. Install one to continue."
  exit 1
fi

# Function to verify file checksum
verify_checksum() {
  local file="$1"
  local checksum_file="$2"
  local actual_checksum
  local expected_checksum

  if command -v sha256sum &> /dev/null; then
    actual_checksum=$(sha256sum "$file" | awk '{print $1}')
  else
    actual_checksum=$(shasum -a 256 "$file" | awk '{print $1}')
  fi

  expected_checksum=$(cut -d ' ' -f 1 "$checksum_file")

  if [[ "$actual_checksum" != "$expected_checksum" ]]; then
    echo "Checksum verification failed for $file"
    exit 1
  fi
}

# Function to check and install Go
install_go() {
  local GO_VERSION="1.19.3"
  local go_archive=""
  
  if ! command -v go &> /dev/null; then
    echo "Go is not installed. Installing Go..."
    
    if [[ "$OS" == "Linux" ]]; then
      go_archive="go${GO_VERSION}.linux-amd64.tar.gz"
    elif [[ "$OS" == "Darwin" ]]; then
      go_archive="go${GO_VERSION}.darwin-amd64.tar.gz"
    else
      echo "Unsupported OS: $OS"
      exit 1
    fi

    # Download the archive and its checksum file
    wget "https://go.dev/dl/${go_archive}" -O "$go_archive"
    wget "https://go.dev/dl/${go_archive}.sha256" -O "${go_archive}.sha256"

    # Verify the downloaded archive
    verify_checksum "$go_archive" "${go_archive}.sha256"

    # Remove any previous Go installation
    sudo rm -rf /usr/local/go

    # Extract the archive to /usr/local
    sudo tar -C /usr/local -xzf "$go_archive"

    # Clean up downloaded files
    rm "$go_archive" "${go_archive}.sha256"

    # Update PATH for the current session and future sessions
    export PATH="$PATH:/usr/local/go/bin"
    if [[ "$OS" == "Linux" ]]; then
      if ! grep -q "/usr/local/go/bin" ~/.bashrc; then
        echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
      fi
    elif [[ "$OS" == "Darwin" ]]; then
      if ! grep -q "/usr/local/go/bin" ~/.zshrc; then
        echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc
      fi
    fi

    echo "Go installed successfully."
  else
    echo "Go is already installed."
  fi
}

# Function to check and install Git
install_git() {
  if ! command -v git &> /dev/null; then
    echo "Git is not installed. Installing Git..."
    if [[ "$OS" == "Linux" ]]; then
      sudo apt update && sudo apt install -y git
    elif [[ "$OS" == "Darwin" ]]; then
      if command -v brew &> /dev/null; then
        brew install git
      else
        echo "Homebrew is not installed. Please install Homebrew first."
        exit 1
      fi
    else
      echo "Unsupported OS: $OS"
      exit 1
    fi
    echo "Git installed successfully."
  else
    echo "Git is already installed."
  fi
}

# Main execution function
main() {
  install_go
  install_git
}

main