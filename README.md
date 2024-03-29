# Dev Tools CLI

Dev Tools CLI is a command-line interface (CLI) tool written in Go that simplifies the installation of various development tools for Go programming on different operating systems.

## Features

- Installs essential development tools for Go programming.
- Provides a user-friendly menu interface.
- Supports multiple operating systems, including Ubuntu, macOS, and Windows.

## Prerequisites

- Go programming language should be installed on your system. You can download and install it from the [official Go website](https://golang.org/dl/).

## Installation

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/diegocastro-r/dev-tools-cli.git
   ```

2. Navigate to the project directory:

   ```bash
   cd dev-tools-cli
   ```

3. Build the Go program:

   ```bash
   go build -o dev-tools main.go
   ```

4. Run the CLI tool:

   ```bash
   ./dev-tools
   ```

## Usage

1. Upon running the CLI tool, you will be presented with a menu offering options to select your operating system and choose the development tools to install.
2. Choose your operating system from the menu (e.g., Ubuntu, macOS, or Windows).
3. Select the development tools you want to install (e.g., GoLand IDE, Docker, or VS Code).
4. Follow the on-screen prompts to complete the installation process.

## Supported Tools

- GoLand IDE
- Docker
- Visual Studio Code (VS Code)
- Git
- And more...

## Contributing

Contributions are welcome! If you find any bugs or have suggestions for improvement, please feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
