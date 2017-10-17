package tvshow

import (
	"regexp"
	"fmt"
	"strings"
	"strconv"
	"errors"
)

var punctuationReplace = strings.NewReplacer(".", " ",
	"_", " ",
)

var articleReplace = strings.NewReplacer(" Of ", " of ",
	" The ", " the ",
	" On ", " on ",
	" In ", " in ",
	" And ", " and ",
	" Vs ", " vs ",
	" Del ", " del ",
	" El ", " el ",
	" La ", " la ",
	" En ", " en ",
)

type ParserFunc func(string, string) (*TvShow, error)

type rule struct {
	Function ParserFunc
	Regex    string
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
	rules := make([]rule, 6)
	rules[0] = rule{
		Regex:    `(^[0-9A-Za-z._ ]*)(^*[Ss][0-9]{2})(^*[Ee][0-9]{2})`,
		Function: threeGroups,
	}
	rules[1] = rule{
		Regex:    `(^[0-9A-Za-z_]*)(^*[.][0-9]{2})(^*[x][0-9]{2})`,
		Function: threeGroups,
	}
	rules[2] = rule{
		Regex:    `(^[0-9A-Za-z_ ]*)(^*- )(^*[0-9]{1})(^*x)(^*[0-9]{2})`,
		Function: fiveGroups,
	}
	rules[3] = rule{
		Regex:    `(^[0-9A-Za-z ]*)(^*Temporada [0-9]* )(Capitulo [0-9]*$)`,
		Function: threeGroupsFullWords,
	}
	rules[4] = rule{
		Regex:    `(^[0-9A-Za-z]*)(^*720p_)(^*[0-9]{3})`,
		Function: threeGroupsCamelCaseQuality,
	}
	rules[5] = rule{
		Regex:    `(^[0-9A-Za-z]*)(^*_)(^*[0-9]{3})`,
		Function: threeGroupsCamelCase,
	}
	return &Parser{Rules: rules}
}

func threeGroups(filename, regex string) (*TvShow, error) {
	r := regexp.MustCompile(regex)
	findGroup := r.FindStringSubmatch(filename)
	if len(findGroup) < 4 {
		return nil, errors.New("not a match")
	}

	rawName := findGroup[1]

	escapedName := punctuationReplace.Replace(rawName)
	name := articleReplace.Replace(strings.Title(strings.TrimSpace(escapedName)))

	season := findGroup[2]
	seasonNumber, err := strconv.Atoi(strings.Trim(season[1:], " "));
	if err != nil {
		return nil, fmt.Errorf("unable to parse season number from %s", filename)
	}

	episode := findGroup[3]
	episodeNumber, err := strconv.Atoi(episode[1:]);
	if err != nil {
		return nil, fmt.Errorf("unable to parse episode number from %s", filename)
	}

	return &TvShow{
		Name:    name,
		Season:  seasonNumber,
		Episode: episodeNumber}, nil
}

func threeGroupsCamelCaseQuality(filename, regex string) (*TvShow, error) {
	r := regexp.MustCompile(regex)
	findGroup := r.FindStringSubmatch(filename)
	if len(findGroup) < 3 {
		return nil, errors.New("not a match")
	}

	rawName := findGroup[1]
	escapedName := punctuationReplace.Replace(blanks(rawName))
	name := articleReplace.Replace(strings.Title(strings.TrimSpace(escapedName)))

	season := findGroup[3][:1]
	seasonNumber, err := strconv.Atoi(strings.Trim(season, " "));
	if err != nil {
		return nil, fmt.Errorf("unable to parse season number from %s", filename)
	}

	episode := findGroup[3][0:]
	episodeNumber, err := strconv.Atoi(episode[1:]);
	if err != nil {
		return nil, fmt.Errorf("unable to parse episode number from %s", filename)
	}

	return &TvShow{
		Name:    name,
		Season:  seasonNumber,
		Episode: episodeNumber}, nil
}

func threeGroupsCamelCase(filename, regex string) (*TvShow, error) {
	r := regexp.MustCompile(regex)
	findGroup := r.FindStringSubmatch(filename)
	if len(findGroup) < 3 {
		return nil, errors.New("not a match")
	}

	rawName := findGroup[1]
	escapedName := punctuationReplace.Replace(blanks(rawName))
	name := articleReplace.Replace(strings.Title(strings.TrimSpace(escapedName)))

	season := findGroup[3][:1]
	seasonNumber, err := strconv.Atoi(strings.Trim(season, " "))
	if err != nil {
		return nil, fmt.Errorf("unable to parse season number from %s", filename)
	}

	episode := findGroup[3][0:]
	episodeNumber, err := strconv.Atoi(episode[1:])
	if err != nil {
		return nil, fmt.Errorf("unable to parse episode number from %s", filename)
	}

	return &TvShow{
		Name:    name,
		Season:  seasonNumber,
		Episode: episodeNumber}, nil
}

func fiveGroups(filename, regex string) (*TvShow, error) {
	r := regexp.MustCompile(regex)
	findGroup := r.FindStringSubmatch(filename)
	if len(findGroup) < 6 {
		return nil, errors.New("not a match")
	}

	rawName := findGroup[1]
	escapedName := punctuationReplace.Replace(rawName)
	name := articleReplace.Replace(strings.Title(strings.TrimSpace(escapedName)))

	season := findGroup[3]
	seasonNumber, err := strconv.Atoi(season);
	if err != nil {
		return nil, fmt.Errorf("unable to parse season number from %s", filename)
	}

	episode := findGroup[5]
	episodeNumber, err := strconv.Atoi(episode);
	if err != nil {
		return nil, fmt.Errorf("unable to parse episode number from %s", filename)
	}

	return &TvShow{
		Name:    name,
		Season:  seasonNumber,
		Episode: episodeNumber}, nil
}

func threeGroupsFullWords(filename, regex string) (*TvShow, error) {
	r := regexp.MustCompile(regex)
	findGroup := r.FindStringSubmatch(filename)

	if len(findGroup) < 4 {
		return nil, errors.New("not a match")
	}

	rawName := findGroup[1]

	escapedName := punctuationReplace.Replace(rawName)
	name := articleReplace.Replace(strings.Title(strings.TrimSpace(escapedName)))

	season := findGroup[2]
	seasonNumber, err := strconv.Atoi(strings.Trim(season[10:], " "));
	if err != nil {
		return nil, fmt.Errorf("unable to parse season number from %s", filename)
	}

	episode := findGroup[3]
	episodeNumber, err := strconv.Atoi(episode[9:]);
	if err != nil {
		return nil, fmt.Errorf("Unable to parse episode number from %s", filename)
	}

	return &TvShow{
		Name:    name,
		Season:  seasonNumber,
		Episode: episodeNumber}, nil
}

func (p *Parser) FromFilename(filename string) (*TvShow, error) {
	if len(filename) == 0 {
		return nil, errors.New("missing parameter filename")
	}

	for _, rule := range p.Rules {
		show, err := rule.Function(filename, rule.Regex)
		if err != nil {
			continue
		}
		return show, nil
	}
	return nil, fmt.Errorf("unable to parse filename %s", filename)
}
