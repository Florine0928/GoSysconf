
**GOMGR CLI Utility**
=====================

A command-line utility for managing rice components on UNIX systems (excluding MacOS).

**Table of Contents**
-----------------

* [Introduction](#introduction)
* [Features](#features)
* [Installation](#installation)
* [Usage](#usage)
* [Configuration](#configuration)
* [Contributing](#contributing)
* [License](#license)

**Introduction**
---------------

### This utility provides a simple and easy-to-use interface for managing rice components, including bars and wallpaper setters on UNIX systems (excluding MacOS).

**Features**
------------

* Manage rice components, including bars and wallpaper setters
* Multiple colorschemes available: manual, pywal dark, pywal light
* Toggle on/off components such as bar and wallpaper setter
* Configure and reload components as needed on the fly
* Can be used as WM startup script

**Installation**
---------------

To install this utility, simply clone the repository and run the installation script:
```bash
go get github.com/furiousman59/GO59
git clone github.com/furiousman59/gomgr && cd gomgr
make build && sudo make install
```
**Usage**
-----

The utility can be run from the command line using the following syntax:
```bash
gomgr [command] [options]
```
List options:

* `-h`, `--help`: Display help message
* `-h {command}`, `--help {command}`: Display help message for a specific command

Further explanation of functions can be found in wiki.md

**Configuration**
--------------

The utility uses a configuration file to store settings. The configuration file is located at `~/.cache/gomgr/config.json`.

**Contributing**
------------

Contributions are welcome! If you'd like to contribute to this project, please fork the repository and submit a pull request.

**License**
-------

This project is licensed under the GNU General Public License V3. See the LICENSE file for details.
