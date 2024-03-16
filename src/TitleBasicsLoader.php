<?php

namespace Imdb;

use Exception;

class TitleBasicsLoader extends Loader
{
    public function __construct(
        string $filename,
        callable $filterCallback = null
    ) {
        parent::__construct($filename, $filterCallback);

        while (($line = gzgets($this->file)) !== false) {
            $fields = explode("\t", trim($line, "\n"));
            $tconst = $fields[0];
            $titleType = $fields[1];
            $primaryTitle = $fields[2];
            $originalTitle = $fields[3];
            $isAdult = ($fields[4] === '1');
            $startYear = ($fields[5] !== '\\N') ? intval($fields[5]) : null;
            $endYear = ($fields[6] !== '\\N') ? intval($fields[6]) : null;
            $runtimeMinutes = ($fields[7] !== '\\N') ? intval($fields[7]) : null;
            $genres = ($fields[8] !== '\\N') ? explode(',', $fields[8]) : [];

            if (isset($this->data[$tconst])) {
                throw new Exception("Duplicate title ID: $tconst");
            }

            $row = [
                'titleType' => $titleType,
                'primaryTitle' => $primaryTitle,
                'originalTitle' => $originalTitle,
                'isAdult' => $isAdult,
                'startYear' => $startYear,
                'endYear' => $endYear,
                'runtimeMinutes' => $runtimeMinutes,
                'genres' => $genres
            ];

            if ($filterCallback === null || $filterCallback($row)) {
                $this->data[$tconst] = $row;
            }
        }
    }
}