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

// strip takes a string output and removes any blank lines
// It does this by first splitting the string into individual lines
// Then it iterates over each line, checking if the line is empty
// or contains only whitespace characters. If the line is empty
// or contains only whitespace characters, it is skipped. All
// non-empty lines are then joined back together with newline
// characters in between to form the final output.
func strip(output string) string {
	// Split the string into individual lines
	lines := strings.Split(output, "\n")
	
	// Create a new slice to hold the non-empty lines
	var nonEmptyLines []string
	
	// Iterate over each line
	for _, line := range lines {
		// Trim the line of any leading or trailing whitespace
		trimmedLine := strings.TrimSpace(line)
		
		// Check if the line is empty or contains only whitespace
		if trimmedLine != "" {
			// If the line is not empty, add it to the new slice
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}
	
	// Join all the non-empty lines back together with newline
	// characters in between
	return strings.Join(nonEmptyLines, "\n")
}

// execShell runs a command with the given arguments and prints the output
// to the console, stripping out any blank lines in the output.
// If there is an error running the command, it is printed to the console.
func execShell(command string, args ...string) {
	// Create a new exec.Cmd object with the given command and arguments
	cmd := exec.Command(command, args...)
	
	// Run the command and capture its output, both stdout and stderr
	output, err := cmd.CombinedOutput()
	
	// Convert the output from a byte slice to a string
	outputStr := string(output)
	
	// Strip out any blank lines from the output
	strippedOutput := strip(outputStr)
	
	// Print the stripped output to the console
	fmt.Println(strippedOutput)
	
	// Check if there was an error running the command
	if err != nil {
		// If there was an error, print it to the console
		fmt.Println("Error:", err)
	}
}
