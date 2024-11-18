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
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error fetching home directory:", err)
		os.Exit(1)
	}
	cache = filepath.Join(home, ".cache", "GoSys")
}

func FuncUtil() {
    if SESSION == "wayland" {
        util = "swaybg"
        utilCmd := exec.Command("nohup", "sh", "-c", util+" -i "+wallpaper+" > /dev/null 2>&1 &")
        err := utilCmd.Run() 
        if err != nil {
            log.Fatal(err) 
        }
    } else if SESSION == "x11" {
        util = "feh"
        utilCmd := exec.Command("nohup", "sh", "-c", util+" --bg-scale "+wallpaper+" > /dev/null 2>&1 &")
        err := utilCmd.Run()  
        if err != nil {
            log.Fatal(err)  
        }
    } else {
        fmt.Println("Unknown session type:", SESSION)
    }
}


