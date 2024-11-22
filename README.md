
# Gopherconf
    // MacOS is not supported
    unix := [3]string{"Linux", "FreeBSD", "AnyBSD"}
    fmt.Printf("A %v CLI Utility\n", unix)
## What this?

A CLI system management utility for Unix-like systems written in glorious language of the Gopher

## What this does?
- Set a wallpaper, optionally cache it
- Start/Kill/Reload a bar
- Refresh Pywal cache on each wallpaper set
- Seamless change between dark and light mode
- Check the wiki for more info

## Use case?
Probally a backend utility you hook up to your window manager's keybind daemon


## Installation
```bash
Dependencies: golang make
```

```bash
git clone github.com/furiousman59/gopherconf && cd gopherconf && make build && sudo make install
```

## Usage

```bash
gopherconf -h
```

### Thanks readme.so for this readme!!
