<?php

namespace Imdb;

use Exception;

class TitlePrincipalsLoader extends Loader
{
    public function __construct(
        string $filename,
        callable $filterCallback = null
    ) {
        parent::__construct($filename, $filterCallback);

        while (($line = gzgets($this->file)) !== false) {
            $fields = explode("\t", trim($line, "\n"));
            $titleId = $fields[0];
            $ordering = intval($fields[1]);
            $personId = $fields[2];
            $category = $fields[3];
            $job = $fields[4] !== '\\N' ? $fields[4] : null;
            $characters = $fields[5] !== '\\N' ? $fields[5] : null;

            if (isset($this->data[$titleId][$ordering])) {
                throw new Exception("Duplicate title ID and order: $titleId, $ordering");
            }

            $row = [
                'personId' => $personId,
                'category' => $category,
                'job' => $job,
                'characters' => $characters
            ];

            if ($filterCallback === null || $filterCallback($row)) {
                $this->data[$titleId][$ordering] = $row;
            }
        }
    }

    public function getPrincipals(string $titleId): ?array
    {
        return $this->data[$titleId] ?? null;
    }

    public function getPrincipalsByPersonId(string $personId): ?array
    {
        $principals = [];

        foreach ($this->data as $titleId => $titlePrincipals) {
            foreach ($titlePrincipals as $ordering => $row) {
                if ($row['personId'] === $personId) {
                    $principals[$titleId][$ordering] = $row;
                }
            }
        }

        return $principals;
    }
}