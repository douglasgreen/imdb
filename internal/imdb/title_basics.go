package imdb

import (
	"fmt"
	"strconv"
	"strings"
)

type TitleBasics struct {
	TitleID        string
	TitleType      string
	PrimaryTitle   string
	OriginalTitle  string
	IsAdult        bool
	StartYear      *int
	EndYear        *int
	RuntimeMinutes *int
	Genres         []string
}

var ValidGenres = []string{
	"Action",
	"Adult",
	"Adventure",
	"Animation",
	"Biography",
	"Comedy",
	"Crime",
	"Documentary",
	"Drama",
	"Family",
	"Fantasy",
	"Film-Noir",
	"Game-Show",
	"History",
	"Horror",
	"Music",
	"Musical",
	"Mystery",
	"News",
	"Reality-TV",
	"Romance",
	"Sci-Fi",
	"Short",
	"Sport",
	"Talk-Show",
	"Thriller",
	"War",
	"Western",
}

var ValidTitleTypes = []string{
	"movie",
	"short",
	"tvEpisode",
	"tvMiniSeries",
	"tvMovie",
	"tvPilot",
	"tvSeries",
	"tvShort",
	"tvSpecial",
	"video",
	"videoGame",
}

type TitleBasicsLoader struct {
	Data map[string]TitleBasics
}

func NewTitleBasicsLoader(
	filename string,
	filter func(TitleBasics) bool,
	mapper func(TitleBasics) TitleBasics,
) (*TitleBasicsLoader, error) {
	r, err := openGzipFile(filename)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	header, err := r.readLine()
	if err != nil {
		return nil, fmt.Errorf("header not found: %w", err)
	}

	expected := []string{
		"tconst",
		"titleType",
		"primaryTitle",
		"originalTitle",
		"isAdult",
		"startYear",
		"endYear",
		"runtimeMinutes",
		"genres",
	}
	if !equalFields(strings.Split(header, "\t"), expected) {
		return nil, fmt.Errorf("format not recognized: %s", filename)
	}

	data := make(map[string]TitleBasics)

	for {
		line, err := r.readLine()
		if err != nil {
			if err.Error() == "EOF" || err == ioEOF() {
				break
			}
			return nil, fmt.Errorf("read line: %w", err)
		}
		fields := strings.Split(line, "\t")
		if len(fields) < 9 {
			continue
		}

		titleID := fields[0]
		if _, ok := data[titleID]; ok {
			return nil, fmt.Errorf("duplicate title ID: %s", titleID)
		}

		isAdult := fields[4] == "1"

		startYear, err := parseOptionalInt(fields[5])
		if err != nil {
			return nil, fmt.Errorf("parse startYear: %w", err)
		}
		endYear, err := parseOptionalInt(fields[6])
		if err != nil {
			return nil, fmt.Errorf("parse endYear: %w", err)
		}
		runtimeMinutes, err := parseOptionalInt(fields[7])
		if err != nil {
			return nil, fmt.Errorf("parse runtimeMinutes: %w", err)
		}

		genres := []string{}
		if fields[8] != `\N` {
			genres = strings.Split(fields[8], ",")
		}

		row := TitleBasics{
			TitleID:        titleID,
			TitleType:      fields[1],
			PrimaryTitle:   fields[2],
			OriginalTitle:  fields[3],
			IsAdult:        isAdult,
			StartYear:      startYear,
			EndYear:        endYear,
			RuntimeMinutes: runtimeMinutes,
			Genres:         genres,
		}

		if filter == nil || filter(row) {
			if mapper != nil {
				row = mapper(row)
			}
			data[titleID] = row
		}
	}

	return &TitleBasicsLoader{Data: data}, nil
}
