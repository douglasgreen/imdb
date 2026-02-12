package imdb

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type TitleRating struct {
	TitleID       string
	AverageRating float64
	NumVotes      int
}

type TitleRatingsLoader struct {
	Data map[string]TitleRating
}

func NewTitleRatingsLoader(
	filename string,
	filter func(TitleRating) bool,
	mapper func(TitleRating) TitleRating,
) (*TitleRatingsLoader, error) {
	r, err := openGzipFile(filename)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	header, err := r.readLine()
	if err != nil {
		return nil, fmt.Errorf("header not found: %w", err)
	}

	expected := []string{"tconst", "averageRating", "numVotes"}
	if !equalFields(strings.Split(header, "\t"), expected) {
		return nil, fmt.Errorf("format not recognized: %s", filename)
	}

	data := make(map[string]TitleRating)

	for {
		line, err := r.readLine()
		if err != nil {
			if err.Error() == "EOF" || err == ioEOF() {
				break
			}
			return nil, fmt.Errorf("read line: %w", err)
		}
		fields := strings.Split(line, "\t")
		if len(fields) < 3 {
			continue
		}

		titleID := fields[0]
		if _, ok := data[titleID]; ok {
			return nil, fmt.Errorf("duplicate title ID: %s", titleID)
		}

		avg, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			return nil, fmt.Errorf("parse averageRating: %w", err)
		}

		numVotes, err := strconv.Atoi(fields[2])
		if err != nil {
			return nil, fmt.Errorf("parse numVotes: %w", err)
		}

		row := TitleRating{
			TitleID:       titleID,
			AverageRating: avg,
			NumVotes:      numVotes,
		}

		if filter == nil || filter(row) {
			if mapper != nil {
				row = mapper(row)
			}
			data[titleID] = row
		}
	}

	return &TitleRatingsLoader{Data: data}, nil
}

func (l *TitleRatingsLoader) TopRated(limit int) []TitleRating {
	items := make([]TitleRating, 0, len(l.Data))
	for _, v := range l.Data {
		items = append(items, v)
	}

	sort.Slice(items, func(i, j int) bool {
		ri := int(items[i].AverageRating * 1000)
		rj := int(items[j].AverageRating * 1000)
		if ri == rj {
			return items[i].NumVotes > items[j].NumVotes
		}
		return ri > rj
	})

	if limit > 0 && limit < len(items) {
		return items[:limit]
	}
	return items
}

func (l *TitleRatingsLoader) TopVoted(limit int) []TitleRating {
	items := make([]TitleRating, 0, len(l.Data))
	for _, v := range l.Data {
		items = append(items, v)
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].NumVotes == items[j].NumVotes {
			ri := int(items[i].AverageRating * 1000)
			rj := int(items[j].AverageRating * 1000)
			return ri > rj
		}
		return items[i].NumVotes > items[j].NumVotes
	})

	if limit > 0 && limit < len(items) {
		return items[:limit]
	}
	return items
}
