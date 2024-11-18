package main

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "log"
)

func Usage() {
	fmt.Println("Usage:", os.Args[0], "[OPTIONS]")
	fmt.Println("	Options:")
	fmt.Println("	-help, -h (Option)			Shows this help message unless a option is specified")
	fmt.Println("	-wallpaper, -w <path>			Wallpaper Path")
	fmt.Println("	-bar, -b <etc>				Which Bar you use")
	fmt.Println("	-scheme, -m <dark/light>		Colorscheme")
	fmt.Println("	-reload, -r				Refreshes Components such as bar")
	fmt.Println("	-echo					Echoes a environment variable")
}

var cache string 
var wallpaper string                                    
var cachedwallpaper string  
var SESSION = os.Getenv("XDG_SESSION_TYPE")
var util string

func init() {
    // Attempt to retrieve the user's home directory
    home, err := os.UserHomeDir()
    // Check if there was an error in fetching the home directory
    if err != nil {
        // Print the error message to the console
        fmt.Println("Error fetching home directory:", err)
        // Exit the program with a non-zero status to indicate failure
        os.Exit(1)
    }
    // Construct the cache path by joining the home directory
    // with ".cache" and "GoSys" to create a specific cache directory
    cache = filepath.Join(home, ".cache", "GoSys")
}

// This function is used to set the wallpaper based on the session type
// It should be called after the wallpaper path has been set
func FuncUtil() {
    // Check if the session type is Wayland
    if SESSION == "wayland" {
        util = "swaybg"
        
        // Create a command to run swaybg with the following arguments:
        // * -i specifies the input image
        // * wallpaper is the path to the wallpaper image
        // * > /dev/null 2>&1 & runs the command in the background and redirects stdout and stderr to /dev/null
        utilCmd := exec.Command("nohup", "sh", "-c", util+" -i "+wallpaper+" > /dev/null 2>&1 &")
        
        // Run the command and check for errors
        err := utilCmd.Run() 
        if err != nil {
            log.Fatal(err) 
        }
    // If the session type is X11
    } else if SESSION == "x11" {
        util = "feh"
        
        // Create a command to run feh with the following arguments:
        // * --bg-scale specifies the image is to be scaled to the desktop size
        // * wallpaper is the path to the wallpaper image
        // * > /dev/null 2>&1 & runs the command in the background and redirects stdout and stderr to /dev/null
        utilCmd := exec.Command("nohup", "sh", "-c", util+" --bg-scale "+wallpaper+" > /dev/null 2>&1 &")
        
        // Run the command and check for errors
        err := utilCmd.Run()  
        if err != nil {
            log.Fatal(err)  
        }
    } else {
        fmt.Println("Unknown session type:", SESSION)
    }
}

// This function takes a string of text as input and returns a new string with any blank or empty lines removed.
// The purpose of this function is to remove any unnecessary whitespace from command output when printing to the console.
func strip(output string) string {
	// Split the input string into an array of strings, one for each line of text
	lines := strings.Split(output, "\n")

	// Initialize an empty array of strings that will store the non-empty lines
	var nonEmptyLines []string

	// Iterate over the array of lines
	for _, line := range lines {
		// Check if the line is not empty or blank
		if strings.TrimSpace(line) != "" {
			// If the line is non-empty, add it to the array of non-empty lines
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}

	// Join the non-empty lines back together with a newline character in between each line
	return strings.Join(nonEmptyLines, "\n")
}

// This function runs a shell command with the given arguments and handles any errors.
// It takes 3 parameters:
// 1. command: the command to run
// 2. suppressOutput: a boolean indicating whether to suppress the output of the command
// 3. args: a variable number of arguments to pass to the command
func execShell(command string, suppressOutput bool, args ...string) {
	// Create a new command to run
	cmd := exec.Command(command, args...)

	// Run the command and capture its output and any errors
	output, err := cmd.CombinedOutput()

	// Check if there was an error
	if err != nil {
		// If suppressOutput is false, print the command output even if there was an error
		if !suppressOutput {
			// Print the command output and any error message
			fmt.Println("Command output:", strip(string(output)))
		}

		// Print the error message
		fmt.Println("Error:", err)

		// Return to prevent further execution of the function
		return
	}

	// If there was no error and suppressOutput is false, print the command output
	if !suppressOutput {
		// Print the command output, stripped of any blank or empty lines
		fmt.Println(strip(string(output)))
	}
}
