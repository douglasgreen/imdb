package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/douglasgreen/imdb/internal/imdb"
)

func main() {
	titleLoader, err := imdb.NewTitleBasicsLoader(
		"data/title.basics.tsv.gz",
		nil,
		func(row imdb.TitleBasics) imdb.TitleBasics {
			return imdb.TitleBasics{
				TitleID:        row.TitleID,
				TitleType:      row.TitleType,
				StartYear:      row.StartYear,
				RuntimeMinutes: row.RuntimeMinutes,
				Genres:         row.Genres,
			}
		},
	)
	if err != nil {
		exitErr(err)
	}
	titles := titleLoader.Data

	ratingLoader, err := imdb.NewTitleRatingsLoader("data/title.ratings.tsv.gz", nil, nil)
	if err != nil {
		exitErr(err)
	}
	ratings := ratingLoader.Data

	counts := map[string]int{}
	rates := map[string]float64{}
	votes := map[string]int{}

	for titleID, rating := range ratings {
		title, ok := titles[titleID]
		if !ok {
			continue
		}
		t := title.TitleType
		counts[t]++
		rates[t] += rating.AverageRating
		votes[t] += rating.NumVotes
	}

	type pair struct {
		Type  string
		Count int
	}
	var list []pair
	for k, v := range counts {
		list = append(list, pair{Type: k, Count: v})
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Count > list[j].Count
	})

	fmt.Println("### Type Counts and Ratings")
	fmt.Println()
	fmt.Println("| Title Type | Count | Average Rating | Average Number of Votes |")
	fmt.Println("|------------|-------|----------------|-------------------------|")

	for _, p := range list {
		avgRate := rates[p.Type] / float64(p.Count)
		avgVotes := float64(votes[p.Type]) / float64(p.Count)
		fmt.Printf("| %s | %d | %.2f | %.0f |\n", p.Type, p.Count, avgRate, avgVotes)
	}

	fmt.Println()

	genreCounts := map[string]int{}
	for _, title := range titles {
		if len(title.Genres) == 0 {
			continue
		}
		for _, g := range title.Genres {
			genreCounts[g]++
		}
	}

	type gpair struct {
		Genre string
		Count int
	}
	var glist []gpair
	for k, v := range genreCounts {
		glist = append(glist, gpair{Genre: k, Count: v})
	}
	sort.Slice(glist, func(i, j int) bool {
		return glist[i].Count > glist[j].Count
	})

	fmt.Println("### Genre Counts")
	fmt.Println()
	fmt.Println("| Genre | Count |")
	fmt.Println("|-------|-------|")
	for _, p := range glist {
		fmt.Printf("| %s | %d |\n", p.Genre, p.Count)
	}

	fmt.Println()

	runtimeCounts := map[string]int{}
	for _, title := range titles {
		if title.TitleType != "movie" {
			continue
		}
		if title.RuntimeMinutes == nil || *title.RuntimeMinutes <= 0 {
			continue
		}
		r := *title.RuntimeMinutes
		rounded := roundToNearestTen(r)
		key := fmt.Sprintf("%d", rounded)
		if rounded >= 300 {
			key = "300+"
		}
		runtimeCounts[key]++
	}

	type rpair struct {
		Key   string
		Value int
	}
	var rlist []rpair
	for k, v := range runtimeCounts {
		rlist = append(rlist, rpair{Key: k, Value: v})
	}
	sort.Slice(rlist, func(i, j int) bool {
		return runtimeKeyLess(rlist[i].Key, rlist[j].Key)
	})

	fmt.Println("### Movie Runtimes")
	fmt.Println()
	fmt.Println("| Runtime (minutes) | Count |")
	fmt.Println("|-------------------|-------|")
	for _, p := range rlist {
		fmt.Printf("| %s | %d |\n", p.Key, p.Value)
	}
	fmt.Println()
}

func roundToNearestTen(v int) int {
	return int((v + 5) / 10 * 10)
}

func runtimeKeyLess(a, b string) bool {
	if a == "300+" {
		return false
	}
	if b == "300+" {
		return true
	}
	return a < b
}

func exitErr(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}
