# tvshow-parse
Parse tv show filenames

# How to Install

## Setup repository

Follow the instructions [here](https://adelolmo.github.io).

## Install package
```
# apt-get install tvshow-parse
```

# How to Use
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