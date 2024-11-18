package main

import (
	"fmt"
	"os"
	"os/exec"
	"log"
)

func main() {
	// Loop through the command line arguments
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		// Get the current argument
		arg := args[i]
		// Switch statement to handle different arguments
		switch arg {
		// Help argument
		case "-help", "-h":
			// Print the usage information
			Usage()
		// Wallpaper argument
		case "-wallpaper", "-w":
			// Check if the next argument is the wallpaper path
			if i+1 < len(args) {
				// Set the wallpaper path
				wallpaper = args[i+1]
				// Create a command to copy the wallpaper to the cache
				// directory
				cpCmd := exec.Command("cp", wallpaper, cache)
				// Run the command and check for errors
				if err := cpCmd.Run(); err != nil {
					// Print the error message to the console
					log.Fatal(err)
				}
				// Print a message indicating that the wallpaper path has been set
				fmt.Println("Wallpaper Path set to:", wallpaper)
				// Call the function to set the wallpaper based on the session type
				FuncUtil()
				// Increment the loop counter to skip the wallpaper path
				i++
			} else {
				// Print an error message if the wallpaper path is not provided
				fmt.Println("ERROR: No path defined")
				// Exit the program with a non-zero status to indicate failure
				return
			}
		default:
			// Print a message if an unknown argument is encountered
			fmt.Printf("Unknown operation: %s\n", arg)
		}
	}
}
