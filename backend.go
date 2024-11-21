package main

import (
    "fmt"
    "os"
    es "github.com/furiousman59/GO59" // es.ExecShell es.StringRNG es.FileExists es.Join es.Read
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

func CacheDir() {
    if es.FileExists(cache) == false {
        // Attempt to create the cache directory
        err := os.MkdirAll(cache, 0755)
        // Check if there was an error in creating the directory
        if err != nil {
        // Print the error message to the console
        fmt.Println("Error creating cache directory:", err)
        // Exit the program with a non-zero status to indicate failure
        os.Exit(1)
    }
    if !es.FileExists(cachewp) {
        file, err := os.Create(cachewp)
        if err != nil {
            fmt.Println("Error creating cache file:", err)
            return
        }
        defer file.Close() // Ensure the file is closed when done
        fmt.Println("Cache file created successfully")
    } else {
        fmt.Println("Cache file already exists")
    }
}
    }



var cache string // ~/.cache/gopherconf
var wallpaper string                 
var cachewp string  // cached wallpaper
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
    // with ".cache" and "gopherconf" to create a specific cache directory
    cache = es.Join(home, ".cache", "gopherconf")

    // Construct the cached wallpaper path
    cachewp = cache + "/" + "CACHED_WALLPAPER"

    // Check if the cache directory exists, if not, create it
    CacheDir()
}

// This function is used to set the wallpaper based on the session type
// It should be called after the wallpaper path has been set
func FuncUtil() {
    if SESSION == "wayland" {
        util = "swaybg"
        
        // Create a command to run swaybg with the following arguments:
        // * -i specifies the input image
        // * wallpaper is the path to the wallpaper image, it being a string variable fetched from flag "-w" 
        // * > /dev/null 2>&1 & runs the command in the background and redirects stdout and stderr to /dev/null
        es.ExecShell("nohup", true, "sh", "-c", util+" -i "+wallpaper+" > /dev/null 2>&1 &")
    } else if SESSION == "x11" {
        util = "feh"
        
        // Create a command to run feh with the following arguments:
        // * --bg-scale specifies the image is to be scaled to the desktop size
        // * wallpaper is the path to the wallpaper image
        // * > /dev/null 2>&1 & runs the command in the background and redirects stdout and stderr to /dev/null
        es.ExecShell("nohup", true, "sh", "-c", util+" --bg-scale "+wallpaper+" > /dev/null 2>&1 &")
    } else {
        fmt.Println("Unknown session type:", SESSION)
    }
}


