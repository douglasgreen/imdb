<?php

declare(strict_types=1);

namespace DouglasGreen\Imdb;

use DouglasGreen\Utility\Data\DataException;
use DouglasGreen\Utility\Data\ValueException;

class TitleEpisodeLoader extends Loader
{
    public const HEADERS = ['tconst', 'parentTconst', 'seasonNumber', 'episodeNumber'];

    /**
     * @throws DataException
     * @throws ValueException
     */
    public function __construct(
        string $filename,
        callable $filterCallback = null,
        callable $processRow = null,
    ) {
        parent::__construct($filename);

        $line = gzgets($this->file);
        if ($line === false) {
            throw new DataException('Header not found: ' . $filename);
        }

        $fields = explode("\t", trim($line, PHP_EOL));
        if ($fields !== self::HEADERS) {
            throw new DataException('Format not recognized: ' . $filename);
        }

        while (($line = gzgets($this->file)) !== false) {
            $fields = explode("\t", trim($line, PHP_EOL));
            $episodeId = $fields[0];
            $parentId = $fields[1];
            $seasonNumber = $fields[2] !== '\N' ? intval($fields[2]) : null;
            $episodeNumber = $fields[3] !== '\N' ? intval($fields[3]) : null;

            if (isset($this->data[$episodeId])) {
                throw new ValueException('Duplicate episode ID: ' . $episodeId);
            }

            $row = [
                'episodeId' => $episodeId,
                'parentId' => $parentId,
                'seasonNumber' => $seasonNumber,
                'episodeNumber' => $episodeNumber,
            ];

            if ($filterCallback === null || $filterCallback($row)) {
                if ($processRow !== null) {
                    $row = $processRow($row);
                }

                $this->data[$episodeId] = $row;
            }
        }
    }

    /**
     * @return array<string, mixed>
     */
    public function getEpisode(string $episodeId): ?array
    {
        return $this->data[$episodeId] ?? null;
    }

    /**
     * @return array<string, array<string, mixed>>
     */
    public function getEpisodesByParentId(string $parentId): ?array
    {
        $episodes = [];

        foreach ($this->data as $episodeId => $row) {
            if ($row['parentId'] === $parentId) {
                $episodes[$episodeId] = $row;
            }
        }

        return $episodes;
    }
}
