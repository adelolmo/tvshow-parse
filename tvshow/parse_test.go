package tvshow_test

import (
	"github.com/adelolmo/tvshow-parse/tvshow"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoArgument(t *testing.T) {
	parser := tvshow.NewParser()
	_, err := parser.FromFilename("")
	assert.Equal(t, "missing parameter filename", err.Error())
}

func TestOneWordShow(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("Westworld.S12E11.PROPER.720p.HDTV.x264-BATV.mkv")
	assert.Equal(t, "Westworld", show.Name)
	assert.Equal(t, 12, show.Season)
	assert.Equal(t, 11, show.Episode)
}

func TestLowerCaseShow(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("westworld.S02E01.PROPER.720p.HDTV.x264-BATV.mkv")
	assert.Equal(t, "Westworld", show.Name)
	assert.Equal(t, 2, show.Season)
	assert.Equal(t, 1, show.Episode)
}

func TestArticleLowerCase(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("the.man.in.the.high.castle.s03e10.720p.web.h264-skgtv.mkv")
	assert.Equal(t, "The Man in the High Castle", show.Name)
	assert.Equal(t, 3, show.Season)
	assert.Equal(t, 10, show.Episode)
}

func TestMultipleWordsShow(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("The.Expanse.S02E01.Dulcinea.1080p.WEB-DL.DD5.1.H.264-VietHD.mkv")
	assert.Equal(t, "The Expanse", show.Name)
	assert.Equal(t, 2, show.Season)
	assert.Equal(t, 1, show.Episode)
}

func TestMultipleWordsWithUpperCaseShow(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("DCs.Legends.of.Tomorrow.S02E01.720p.HDTV.x264-SVA.mkv")
	assert.Equal(t, "DCs Legends of Tomorrow", show.Name)
	assert.Equal(t, 2, show.Season)
	assert.Equal(t, 1, show.Episode)
}

func TestMultipleWordsLowerCaseShow(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("the.expanse.S02E01.Dulcinea.1080p.WEB-DL.DD5.1.H.264-VietHD.mkv")
	assert.Equal(t, "The Expanse", show.Name)
	assert.Equal(t, 2, show.Season)
	assert.Equal(t, 1, show.Episode)
}

func TestMultipleWordsUnderscoreShow(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("doctor_who_2005.12x11.720p_hdtv_x264-fov.mkv")
	assert.Equal(t, "Doctor Who 2005", show.Name)
	assert.Equal(t, 12, show.Season)
	assert.Equal(t, 11, show.Episode)
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

func TestDontUpperCaseMiddleArticlesOrPrepositions(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("Game.of.Thrones.S02E01.720p.HDTV.x264-AVS.mkv")
	assert.Equal(t, "Game of Thrones", show.Name)
	assert.Equal(t, 2, show.Season)
	assert.Equal(t, 1, show.Episode)
}

func TestOneDigitSeasonDashSeparator(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("el ministerio del tiempo - 2x11")
	assert.Equal(t, "El Ministerio del Tiempo", show.Name)
	assert.Equal(t, 2, show.Season)
	assert.Equal(t, 11, show.Episode)
}

func TestOneDigitSeason(t *testing.T) {
	parser := tvshow.NewParser()
	show, err := parser.FromFilename("el   otro lado 1x3")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "El Otro Lado", show.Name)
	assert.Equal(t, 1, show.Season)
	assert.Equal(t, 3, show.Episode)
}

func TestOneDigitSeasonWithFileExtension(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("El ministerio del tiempo - 2x11 (EliteTorrent.net).mp4")
	assert.Equal(t, "El Ministerio del Tiempo", show.Name)
	assert.Equal(t, 2, show.Season)
	assert.Equal(t, 11, show.Episode)
}

func TestDoubleSeasonDigitsDoubleEpisodeDigitsWithBlanks(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("Doctor Who 2005 S10E01 720p HDTV x264 FoV")
	assert.Equal(t, "Doctor Who 2005", show.Name)
	assert.Equal(t, 10, show.Season)
	assert.Equal(t, 1, show.Episode)
}

func TestSpanishSeasonOneDigitFullWords(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("El Ministerio Del Tiempo Temporada 2 Capitulo 10")
	assert.Equal(t, "El Ministerio del Tiempo", show.Name)
	assert.Equal(t, 2, show.Season)
}

func TestSpanishSeasonTwoDigitsFullWords(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("El Ministerio Del Tiempo Temporada 12 Capitulo 10")
	assert.Equal(t, "El Ministerio del Tiempo", show.Name)
	assert.Equal(t, 12, show.Season)
}

func TestSpanishEpisodeOneDigitFullWords(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("El Ministerio Del Tiempo Temporada 2 Capitulo 1")
	assert.Equal(t, "El Ministerio del Tiempo", show.Name)
	assert.Equal(t, 1, show.Episode)
}

func TestSpanishEpisodeTwoDigitsFullWords(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("El Ministerio Del Tiempo Temporada 2 Capitulo 10")
	assert.Equal(t, "El Ministerio del Tiempo", show.Name)
	assert.Equal(t, 10, show.Episode)
}

func TestNoSpaceInFilenameNorSeasonEpisodeWithVideoQuality(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("ElMinisterioDelTiempo720p_201_WWW.NEWPCT1.COM.mkv")
	assert.Equal(t, "El Ministerio del Tiempo", show.Name)
	assert.Equal(t, 2, show.Season)
	assert.Equal(t, 1, show.Episode)
}

func TestQualityWithoutPBeforeEpisodeAndSeason(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("Two Words 720 2x11 [www.url.com].mkv")
	assert.Equal(t, "Two Words", show.Name)
	assert.Equal(t, 2, show.Season)
	assert.Equal(t, 11, show.Episode)
}

func TestQualityWithPBeforeEpisodeAndSeason(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("Two Words 720p 2x11 [www.url.com].mkv")
	assert.Equal(t, "Two Words", show.Name)
	assert.Equal(t, 2, show.Season)
	assert.Equal(t, 11, show.Episode)
}

func TestNoSpaceInFilenameNorSeasonEpisodeWithX(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("Title720p2x11 [www.url.com].mkv")
	assert.Equal(t, "Title", show.Name)
	assert.Equal(t, 2, show.Season)
	assert.Equal(t, 11, show.Episode)
}

func TestNoSpaceInFilenameNorSeasonEpisode(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("ElMinisterioDelTiempo_201_WWW.NEWPCT1.COM.mkv")
	assert.Equal(t, "El Ministerio del Tiempo", show.Name)
	assert.Equal(t, 2, show.Season)
	assert.Equal(t, 1, show.Episode)
}

func TestNoSpaceInFilenameNorSeasonEpisodeTwoDigitsSeason(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("ElMinisterioDelTiempo_211_WWW.NEWPCT1.COM.mkv")
	assert.Equal(t, "El Ministerio del Tiempo", show.Name)
	assert.Equal(t, 2, show.Season)
	assert.Equal(t, 11, show.Episode)
}

func TestTitleWithDash(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("the.x-files.s02e11.720p.web.x264-tbs.mkv")
	assert.Equal(t, "The X-Files", show.Name)
	assert.Equal(t, 2, show.Season)
	assert.Equal(t, 11, show.Episode)
}

func TestTvShowSeasonEpisode(t *testing.T) {
	parser := tvshow.NewParser()
	show, _ := parser.FromFilename("vota juan 2x16")
	assert.Equal(t, "Vota Juan", show.Name)
	assert.Equal(t, 2, show.Season)
	assert.Equal(t, 16, show.Episode)
}

func TestFilm(t *testing.T) {
	parser := tvshow.NewParser()
	_, err := parser.FromFilename("Logan.2017.1080p.WEB-DL.H264.AC3-EVO[EtHD].mkv")
	assert.Equal(t, "unable to parse filename 'Logan.2017.1080p.WEB-DL.H264.AC3-EVO[EtHD].mkv'", err.Error())
}
