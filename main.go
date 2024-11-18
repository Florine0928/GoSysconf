package main

import (
	"fmt"
	"os"
	"os/exec"
	"log"
	"strings"
)

func main() {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "-help", "-h":
			Usage()
//		case "-bar", "-b":
//			if i+1 < len(args) {
//				
//			}
		case "-wallpaper", "-w":
			if i+1 < len(args) {
				cmd := exec.Command("realpath", args[i+1])
				output, err := cmd.Output()
				if err != nil {
					log.Fatal(err)
				}
				wallpaper = strings.TrimSpace(string(output))
				
				cpCmd := exec.Command("cp", wallpaper, cache)
				err = cpCmd.Run() 
				i++ 
				fmt.Println("Wallpaper Path set to:", wallpaper)
				FuncUtil()
			} else {
				fmt.Println("ERROR: No path defined")
				return
			}
		default:
			fmt.Printf("Unknown operation: %s\n", arg)
		}
	}
}
