
# Wiki
- The Documentation
# Pywal Linker
So basically, after the Pywal cache is built, a CSS file is generated, which programs such as Waybar which are styled in CSS can @import and make use of the color table of Pywal. The linker just symlinks the generated file to ~/.config/waybar and after the user can @import the file in style.css and use the color table
# Window Manager Usage
You can use this program as a startup script for WM technically. Just normally configure Gopherconf by following the -help command, then include a exec in your WM config file, here is a example for Hyprland:
```conf
exec = gopherconf -r all
```
What this does is call the program to -reload all components *that are enabled
# Config File
The config is written in JSON, its located at $HOME/.cache/gopherconf/config.json
