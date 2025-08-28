package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/lftk/anki"
)

func main() {
	var path, output string
	switch args := os.Args[1:]; len(args) {
	case 1:
		path = args[0]
		output = strings.TrimSuffix(
			filepath.Base(path), filepath.Ext(path),
		)
	case 2:
		path = args[0]
		output = args[1]
	default:
		fmt.Println("usage: anki-unpkg <path> [<dir>]")
		os.Exit(1)
	}

	if err := unpack(path, output); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func unpack(path, output string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	dir := filepath.Join(wd, output)
	if _, err := os.Stat(dir); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
		// Directory does not exist, create it.
		if err = os.Mkdir(dir, 0755); err != nil {
			return err
		}
	} else {
		// Directory exists, check if it's empty.
		f, err := os.Open(dir)
		if err != nil {
			return err
		}
		defer f.Close()

		// Try to read just one entry. If we get EOF, it's empty.
		if _, err = f.ReadDir(1); err != io.EOF {
			if err == nil {
				// We successfully read one entry, so it's not empty.
				err = fmt.Errorf("%q already exists", output)
			}
			return err
		}
	}

	z, err := zip.OpenReader(path)
	if err != nil {
		return err
	}
	defer z.Close()

	return anki.Unpack(&z.Reader, dir)
}
