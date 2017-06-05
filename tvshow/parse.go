package tvshow

import (
	"regexp"
	"fmt"
	"strings"
	"strconv"
	"errors"
)

type rule struct {
	Regex        string
	GroupSize    int
	NameGroup    int
	SeasonGroup  int
	EpisodeGroup int
}

type Parser struct {
	Rules []rule
}

type TvShow struct {
	Name    string
	Season  int
	Episode int
}

func NewParser() *Parser {
	rules := make([]rule, 2)
	rules[0] = rule{
		Regex:`(^[0-9A-Za-z.]*)(^*[Ss][0-9]{2})(^*[Ee][0-9]{2})`,
		GroupSize:4,
		NameGroup:1,
		SeasonGroup:2,
		EpisodeGroup:3,
	}
	rules[1] = rule{
		Regex:`(^[0-9A-Za-z_]*)(^*[.][0-9]{2})(^*[x][0-9]{2})`,
		GroupSize:4,
		NameGroup:1,
		SeasonGroup:2,
		EpisodeGroup:3,
	}

	return &Parser{Rules:rules}
}

func (p *Parser) FromFilename(filename string) (*TvShow, error) {
	if (len(filename) == 0) {
		return nil, errors.New("Missing parameter filename")
	}

	punctuationReplace := strings.NewReplacer(".", " ",
		"_", " ",
	)
	articleReplace := strings.NewReplacer(" Of ", " of ",
		" The ", " the ",
		" On ", " on ",
		" In ", " in ",
		" And ", " and ",
		" Vs ", " vs ",
	)
	for _, rule := range p.Rules {
		regex := rule.Regex
		r := regexp.MustCompile(regex)
		findGroup := r.FindStringSubmatch(filename)
		if (len(findGroup) < rule.GroupSize) {
			continue
		}

		rawName := findGroup[rule.NameGroup]

		escapedName := punctuationReplace.Replace(rawName)
		name := articleReplace.Replace(strings.Title(strings.TrimSpace(escapedName)))

		season := findGroup[rule.SeasonGroup]
		seasonNumber, err := strconv.Atoi(season[1:]);
		if err != nil {
			return nil, fmt.Errorf("Unable to parse season number from %s", filename)
		}

		episode := findGroup[rule.EpisodeGroup]
		episodeNumber, err := strconv.Atoi(episode[1:]);
		if err != nil {
			return nil, fmt.Errorf("Unable to parse episode number from %s", filename)
		}

		return &TvShow{
			Name:name,
			Season:seasonNumber,
			Episode:episodeNumber}, nil
	}
	return nil, fmt.Errorf("Unable to parse filename %s", filename)
}