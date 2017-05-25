package main

import (
	"github.com/adelolmo/tvshow/tvshow"
	"fmt"
	"flag"
	"os"
)

func main() {
	filename := flag.String("filename", "", "tv show's filename")
	filter := flag.String("filter", "name", "tv show's filter. (name|season|episode)")

	flag.Parse()
	if len(*filename) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	parser := tvshow.NewParser();
	show, err := parser.FromFilename(*filename);
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	switch *filter {
	case "name":
		fmt.Print(show.Name)
	case "season":
		fmt.Print(show.Season)
	case "episode":
		fmt.Print(show.Episode)
	}
}
