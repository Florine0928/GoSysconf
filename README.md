
# Gopherconf
    // MacOS is not supported
    unix := [3]string{"Linux", "FreeBSD", "AnyBSD"}
    fmt.Printf("A %v CLI Utility\n", unix)
## What this?

A CLI system management utility for Unix-like systems written in glorious language of the Gopher

## What this does?
- Set wallpaper (Swaybg and Feh for Wayland and X11)
- Start, Kill, Reload Waybar and AGS Bar on the fly
- Optional Pywal integration, and linker to Waybar (Check wiki.md for more details)
- Gopher
- Extra info in Wiki.md

## Use case?
Probally a backend utility you hook up to your window manager's keybind daemon


## Installation
```bash
Dependencies: golang make
Go Dependencies: GO59
```

```bash
go get github.com/furiousman59/GO59@Latest
go install github.com/furiousman59/gopherconf@latest
```

## Usage

```bash
gopherconf -help
```

### Thanks readme.so for this readme!!
