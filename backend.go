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


