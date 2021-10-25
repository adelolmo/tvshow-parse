# tvshow-parse
Parse tv show filenames

## Build
The build process supports cross-platform compilation.
To set the target architecture add the argument `ARCH` with the value when calling `make`.
Supported architectures:
* amd64
* i386
* armhf
* arm64


    make ARCH=armhf

## Install

### Package
After building the package you can install it with `dpkg` command.

    sudo dpkg -i tvshow-parse_1.0.0_armhf.deb

### Binary only
You can also install the binary without building the package.

    sudo make install

## Run
Execute the program running the following command:

    tvshow-parse

There are three filters available: name, season and episode.

```
tvshow-parse -filename=The.Americans.2013.S01E09.720p.HDTV.X264-DIMENSION.mkv -filter=name
The Americans 2013
```
```
tvshow-parse -filename=The.Americans.2013.S01E09.720p.HDTV.X264-DIMENSION.mkv -filter=season
1
```
```
tvshow-parse -filename=The.Americans.2013.S01E09.720p.HDTV.X264-DIMENSION.mkv -filter=episode
9
```
