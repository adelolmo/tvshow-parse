package tvshow

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const debug = false

var punctuationReplace = strings.NewReplacer(".", " ",
	"_", " ",
)

var replace = func(word string) string {
	switch word {
	case "Of", "The", "On", "In", "And", "Vs", "Del", "El", "La", "En":
		return strings.ToLower(word)

	case "of", "the", "on", "in", "and", "vs", "del", "el", "la", "en":
		return word
	}
	return strings.Title(word)
}

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
	rules := make([]rule, 10)
	rules[0] = rule{
		Regex:    `(^[0-9A-Za-z._\- ]*)(^*[Ss][0-9]{2})(^*[Ee][0-9]{2})`,
		Function: threeGroups,
	}
	rules[1] = rule{
		Regex:    `(^[0-9A-Za-z_\- ]*)(^*[.][0-9]{2})(^*[x][0-9]{2})`,
		Function: threeGroups,
	}
	rules[2] = rule{
		Regex:    `(^[A-Za-z ]*)(^*[ 0-9]{2})(^*[x][0-9]{2})`,
		Function: threeGroups,
	}
	rules[3] = rule{
		Regex:    `(^[0-9A-Za-z_ ]*)(^*- )(^*[0-9]{1})(^*x)(^*[0-9]{2})`,
		Function: fiveGroups,
	}
	rules[4] = rule{
		Regex:    `(^[0-9A-Za-z ]*)(^* 720 )(^*[0-9]{1})(^*x)(^*[0-9]{2})`,
		Function: fiveGroups,
	}
	rules[5] = rule{
		Regex:    `(^[0-9A-Za-z ]*)(^* 720p )(^*[0-9]{1})(^*x)(^*[0-9]{2})`,
		Function: fiveGroups,
	}
	rules[6] = rule{
		Regex:    `(^[0-9A-Za-z]*)(^*720p)(^*[0-9]{1})(^*x)(^*[0-9]{2})`,
		Function: fiveGroups,
	}
	rules[7] = rule{
		Regex:    `(^[0-9A-Za-z ]*)(^*Temporada [0-9]* )(Capitulo [0-9]*$)`,
		Function: threeGroupsFullWords,
	}
	rules[8] = rule{
		Regex:    `(^[0-9A-Za-z]*)(^*720p_)(^*[0-9]{3})`,
		Function: threeGroupsCamelCaseQuality,
	}
	rules[9] = rule{
		Regex:    `(^[0-9A-Za-z]*)(^*_)(^*[0-9]{3})`,
		Function: threeGroupsCamelCase,
	}
	return &Parser{Rules: rules}
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
	return nil, fmt.Errorf("unable to parse filename '%s'", filename)
}

func threeGroups(filename, regex string) (*TvShow, error) {
	r := regexp.MustCompile(regex)
	findGroup := r.FindStringSubmatch(filename)
	if debug {
		_, _ = fmt.Printf("threeGroups len:%d  %s\n", len(findGroup), filename)
		for i := 0; i < len(findGroup); i++ {
			fmt.Printf("findGroup[%d]  %s\n", i, findGroup[i])
		}
	}
	if len(findGroup) < 4 {
		return nil, errors.New("not a match")
	}

	rawName := findGroup[1]
	escapedName := punctuationReplace.Replace(rawName)
	name := title(escapedName)
	if debug {
		fmt.Printf("rawName: %s\n", rawName)
		fmt.Printf("escapedName: %s\n", escapedName)
		fmt.Printf("name: %s\n", name)
	}

	season := findGroup[2]
	if debug {
		fmt.Printf("season: %s\n", season)
	}
	seasonNumber, err := strconv.Atoi(strings.Trim(season[1:], " "))
	if debug {
		fmt.Printf("seasonNumber: %d\n", seasonNumber)
	}
	if err != nil {
		return nil, fmt.Errorf("unable to parse season number from %s", filename)
	}

	episode := findGroup[3]
	if debug {
		fmt.Printf("episode: %s\n", episode)
	}
	episodeNumber, err := strconv.Atoi(episode[1:])
	if debug {
		fmt.Printf("episodeNumber: %d\n", episodeNumber)
	}
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
	if debug {
		fmt.Printf("threeGroupsCamelCaseQuality len:%d  %s\n", len(findGroup), filename)
	}
	if len(findGroup) < 3 {
		return nil, errors.New("not a match")
	}

	rawName := findGroup[1]
	escapedName := punctuationReplace.Replace(blanks(rawName))
	name := title(escapedName)
	if debug {
		fmt.Printf("rawName: %s\n", rawName)
		fmt.Printf("escapedName: %s\n", escapedName)
		fmt.Printf("name: %s\n", name)
	}

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

func threeGroupsCamelCase(filename, regex string) (*TvShow, error) {
	r := regexp.MustCompile(regex)
	findGroup := r.FindStringSubmatch(filename)
	if debug {
		fmt.Printf("threeGroupsCamelCase len:%d  %s\n", len(findGroup), filename)
	}
	if len(findGroup) < 3 {
		return nil, errors.New("not a match")
	}

	rawName := findGroup[1]
	escapedName := punctuationReplace.Replace(blanks(rawName))
	name := title(escapedName)
	if debug {
		fmt.Printf("rawName: %s\n", rawName)
		fmt.Printf("escapedName: %s\n", escapedName)
		fmt.Printf("name: %s\n", name)
	}

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
	if debug {
		fmt.Printf("fiveGroups len:%d  %s\n", len(findGroup), filename)
	}
	if len(findGroup) < 6 {
		return nil, errors.New("not a match")
	}

	rawName := findGroup[1]
	escapedName := punctuationReplace.Replace(rawName)
	name := title(escapedName)
	if debug {
		fmt.Printf("rawName: %s\n", rawName)
		fmt.Printf("escapedName: %s\n", escapedName)
		fmt.Printf("name: %s\n", name)
	}

	season := findGroup[3]
	seasonNumber, err := strconv.Atoi(season)
	if err != nil {
		return nil, fmt.Errorf("unable to parse season number from %s", filename)
	}

	episode := findGroup[5]
	episodeNumber, err := strconv.Atoi(episode)
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

	if debug {
		fmt.Printf("threeGroupsFullWords len:%d  %s\n", len(findGroup), filename)
	}
	if len(findGroup) < 4 {
		return nil, errors.New("not a match")
	}

	rawName := findGroup[1]
	escapedName := punctuationReplace.Replace(rawName)
	name := title(escapedName)
	if debug {
		fmt.Printf("rawName: %s\n", rawName)
		fmt.Printf("escapedName: %s\n", escapedName)
		fmt.Printf("name: %s\n", name)
	}

	season := findGroup[2]
	seasonNumber, err := strconv.Atoi(strings.Trim(season[10:], " "))
	if err != nil {
		return nil, fmt.Errorf("unable to parse season number from %s", filename)
	}

	episode := findGroup[3]
	episodeNumber, err := strconv.Atoi(episode[9:])
	if err != nil {
		return nil, fmt.Errorf("unable to parse episode number from %s", filename)
	}

	return &TvShow{
		Name:    name,
		Season:  seasonNumber,
		Episode: episodeNumber}, nil
}

func title(name string) string {
	rx := regexp.MustCompile(`\w+`)
	title := rx.ReplaceAllStringFunc(name, replace)
	return strings.Title(title[0:1]) + strings.TrimSpace(title[1:])
}
