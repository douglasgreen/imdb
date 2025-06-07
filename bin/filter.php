#!/usr/bin/env php
<?php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use DouglasGreen\Imdb\TitleBasicsLoader;
use DouglasGreen\Imdb\TitleRatingsLoader;
use DouglasGreen\OptParser\OptParser;
use DouglasGreen\Utility\Data\ArgumentException;

$optParser = new OptParser('IMDB Processor', 'Process IMDB non-commercial datasets');

$optParser
    ->addParam(['min-year', 'y'], 'INT', 'Minimum year', static function (int $value): int {
        if ($value < 1900 || $value > date('Y')) {
            throw new ArgumentException('Year not within range: ' . $value);
        }

        return $value;
    })
    ->addParam(['title-type', 't'], 'STRING', 'Title type', static function (
        string $value,
    ): string {
        if (! in_array($value, TitleBasicsLoader::VALID_TITLE_TYPES, true)) {
            throw new ArgumentException(
                'Title not valid; must be one of: ' .
                    implode(', ', TitleBasicsLoader::VALID_TITLE_TYPES),
            );
        }

        return $value;
    })
    ->addParam(['genre', 'g'], 'STRING', 'Genre', static function (string $value): string {
        if (! in_array($value, TitleBasicsLoader::VALID_GENRES, true)) {
            throw new ArgumentException(
                'Genre not valid; must be one of: ' .
                    implode(', ', TitleBasicsLoader::VALID_GENRES),
            );
        }

        return $value;
    })
    ->addParam(['min-rating', 'r'], 'FLOAT', 'Minimum rating', static function (
        float $value,
    ): float {
        if ($value < 0.0 || $value > 10.0) {
            throw new ArgumentException('Value not in range 0 to 10');
        }

        return $value;
    })
    ->addParam(['min-votes', 'v'], 'INT', 'Minimum votes', static function (int $value): int {
        if ($value < 1) {
            throw new ArgumentException('Value must be greater than 0');
        }

        return $value;
    })
    ->addFlag(['adult', 'a'], 'Include only adult films')
    ->addFlag(['sort-by-votes', 's'], 'Sort by votes')
    ->addUsageAll();

$input = $optParser->parse();

$minYear = $input->get('min-year');
$titleType = $input->get('title-type');
$genre = $input->get('genre');
$minRating = $input->get('min-rating');
$minVotes = $input->get('min-votes');
$adult = (bool) $input->get('adult');
$sortByVotes = (bool) $input->get('sort-by-votes');

$titleLoader = new TitleBasicsLoader(
    __DIR__ . '/../assets/data/title.basics.tsv.gz',
    static function ($row) use ($minYear, $titleType, $genre, $adult): bool {
        if ($minYear && $row['startYear'] < $minYear) {
            return false;
        }

        if ($titleType && $row['titleType'] !== $titleType) {
            return false;
        }

        if ($genre && is_array($row['genres']) && ! in_array($genre, $row['genres'], true)) {
            return false;
        }

        return ! ($adult && ! $row['isAdult']);
    },
    static fn($row): array => [
        'primaryTitle' => $row['primaryTitle'],
        'startYear' => $row['startYear'],
        'genres' => $row['genres'],
    ],
);
$titles = $titleLoader->getData();

$ratingLoader = new TitleRatingsLoader(
    __DIR__ . '/../assets/data/title.ratings.tsv.gz',
    static function (array $row) use ($titles, $minRating, $minVotes): bool {
        if (! isset($titles[$row['titleId']])) {
            return false;
        }

        if ($minRating && $row['averageRating'] < $minRating) {
            return false;
        }

        return ! ($minVotes && $row['numVotes'] < $minVotes);
    },
);
$ratings = $sortByVotes ? $ratingLoader->getTopVotedTitles() : $ratingLoader->getTopRatedTitles();

foreach ($ratings as $titleId => $rating) {
    $title = $titles[$titleId];
    $genres = $title['genres'];
    $primaryTitle = $title['primaryTitle'];
    $startYear = $title['startYear'];
    $averageRating = $rating['averageRating'];
    $numVotes = $rating['numVotes'];
    $genreDesc = $genres ? ' (' . implode(', ', $genres) . ')' : '';
    echo sprintf(
        '%s (%s): %s * %s%s%s',
        $primaryTitle,
        $startYear,
        $averageRating,
        $numVotes,
        $genreDesc,
        PHP_EOL,
    );
}
