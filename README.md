# tvshow-parse
Parse tv show filenames

# How to Install

## Setup repository

```
sudo apt-get install apt-transport-https
```

For amd64:
```
wget -O - http://adelolmo.github.io/andoni.delolmo@gmail.com.gpg.key | sudo apt-key add -
echo "deb http://adelolmo.github.io xenial main" | sudo tee /etc/apt/sources.list.d/adelolmo.list
sudo apt-get update
```
For arm:
```
wget -O - http://adelolmo.github.io/andoni.delolmo@gmail.com.gpg.key | sudo apt-key add -
echo "deb http://adelolmo.github.io jessie main" | sudo tee /etc/apt/sources.list.d/adelolmo.list
sudo apt-get update
```

## Install package
```
sudo apt-get install tvshow-parse
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