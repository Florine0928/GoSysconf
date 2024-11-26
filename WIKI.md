
# Wiki
- The Documentation
# Pywal Linker
So basically, after the Pywal cache is built, a CSS file is generated, which programs such as Waybar which are styled in CSS can @import and make use of the color table of Pywal. The linker just symlinks the generated file to ~/.config/waybar and after the user can @import the file in style.css and use the color table
# Window Manager Usage
You can use this program as a startup script for WM technically. Just normally configure Gopherconf by following the -help command, then include a exec in your WM config file, here is a example for Hyprland:
```conf
exec = gomgr -r all
```
What this does is call the program to -reload all components *that are enabled
# Config File
The config is written in JSON, its located at $HOME/.cache/gopherconf/config.json
# Looper Usage
This is basically a on/off toggle for components, you could bind a looper event to your keybind daemon to toggle on/off a desired component, if you don't get what I mean, here is my Hyprland config:
```bash
gomgr -l bar
gomgr -l wallpaper
# gomgr -l all
```

```conf
bind = $mainMod, D, exec, gomgr exec looper.wallpaper
# Will turn on/off the wallpaper setter every time you press Win + D

bind = $mainMod, B, exec, gomgr exec looper.bar
# Will turn on/off the bar every time you press Win + B
```
