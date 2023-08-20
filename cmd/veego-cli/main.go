package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

// TO INSTALL THE EXECUTABLE RUN: go install
func main() {
    app := &cli.App{
        Commands: []*cli.Command{
            {
                Name:  "init",
                Usage: "Initialize a new Go Boilerplate",
                Action: func(c *cli.Context) error {
                    projectName := c.Args().First()  //==> Get the first argument which is the project name
                    if projectName == "" {
                        return fmt.Errorf(color.RedString("Project name is required"))
                    }

                    // Print to the terminal that the process as started
                    color.Green("Initializing Veego for setup...")
                    time.Sleep(4 * time.Second)

                    color.Green("Veego is running, ")

                    // Draw a simple box using ASCII characters
                    box := `
                    ╔══════════════╗
                    ║              ║
                    ║   VEEGO CLI  ║
                    ║              ║
                    ╚══════════════╝ , created with Go
                    `
                    fmt.Println(box)

                    currentDir, err := os.Getwd() //==> Get the current working directory
                    if err != nil {
                        return err
                    }
    
                    projectPath := filepath.Join(currentDir, projectName) 
                    err = os.Mkdir(projectPath, 0755)
                    if err != nil {
                        return err
                    }
    
                    // start creating the Go project
                    time.Sleep(4 * time.Second)
                    color.Green("Creating boilerplate: %s\n", projectName)

                    time.Sleep(4 * time.Second)
                    cmd := exec.Command("go", "mod", "init", projectName)
                    cmd.Dir = projectPath
                    cmd.Stdout = os.Stdout
                    cmd.Stderr = os.Stderr
                    if err := cmd.Run(); err != nil {
                        color.Red("Project initialization failed: %v\n", err)
                        return err
                    }
    
                    time.Sleep(6 * time.Second)
                    color.Green("Project %s initialized successfully!\n", projectName)
                    return nil
                },
            },
        },

    }

    if err := app.Run(os.Args); err != nil {
        fmt.Println(err)
    }
}
