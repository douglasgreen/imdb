#!/usr/bin/env bash

# Change to the directory of the script
cd "$(dirname "$0")"

echo "Fetching name.basics.tsv"
wget https://datasets.imdbws.com/name.basics.tsv.gz

echo "Fetching title.ratings.tsv"
wget https://datasets.imdbws.com/title.ratings.tsv.gz

# echo "Fetching title.akas.tsv"
# wget https://datasets.imdbws.com/title.akas.tsv.gz
#
# echo "Fetching title.basics.tsv"
# wget https://datasets.imdbws.com/title.basics.tsv.gz
#
# echo "Fetching title.crew.tsv"
# wget https://datasets.imdbws.com/title.crew.tsv.gz
#
# echo "Fetching title.episode.tsv"
# wget https://datasets.imdbws.com/title.episode.tsv.gz
#
# echo "Fetching title.principals.tsv"
# wget https://datasets.imdbws.com/title.principals.tsv.gz

mv *.tsv.gz data
