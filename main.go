package main

import (
	"encoding/json"
	"fmt"
	es "github.com/furiousman59/GO59" // es.ExecShell es.StringRNG es.FileExists es.Join es.Read
	"os"
	"path/filepath"
)

var config struct { // config
	Wallpaper string `json:"wallpaper"` // path to wallpaper
	Util      string `json:"util"`      // enabled or disabled - related to wallpaper setter
	Bar       string `json:"bar"`       // self explanatory
	Utile     string `json:"utile"`     // enabled or disabled - related to bar
	Scheme    string `json:"scheme"`    // Colorscheme - semi-related to pywal
}

func main() {
	var err error
	home, err = os.UserHomeDir()
	if err != nil {
		fmt.Println("Error fetching home directory:", err)
		os.Exit(1)
	}
	cache = es.Join(home, ".cache", "gopherconf")
	CacheDir()

	if SESSION == "wayland" {
		util = "swaybg"
	} else if SESSION == "x11" {
		util = "feh"
	}

	configPath := filepath.Join(cache, "config.json")
	b, err := os.ReadFile(configPath)
	if err != nil {
		config = struct {
			Wallpaper string `json:"wallpaper"`
			Util      string `json:"util"`
			Bar       string `json:"bar"`
			Utile     string `json:"utile"`
			Scheme    string `json:"scheme"`
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
			if i+1 < len(args) {
				switch args[i+1] {
				case "wallpaper", "w":
					Usage("wallpaper")
					i++
				case "bar", "b":
					Usage("bar")
					i++
				case "scheme", "m":
					Usage("scheme")
					i++
				case "reload", "r":
					Usage("reload")
					i++
				case "disable", "d":
					Usage("disable")
					i++
				default:
					Usage("general")
				}
			} else {
				Usage("general")
			}
		case "-eval", "-e":
			Linker()
			InitPywal()
			// used for testing functions generally, don't use this unless you know what you're doing
		case "-scheme", "-m":
			if i+1 < len(args) {
				if args[i+1] == "manual" {
					config.Scheme = "manual"
					WriteConfig()
					Garbage("pywal")
					Reload("all")
					fmt.Println("Colorscheme set to manual")
					i++
				} else if args[i+1] == "pyd" {
					config.Scheme = "pyd"
					WriteConfig()
					InitPywal()
					Linker()
					Reload("all")
					fmt.Println("Colorscheme set to pywal dark")
					i++
				} else if args[i+1] == "pyl" {
					config.Scheme = "pyl"
					WriteConfig()
					InitPywal()
					Linker()
					Reload("all")
					fmt.Println("Colorscheme set to pywal light")
					i++
				}
			}
		case "-wallpaper", "-w":
			if i+1 < len(args) {
				config.Wallpaper = args[i+1]
				if config.Util == "disabled" {
					config.Util = "enabled"
				}
				WriteConfig()
				InitPywal()
				FuncUtil()
				fmt.Println("Wallpaper Path set to:", config.Wallpaper)
				i++
			} else if config.Wallpaper != "" {
				FuncUtil()
				if config.Util == "disabled" {
					config.Util = "enabled"
				}
				WriteConfig()
				fmt.Println("Cached Wallpaper:", config.Wallpaper)
			} else {
				fmt.Println("Error: No wallpaper path provided")
				return
			}
		case "-bar", "-b":
			if i+1 < len(args) {
				if SESSION == "wayland" {
					if args[i+1] == "waybar" {
						config.Bar = "waybar"
						config.Utile = "enabled"
						WriteConfig()
						KillUtil("bar")
						FuncBar()
					} else if args[i+1] == "ags" {
						config.Bar = "ags"
						config.Utile = "enabled"
						WriteConfig()
						KillUtil("bar")
						FuncBar()
					}
				} else if SESSION == "x11" {
					if args[i+1] == "polybar" {
						config.Bar = "polybar"
						config.Utile = "enabled"
						WriteConfig()
						KillUtil("bar")
						FuncBar()
					}
				}
			} else {
				fmt.Println("Error: Invalid bar provided:", args[i+1])
				return
			}
			i++

		case "-disable", "-d":
			if i+1 < len(args) {
				if args[i+1] == "wallpaper" {
					if config.Util != "disabled" {
						KillUtil("util")
						fmt.Println(util, "has been disabled")
						config.Util = "disabled"
						WriteConfig()
					} else {
						fmt.Println("Wallpaper setter is already disabled")
					}
					i++
				} else if args[i+1] == "bar" {
					if config.Utile != "disabled" {
						KillUtil("bar")
						fmt.Println(config.Bar, "has been disabled")
						config.Utile = "disabled"
						WriteConfig()
					} else {
						fmt.Println("Bar is already disabled")
					}
					i++
				} else if args[i+1] == "all" {
					if config.Util != "disabled" && config.Utile != "disabled" {
						KillUtil("all")
						config.Util = "disabled"
						config.Utile = "disabled"
						WriteConfig()
						fmt.Println("All components have been disabled")
					} else {
						fmt.Println("All components are already disabled")
					}
					i++
				} else {
					fmt.Println("Error: Invalid component to disable:", args[i+1])
					return
				}
			}
		case "-reload", "-r":
			if i+1 < len(args) {
				if args[i+1] == "wallpaper" {
					if config.Util != "disabled" {
						KillUtil("util")
						FuncUtil()
						fmt.Println("Wallpaper setter has been reloaded")
						i++
					}
				} else if args[i+1] == "bar" {
					if config.Utile != "disabled" {
						KillUtil("bar")
						FuncBar()
						fmt.Println("Bar has been reloaded")
						i++
					}
				} else if args[i+1] == "all" {
					if config.Util != "disabled" && config.Utile != "disabled" {
						KillUtil("all")
						FuncUtil()
						FuncBar()
						fmt.Println("All components have been reloaded")
						i++
					}
				} else if config.Util == "disabled" && config.Utile == "disabled" {
					fmt.Println("Error: all components are disabled")
					return
				} else {
					fmt.Println("Error: Invalid component to reload:", args[i+1])
					return
				}
			}
		default:
			Usage("general")
		}
	}
}
