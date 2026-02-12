package imdb

import (
	"bufio"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type gzipReader struct {
	file *os.File
	gz   *gzip.Reader
	br   *bufio.Reader
}

func openGzipFile(filename string) (*gzipReader, error) {
	if _, err := os.Stat(filename); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("file does not exist: %s", filename)
		}
		return nil, fmt.Errorf("stat file: %w", err)
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}

	gz, err := gzip.NewReader(f)
	if err != nil {
		_ = f.Close()
		return nil, fmt.Errorf("open gzip: %w", err)
	}

	return &gzipReader{
		file: f,
		gz:   gz,
		br:   bufio.NewReader(gz),
	}, nil
}

func (r *gzipReader) Close() error {
	if r == nil {
		return nil
	}
	var err1 error
	if r.gz != nil {
		err1 = r.gz.Close()
	}
	err2 := r.file.Close()
	if err1 != nil {
		return err1
	}
	return err2
}

func (r *gzipReader) readLine() (string, error) {
	line, err := r.br.ReadString('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		return "", err
	}
	if errors.Is(err, io.EOF) && len(line) == 0 {
		return "", io.EOF
	}
	return strings.TrimRight(line, "\r\n"), nil
}
