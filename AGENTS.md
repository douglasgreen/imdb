# IMDb Dataset Processor

Go project that processes IMDb non-commercial datasets (title.basics.tsv.gz, title.ratings.tsv.gz).

## Commands

| Command | Description |
|---------|-------------|
| `go build ./cmd/filter` | Build filter binary |
| `go build ./cmd/printstats` | Build stats binary |
| `./fetch.sh` | Download data files to `data/` (run first) |
| `./filter --title-type movie --min-year 2010` | Filter movies by criteria |
| `./printstats` | Print dataset statistics |

## Filter Options

- `-y, --min-year`: Minimum release year (1900-current)
- `-t, --title-type`: Title type (movie, tvEpisode, tvSeries, short, etc.)
- `-g, --genre`: Genre filter (Drama, Comedy, Animation, etc.)
- `-r, --min-rating`: Minimum rating 0-10
- `-v, --min-votes`: Minimum vote count
- `-s, --sort-by-votes`: Sort by votes instead of rating
- `-a, --adult`: Include only adult titles

## Data Requirements

Run `./fetch.sh` first. Requires these files in `data/`:
- `title.basics.tsv.gz`
- `title.ratings.tsv.gz`

## Project Structure

- `cmd/filter/main.go` - Filter CLI entrypoint
- `cmd/printstats/main.go` - Statistics CLI entrypoint
- `internal/imdb/` - Core data loading and types
  - `loader.go` - Gzip TSV parsing
  - `title_basics.go` - Title basics schema
  - `title_ratings.go` - Ratings schema
  - `name_basics.go` - Names schema (unused)

## Notes

- Go 1.22 (module: github.com/douglasgreen/imdb)
- No tests, linting, or CI configured
- No external dependencies beyond stdlib
- Data files are large (~500MB+ each); not committed