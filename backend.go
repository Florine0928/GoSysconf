package main

import (
    "fmt"
    "os"
    "encoding/json"
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
    fmt.Println("	-disable, -d				Disable a component")
}

func CacheDir() {
    if es.FileExists(cache) == false {
        err := os.MkdirAll(cache, 0755)
        if err != nil {
        fmt.Println("Error creating cache directory:", err)
        os.Exit(1)
    }
    if !es.FileExists(cachewp) {
        file, err := os.Create(cachewp)
        if err != nil {
            fmt.Println("Error creating cache file:", err)
            return
        }
        defer file.Close()
        fmt.Println("Cache file created successfully")
    } else {
        fmt.Println("Cache file already exists")
    }
}
    }

var cache string // ~/.cache/gopherconf               
var cachewp string  // cached wallpaper
var SESSION = os.Getenv("XDG_SESSION_TYPE")
var util string

func init() {
    home, err := os.UserHomeDir()
    if err != nil {
        fmt.Println("Error fetching home directory:", err)
        os.Exit(1)
    }
    cache = es.Join(home, ".cache", "gopherconf")
    cachewp = cache + "/" + "CACHED_WALLPAPER"
    CacheDir()

    if SESSION == "wayland" {
        util = "swaybg"
    } else if SESSION == "x11" {
        util = "feh"
    }

}

func FuncUtil() {
    if SESSION == "wayland" {
        KillBar()
        es.ExecShell("nohup", true, "sh", "-c", util+" -i "+config.Wallpaper+" > /dev/null 2>&1 &")
    } else if SESSION == "x11" {
        KillBar()
        es.ExecShell("nohup", true, "sh", "-c", util+" --bg-scale "+config.Wallpaper+" > /dev/null 2>&1 &")
    } else {
        fmt.Println("Unknown session type:", SESSION)
    }
}

func KillBar() {
    es.ExecShell("pkill", true, util)
}

func WriteConfig() {
    configPath := es.Join(cache, "config.json")
    b, err := json.Marshal(config)
    if err != nil {
        fmt.Println("Error: Could not marshal configuration:", err)
        return
    }
    err = os.WriteFile(configPath, b, 0644)
    if err != nil {
        fmt.Println("Error: Could not write configuration:", err)
        return
    }
}   

