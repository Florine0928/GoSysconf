package main

import (
	"fmt"
	"os"
	"path/filepath"
    "encoding/json"
)

var config struct { // config
    Wallpaper string `json:"wallpaper"`
    Util string `json:"util"`
}

func main() {
    configPath := filepath.Join(cache, "config.json")
    b, err := os.ReadFile(configPath)
    if err != nil {
        config = struct {
            Wallpaper string `json:"wallpaper"`
            Util string `json:"util"`
        }{}
    } else {
        err = json.Unmarshal(b, &config)
        if err != nil {
            fmt.Println("Error: Could not unmarshal configuration:", err)
            return
        }
    }

    args := os.Args[1:]
    for i := 0; i < len(args); i++ {
        arg := args[i]
        switch arg {
        case "-help", "-h":
            Usage()
        case "-wallpaper", "-w":
            if i+1 < len(args) {
                config.Wallpaper = args[i+1]
                if config.Util == "disabled" {
                    config.Util = "enabled"
                }
                WriteConfig()
                FuncUtil()
                fmt.Println("Wallpaper Path set to:", config.Wallpaper)
                i++
            } else if config.Wallpaper != "" {
                FuncUtil()
                fmt.Println("Cached Wallpaper:", config.Wallpaper)
            } else {
                fmt.Println("Error: No wallpaper path provided")
                return
            }
        case "-disable", "-d":
            if i+1 < len(args) {
                if args[i+1] == "wallpaper" {
                    config.Util = "disabled"
                    WriteConfig()
                    KillBar()
                    fmt.Println(util, "has been disabled")
                    i++
                }
            }
        default:
            fmt.Printf("Unknown operation: %s\n", arg)
        }
    }
}

