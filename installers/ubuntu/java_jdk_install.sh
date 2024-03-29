#!/bin/bash

# Function to check if Java JDK is installed
check_java_installed() {
    if java -version >/dev/null 2>&1; then
        echo "Java JDK is already installed."
        return 0
    else
        return 1
    fi
}

# Check if Java JDK is already installed
if check_java_installed; then
    exit 0
fi

# Update package index
sudo apt update

# Install Java JDK
sudo apt install default-jdk -y

# Verify Java installation
java -version

# Print installation status
echo "Java JDK installation completed."
