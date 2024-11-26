package main

import (
	"encoding/json"
	"fmt"
	es "github.com/furiousman59/GO59" // es.ExecShell es.StringRNG es.FileExists es.Join es.Read
	"os"
)

// Bash Freestyle
// var echo = fmt.Println
// echo("GNU Bash 4.4.19(1)-release (x86_64-pc-linux-gnu)")

func Usage(which string) {
	switch which {
	case "general":
		fmt.Println("Usage:", os.Args[0], "[OPTIONS]")
		fmt.Println("	Options:")
		fmt.Println("	-help, -h (Option)			Shows this help message unless a option is specified")
		fmt.Println("	-wallpaper, -w <path>			Wallpaper Path")
		fmt.Println("	-bar, -b 				Which Bar you use <waybar, polybar, ags>")
		fmt.Println("	-looper, -l				Cycle between enabling/disabling a component")
		fmt.Println("	-scheme, -m                             Colorscheme, pyd/pyl for pywal dark/light")
		fmt.Println("	-reload, -r				Refreshes Components such as bar")
		fmt.Println("	-disable, -d				Disable a component")
		fmt.Println("  exec								Execute a event")
	case "wallpaper":
		fmt.Println("Usage: ", os.Args[0], "-wallpaper, -w path/to/wallpaper.anyformat")
		fmt.Println("Swaybg or Feh must be installed")
	case "bar":
		fmt.Println("Usage: ", os.Args[0], "-bar, -b <waybar/polybar/ags>")
		fmt.Println("One of these must be installed")
	case "scheme":
		fmt.Println("Usage: ", os.Args[0], "-scheme, -m <manual/pyd/pyl>")
		fmt.Println("	Pywal must be installed: https://github.com/dylanaraps/pywal")
		fmt.Println("	Pyd and pyl are shortcuts for dark and light respectively")
	case "reload":
		fmt.Println("Usage: ", os.Args[0], "-reload, -r <all/util/bar>")
		fmt.Println("	Can only reload enabled components")
	case "exec":
		fmt.Println("Usage: ", os.Args[0], "exec <which>")
		fmt.Println("Available events: looper.wallpaper, looper.bar, looper.bar")
	case "disable":
		fmt.Println("Usage: ", os.Args[0], "-disable, -d <all/util/bar>")
	case "looper":
		fmt.Println("Usage: ", os.Args[0], "-looper, -l <all/util/bar>")
		fmt.Println("	Cycle between turning on/off a component")
		fmt.Println("	Available components: wallpaper, bar")
	default:
		fmt.Println("Usage:", os.Args[0], "[OPTIONS]")
	}
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

var home string
var cache string   // ~/.cache/gopherconf
var cachewp string // cached wallpaper
var SESSION = os.Getenv("XDG_SESSION_TYPE")
var util string

func FuncUtil() {
	if SESSION == "wayland" {
		KillUtil("util")
		es.ExecShell("nohup", true, "sh", "-c", util+" -i "+config.Wallpaper+" > /dev/null 2>&1 &")
	} else if SESSION == "x11" {
		KillUtil("util")
		es.ExecShell("nohup", true, "sh", "-c", util+" --bg-scale "+config.Wallpaper+" > /dev/null 2>&1 &")
	} else {
		fmt.Println("Unknown session type:", SESSION)
	}
}

func FuncBar() {
	if config.Bar == "disabled" || config.Bar == "" {
		fmt.Println("Error: No bar specified in configuration")
		return
	}

	if SESSION == "wayland" {
		if config.Bar == "waybar" {
			es.ExecShell("nohup", true, "sh", "-c", "waybar > /dev/null 2>&1 &")
		} else if config.Bar == "ags" {
			es.ExecShell("nohup", true, "sh", "-c", "ags run > /dev/null 2>&1 &")
		} else {
			fmt.Println("Error: Unknown bar:", config.Bar)
			return
		}
	} else if SESSION == "x11" {
		if config.Bar == "polybar" {
			es.ExecShell("nohup", true, "sh", "-c", "polybar mybar > /dev/null 2>&1 &")
		} else {
			fmt.Println("Error: Unknown bar:", config.Bar)
			return
		}
	} else {
		fmt.Println("Error: Unknown session type:", SESSION)
		return
	}
}
func KillUtil(which string) {
	if which == "all" {
		es.ExecShell("pkill", true, util)
		es.ExecShell("pkill", true, config.Bar)
	} else if which == "util" {
		es.ExecShell("pkill", true, util)
	} else if which == "bar" {
		es.ExecShell("pkill", true, config.Bar)
	}
}

func Reload(which string) {
	if which == "all" {
		KillUtil("all")
		FuncUtil()
		FuncBar()
	} else if which == "util" {
		KillUtil("util")
		FuncUtil()
	} else if which == "bar" {
		KillUtil("bar")
		FuncBar()
	}
}

func InitPywal() {
	if config.Scheme == "manual" {
		return
	}

	if config.Wallpaper == "" {
		fmt.Println("Error: No wallpaper specified in configuration")
		return
	}

	var command string
	switch config.Scheme {
	case "pyd":
		command = "wal -i " + config.Wallpaper
	case "pyl":
		command = "wal -l -i " + config.Wallpaper
	default:
		return
	}

	es.ExecShell("sh", true, "-c", command)
}

func Garbage(which string) {
	if which == "pywal" {
		es.ExecShell("rm", true, "-rf", es.Join(home, ".cache", "wal"))
	}
}

func Linker() {
	pywalPath := es.Join(home, ".cache", "wal", "colors-waybar.css")
	if config.Bar == "waybar" {
		waybarPath := es.Join(home, ".config", "waybar", "colors-waybar.css")
		if err := os.Symlink(pywalPath, waybarPath); err != nil {
			return
		}
	}
} // more to be added i guess??? you can @import the css file in waybar style.css now and make use of pywal!!!!!!!!!!!!!!!!11!1!1
// ps: i'm going insane

func Looper(which string) {
	if which == "util" {
		if config.LooperUtil == "enabled" && config.LooperUtilCycle == "first" {
			Reload("util")
			config.LooperUtilCycle = "last"
			WriteConfig()
		} else if config.LooperUtil == "enabled" && config.LooperUtilCycle == "last" {
			KillUtil("util")
			config.LooperUtilCycle = "first"
			WriteConfig()
		}
	} else if which == "bar" {
		// Ensures that the cycle is toggled between "first" and "last"
		if config.LooperUtile == "enabled" && config.LooperUtileCycle == "first" {
			Reload("bar")
			config.LooperUtileCycle = "last"
			WriteConfig()
		} else if config.LooperUtile == "enabled" && config.LooperUtileCycle == "last" {
			KillUtil("bar")
			config.LooperUtileCycle = "first"
			WriteConfig()
		}
	} else if which == "all" {
		Looper("util")
		Looper("bar")
	} else if config.Looper == "disabled" {
		return
	}
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
