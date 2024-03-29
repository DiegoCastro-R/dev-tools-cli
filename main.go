package main

import (
    "fmt"
    "os"
    "os/exec"
)

func executeShellScript(scriptPath string) error {
    cmd := exec.Command("sh", scriptPath)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

func main() {
    fmt.Println("==========================")
    fmt.Println("Welcome to JDK Installer CLI!")
    fmt.Println("Select your OS:")
    fmt.Println("1. 🐧 Ubuntu: .deb Packages")
    fmt.Println("2. Exit")
    fmt.Println("==========================")

    var osChoice int
    fmt.Print("Enter your choice: ")
    fmt.Scanln(&osChoice)

    switch osChoice {
    case 1:
        fmt.Println("👉 You selected Ubuntu - .deb Packages")
        fmt.Println("==========================")
        fmt.Println("Select the package to install:")
		fmt.Println("💻 Dev Tools:")
        fmt.Println("1. Install Java JDK")
        fmt.Println("2. Install IntelliJ IDE")
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@")
        fmt.Println("3. Back to main menu")
        fmt.Println("==========================")

        var packageChoice int
        fmt.Print("Enter your choice: ")
        fmt.Scanln(&packageChoice)

        switch packageChoice {
        case 1:
            fmt.Println("⏳ Installing Java JDK...")
            err := executeShellScript("./installers/ubuntu/java_jdk_install.sh")
            if err != nil {
                fmt.Println("Error installing JDK:", err)
            } else {
                fmt.Println("✅ Java JDK has been installed.")
            }
        case 2:
            fmt.Println("⏳ Installing IntelliJ IDE...")
            err := executeShellScript("./installers/ubuntu/intellij_community_install.sh")
            if err != nil {
                fmt.Println("Error IntelliJ:", err)
            } else {
                fmt.Println("✅ IntelliJ has been installed.")
            }
        case 3:
            // Do nothing, it will go back to the main menu
        default:
            fmt.Println("Invalid choice. Please try again.")
        }

    case 2:
        fmt.Println("Exiting...")
        os.Exit(0)
    default:
        fmt.Println("Invalid choice. Please try again.")
    }
}
