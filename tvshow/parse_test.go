package tvshow_test

import (
	"testing"
	"github.com/bmizerany/assert"
	"github.com/adelolmo/tvshow/tvshow"
)

func TestNoArgument(t *testing.T) {
	parser := tvshow.NewParser()
	_, err := parser.FromFilename("")
	assert.Equal(t, "Missing parameter filename", err.Error())
}

func TestOneWordShow(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("Westworld.S02E01.PROPER.720p.HDTV.x264-BATV.mkv")
	assert.Equal(t, "Westworld", show.Name)
	assert.Equal(t, 2, show.Season)
	assert.Equal(t, 1, show.Episode)
}

func TestMultipleWordsShow(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("The.Expanse.S02E01.Dulcinea.1080p.WEB-DL.DD5.1.H.264-VietHD.mkv")
	assert.Equal(t, "The Expanse", show.Name)
	assert.Equal(t, 2, show.Season)
	assert.Equal(t, 1, show.Episode)
}

func TestShowWithNumbers(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("12.Monkeys.S02E01.720p.HDTV.x264-KILLERS.mkv")
	assert.Equal(t, "12 Monkeys", show.Name)
	assert.Equal(t, 2, show.Season)
	assert.Equal(t, 1, show.Episode)
}

func TestMultipleWordsWithNumbersShow(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("The.Americans.2013.S02E01.HDTV.x264-LOL.[VTV].mp4")
	assert.Equal(t, "The Americans 2013", show.Name)
	assert.Equal(t, 2, show.Season)
	assert.Equal(t, 1, show.Episode)
}

func TestFilm(t *testing.T) {
	parser := tvshow.NewParser()
	_, err := parser.FromFilename("Logan.2017.1080p.WEB-DL.H264.AC3-EVO[EtHD].mkv")
	assert.Equal(t, "Unable to parse filename Logan.2017.1080p.WEB-DL.H264.AC3-EVO[EtHD].mkv", err.Error())
}