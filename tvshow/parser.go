package tvshow

import (
	"regexp"
	"fmt"
	"strings"
	"strconv"
	"errors"
)

type Parser struct {

}

type TvShow struct {
	Name    string
	Season  int
	Episode int
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) FromFilename(filename string) (*TvShow, error) {
	if (len(filename) == 0) {
		return nil, errors.New("Missing parameter filename")
	}

	r := regexp.MustCompile(`(^[0-9A-Za-z.]*)(^*[Ss][0-9]{2})(^*[Ee][0-9]{2})`)
	findGroup := r.FindStringSubmatch(filename)
	if (len(findGroup) < 4) {
		return nil, fmt.Errorf("Unable to parse filename %s", filename)
	}

	rawName := findGroup[1]
	escapedName := strings.Replace(rawName, ".", " ", -1)
	name := strings.TrimSpace(escapedName)

	season := findGroup[2]
	seasonNumber, err := strconv.Atoi(season[1:]);
	if err != nil {
		return nil, fmt.Errorf("Unable to parse season number from %s", filename)
	}

	episode := findGroup[3]
	episodeNumber, err := strconv.Atoi(episode[1:]);
	if err != nil {
		return nil, fmt.Errorf("Unable to parse episode number from %s", filename)
	}

	return &TvShow{
		Name:name,
		Season:seasonNumber,
		Episode:episodeNumber}, nil
}