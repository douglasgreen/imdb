package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/douglasgreen/imdb/internal/imdb"
)

func main() {
	var (
		minYear      int
		titleType    string
		genre        string
		minRating    float64
		minVotes     int
		adultOnly    bool
		sortByVotes  bool
	)

	flag.IntVar(&minYear, "min-year", 0, "Minimum year")
	flag.IntVar(&minYear, "y", 0, "Minimum year (shorthand)")
	flag.StringVar(&titleType, "title-type", "", "Title type")
	flag.StringVar(&titleType, "t", "", "Title type (shorthand)")
	flag.StringVar(&genre, "genre", "", "Genre")
	flag.StringVar(&genre, "g", "", "Genre (shorthand)")
	flag.Float64Var(&minRating, "min-rating", 0.0, "Minimum rating")
	flag.Float64Var(&minRating, "r", 0.0, "Minimum rating (shorthand)")
	flag.IntVar(&minVotes, "min-votes", 0, "Minimum votes")
	flag.IntVar(&minVotes, "v", 0, "Minimum votes (shorthand)")
	flag.BoolVar(&adultOnly, "adult", false, "Include only adult films")
	flag.BoolVar(&adultOnly, "a", false, "Include only adult films (shorthand)")
	flag.BoolVar(&sortByVotes, "sort-by-votes", false, "Sort by votes")
	flag.BoolVar(&sortByVotes, "s", false, "Sort by votes (shorthand)")
	flag.Parse()

	if minYear != 0 {
		year := time.Now().Year()
		if minYear < 1900 || minYear > year {
			exitErr(fmt.Errorf("year not within range: %d", minYear))
		}
	}

	if titleType != "" && !contains(imdb.ValidTitleTypes, titleType) {
		exitErr(fmt.Errorf("title not valid; must be one of: %s", join(imdb.ValidTitleTypes)))
	}

	if genre != "" && !contains(imdb.ValidGenres, genre) {
		exitErr(fmt.Errorf("genre not valid; must be one of: %s", join(imdb.ValidGenres)))
	}

	if minRating < 0.0 || minRating > 10.0 {
		exitErr(fmt.Errorf("value not in range 0 to 10"))
	}

	if minVotes < 0 {
		exitErr(fmt.Errorf("value must be greater than or equal to 0"))
	}

	titleLoader, err := imdb.NewTitleBasicsLoader(
		"data/title.basics.tsv.gz",
		func(row imdb.TitleBasics) bool {
			if minYear != 0 && (row.StartYear == nil || *row.StartYear < minYear) {
				return false
			}
			if titleType != "" && row.TitleType != titleType {
				return false
			}
			if genre != "" && !contains(row.Genres, genre) {
				return false
			}
			if adultOnly && !row.IsAdult {
				return false
			}
			return true
		},
		func(row imdb.TitleBasics) imdb.TitleBasics {
			return imdb.TitleBasics{
				TitleID:      row.TitleID,
				TitleType:    row.TitleType,
				PrimaryTitle: row.PrimaryTitle,
				StartYear:    row.StartYear,
				Genres:       row.Genres,
			}
		},
	)
	if err != nil {
		exitErr(err)
	}

	titles := titleLoader.Data

	ratingLoader, err := imdb.NewTitleRatingsLoader(
		"data/title.ratings.tsv.gz",
		func(row imdb.TitleRating) bool {
			if _, ok := titles[row.TitleID]; !ok {
				return false
			}
			if minRating != 0.0 && row.AverageRating < minRating {
				return false
			}
			if minVotes != 0 && row.NumVotes < minVotes {
				return false
			}
			return true
		},
		nil,
	)
	if err != nil {
		exitErr(err)
	}

	var ratings []imdb.TitleRating
	if sortByVotes {
		ratings = ratingLoader.TopVoted(0)
	} else {
		ratings = ratingLoader.TopRated(0)
	}

	for _, rating := range ratings {
		title := titles[rating.TitleID]
		year := ""
		if title.StartYear != nil {
			year = fmt.Sprintf("%d", *title.StartYear)
		}
		genreDesc := ""
		if len(title.Genres) > 0 {
			genreDesc = " (" + join(title.Genres) + ")"
		}
		fmt.Printf(
			"%s (%s, %s): %.1f * %d%s\n",
			title.PrimaryTitle,
			year,
			title.TitleType,
			rating.AverageRating,
			rating.NumVotes,
			genreDesc,
		)
	}
}

func exitErr(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}

func contains(list []string, v string) bool {
	for _, item := range list {
		if item == v {
			return true
		}
	}
	return false
}

func join(list []string) string {
	if len(list) == 0 {
		return ""
	}
	out := list[0]
	for i := 1; i < len(list); i++ {
		out += ", " + list[i]
	}
	return out
}
