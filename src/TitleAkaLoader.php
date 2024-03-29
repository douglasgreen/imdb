<?php

namespace Imdb;

use Exception;

class TitleAkaLoader extends Loader
{
    const HEADERS = [
        'titleId',
        'ordering',
        'title',
        'region',
        'language',
        'types',
        'attributes',
        'isOriginalTitle'
    ];

    public function __construct(
        string $filename,
        callable $filterCallback = null,
        callable $processRow = null
    ) {
        parent::__construct($filename, $filterCallback);

        $line = gzgets($this->file);
        $fields = explode("\t", trim($line, "\n"));
        if ($fields != self::HEADERS) {
            throw new Exception("Format not recognized: $filename");
        }

        while (($line = gzgets($this->file)) !== false) {
            $fields = explode("\t", trim($line, "\n"));
            $titleId = $fields[0];
            $ordering = intval($fields[1]);
            $title = $fields[2];
            $region = ($fields[3] !== '\\N') ? $fields[3] : null;
            $language = ($fields[4] !== '\\N') ? $fields[4] : null;
            $types = ($fields[5] !== '\\N') ? explode(', ', $fields[5]) : null;
            $attributes = ($fields[6] !== '\\N') ? explode(', ', $fields[6]) : null;
            $isOriginalTitle = ($fields[7] === '1');

            if (isset($this->data[$titleId][$ordering])) {
                throw new Exception("Duplicate title ID and order: $titleId, $ordering");
            }

            $row = [
                'titleId' => $titleId,
                'ordering' => $ordering,
                'title' => $title,
                'region' => $region,
                'language' => $language,
                'types' => $types,
                'attributes' => $attributes,
                'isOriginalTitle' => $isOriginalTitle
            ];

            if ($filterCallback === null || $filterCallback($row)) {
                if ($processRow) {
                    $row = $processRow($row);
                }
                $this->data[$titleId][$ordering] = $row;
            }
        }
    }

    public function getAkas(string $titleId): ?array
    {
        return $this->data[$titleId] ?? null;
    }
}
