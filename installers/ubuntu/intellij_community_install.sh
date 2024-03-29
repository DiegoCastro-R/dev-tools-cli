#!/bin/bash

# Function to check if IntelliJ IDEA Community Edition is already installed
check_intellij_community_installed() {
    if [ -d "/opt/idea-community" ]; then
        echo "IntelliJ IDEA Community Edition is already installed."
        return 0
    else
        return 1
    fi
}

# Check if IntelliJ IDEA Community Edition is already installed
if check_intellij_community_installed; then
    exit 0
fi

# Create a temporary directory
temp_dir=$(mktemp -d)

# Download IntelliJ IDEA Community Edition tar.gz archive to temporary file
temp_file="$temp_dir/idea-community.tar.gz"
wget -O "$temp_file" "https://download.jetbrains.com/idea/ideaIC-2021.3.2.tar.gz"

# Extract the downloaded archive to /opt directory
sudo tar -xf "$temp_file" -C /opt/

# Rename the extracted directory to idea-community
sudo mv /opt/idea-* /opt/idea-community

# Create a symbolic link for easy access
sudo ln -s /opt/idea-community/bin/idea.sh /usr/local/bin/idea

# Clean up temporary directory
rm -r "$temp_dir"

# Print installation status
echo "IntelliJ IDEA Community Edition installation completed. Run 'idea' command to launch IntelliJ IDEA."
