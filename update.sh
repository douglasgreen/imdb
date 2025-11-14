#!/bin/bash

./bin/filter.php --title-type tvEpisode --min-rating 7.39 --min-votes 2020 > var/tvEpisode
./bin/filter.php --title-type movie --min-rating 6.17 --min-votes 36350 > var/movie
./bin/filter.php --title-type short --min-rating 6.83 --min-votes 750 > var/short
./bin/filter.php --title-type tvSeries --min-rating 6.86 --min-votes 15280 > var/tvSeries
./bin/filter.php --title-type video --min-rating 6.58 --min-votes 2000 > var/video
./bin/filter.php --title-type tvMovie --min-rating 6.6 --min-votes 2540 > var/tvMovie
./bin/filter.php --title-type tvMiniSeries --min-rating 7.12 --min-votes 12250 > var/tvMiniSeries
./bin/filter.php --title-type videoGame --min-rating 6.79 --min-votes 3620 > var/videoGame
./bin/filter.php --title-type tvSpecial --min-rating 6.73 --min-votes 2270 > var/tvSpecial
./bin/filter.php --title-type tvShort --min-rating 6.81 --min-votes 1610 > var/tvShort
