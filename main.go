package main

import (
	"fmt"
	"os"
	"path/filepath"
	es "github.com/furiousman59/GO59"
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
                // Create a command to copy the wallpaper to the cache directory
                temp := filepath.Join(cache, es.StringRNG(12) + ".jpg")
                go es.ExecShell("cp", true, wallpaper, temp)
                // Cache the wallpaper path
                go es.ExecShell("nohup", true, "sh", "-c", fmt.Sprintf("echo %s > %s", temp, cachewp))

                // Switch to cached wallpaper
                wallpaper = temp

                // Call the backend function to set the wallpaper based on the session type
                FuncUtil()

                // Print the wallpaper path
                fmt.Println("Wallpaper Path set to:", wallpaper)

                // Increment the loop counter to skip the wallpaper path
                i++
            } else {
				// Print an error message if the wallpaper path is not provided
				fmt.Println("Error: No wallpaper path provided")
				return
			}
        // Default case for unknown arguments
        default:
            fmt.Printf("Unknown operation: %s\n", arg)
        }
    }
}

