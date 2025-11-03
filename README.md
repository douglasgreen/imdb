# imdb

Process [IMDb non-commercial datasets](https://developer.imdb.com/non-commercial-datasets/).

## Project setup

Standard config files for linting and testing are copied into place from a GitHub repository called
[utility](https://github.com/douglasgreen/utility). See that project's README page for details.

## Scripts

-   `bin/fetch.sh` - Fetch data files and put in `data/` directory.
-   `bin/filter.php` - Filter the data for recent highly-rated movies.
-   `bin/print_stats.php` - Print statistics about titles and ratings.

### Using filter

This PHP script, `filter`, allows you to filter the IMDB data sets and find recent highly rated
movies based on various criteria. It provides a convenient way to explore the IMDB data and discover
popular movies that match your preferences.

#### Usage

To use the script, run it from the command line with the desired options:

```
./filter.php [options]
```

#### Options

The script supports the following options:

-   `-y, --min-year`: Specify the minimum year for the movies. Only movies released on or after this
    year will be included in the results.
-   `-t, --title-type`: Filter movies by title type. Examples of valid title types include
    "tvEpisode" for TV episodes and "movie" for feature films.
-   `-g, --genre`: Filter movies by genre. Examples of valid genres include "Documentary" and
    "Animation".
-   `-r, --min-rating`: Specify the minimum average rating for the movies. Only movies with an
    average rating equal to or higher than this value will be included in the results.
-   `-v, --min-votes`: Specify the minimum number of votes required for a movie to be considered.
    Only movies with a number of votes equal to or greater than this value will be included in the
    results.
-   `-s, --sort-by-votes`: Sort by votes instead of ratings.
-   `-a, --adult`: Include only adult films.

#### Examples

Here are a few examples of how to use the script with different options:

1. Find movies released in 2010 or later:

    ```
    ./filter.php --title-type movie --min-year 2010
    ```

2. Find TV episodes in the "Documentary" genre:

    ```
    ./filter.php --title-type tvEpisode --genre Documentary
    ```

3. Find movies with an average rating of 8.0 or higher and at least 10,000 votes:

    ```
    ./filter.php --title-type movie --min-rating 8.0 --min-votes 10000
    ```

4. Find animated movies released in 2015 or later with a minimum rating of 7.5:
    ```
    ./filter.php --title-type movie --min-year 2015 --genre Animation --min-rating 7.5
    ```

The script will output the filtered results, displaying the movie title, release year, average
rating, and number of votes for each matching movie.

Note: Make sure you have the required IMDB data sets (`title.basics.tsv.gz` and
`title.ratings.tsv.gz`) in the `data/` directory before running the script.

## Statistics

Here is a summary of data from the files downloaded on March 16, 2024.

### Type Counts and Ratings

This table shows that:

-   There are a variety of title types.
-   Movies get more votes but lower ratings than TV series.
-   There are 3.3 times as many movies as TV series.
-   There are 7.5 episodes in the average TV series.

| Title Type   | Count  | Average Rating | Average Number of Votes |
| ------------ | ------ | -------------- | ----------------------- |
| tvEpisode    | 704337 | 7.39           | 202                     |
| movie        | 306297 | 6.17           | 3635                    |
| short        | 156668 | 6.83           | 75                      |
| tvSeries     | 94056  | 6.86           | 1528                    |
| video        | 52702  | 6.58           | 200                     |
| tvMovie      | 52484  | 6.6            | 254                     |
| tvMiniSeries | 16824  | 7.12           | 1225                    |
| videoGame    | 16043  | 6.79           | 362                     |
| tvSpecial    | 12183  | 6.73           | 227                     |
| tvShort      | 2285   | 6.81           | 161                     |

### Genre Counts

This table shows that drama and comedy are the most popular genres.

| Genre       | Count   |
| ----------- | ------- |
| Drama       | 3021790 |
| Comedy      | 2102836 |
| Talk-Show   | 1303184 |
| Short       | 1157867 |
| Documentary | 1013042 |
| Romance     | 989789  |
| News        | 948336  |
| Family      | 787216  |
| Reality-TV  | 599165  |
| Animation   | 535258  |
| Crime       | 443956  |
| Action      | 437587  |
| Adventure   | 416566  |
| Music       | 400091  |
| Game-Show   | 387921  |
| Adult       | 333805  |
| Sport       | 253875  |
| Fantasy     | 216141  |
| Mystery     | 215164  |
| Horror      | 190618  |
| Thriller    | 175920  |
| History     | 157445  |
| Biography   | 114825  |
| Sci-Fi      | 114411  |
| Musical     | 90101   |
| War         | 41946   |
| Western     | 30442   |
| Film-Noir   | 883     |

### Movie Runtimes

This table shows that 90 minutes is the most common movie runtime.

| Runtime (minutes) | Count  |
| ----------------- | ------ |
| 0                 | 65     |
| 10                | 344    |
| 20                | 328    |
| 30                | 312    |
| 40                | 529    |
| 50                | 33685  |
| 60                | 39440  |
| 70                | 36983  |
| 80                | 62412  |
| 90                | 100979 |
| 100               | 63510  |
| 110               | 32837  |
| 120               | 20164  |
| 130               | 10792  |
| 140               | 8334   |
| 150               | 5370   |
| 160               | 3061   |
| 170               | 1639   |
| 180               | 1156   |
| 190               | 423    |
| 200               | 342    |
| 210               | 246    |
| 220               | 185    |
| 230               | 127    |
| 240               | 184    |
| 250               | 95     |
| 260               | 78     |
| 270               | 62     |
| 280               | 48     |
| 290               | 43     |
| 300+              | 559    |
