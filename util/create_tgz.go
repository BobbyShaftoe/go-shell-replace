package util

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type TarGzArgs struct {
	Src    string
	DstDir string
}

func (t *TarGzArgs) CreateArchive() error {
	src, dstdir := t.Src, t.DstDir

	fmt.Println("Creating tar file")
	tarFile, err := tarIt(src)
	if err != nil {
		return err
	}
	fmt.Println("Creating gzip file")
	err = gzipIt(tarFile, dstdir)
	if err != nil {
		return err
	}
	err = os.Remove(tarFile)
	return err
}

// Create Gzip file
func gzipIt(source, target string) error {
	reader, err := os.Open(source)
	if err != nil {
		return err
	}

	fileName := filepath.Base(source)
	target = fmt.Sprintf("%s.gz", fileName)
	writer, err := os.Create(target)
	if err != nil {
		return err
	}
	defer writer.Close()

	archiver := gzip.NewWriter(writer)
	archiver.Name = fileName
	defer archiver.Close()

	_, err = io.Copy(archiver, reader)
	return err
}

// Create Tar file
func tarIt(source string) (string, error) {
	fileName := filepath.Base(source)
	fileName = fmt.Sprintf("%s.tar", fileName)
	fmt.Println("Opening file for writing")
	tarFile, err := os.Create(fileName)
	if err != nil {
		return fileName, err
	}
	defer tarFile.Close()

	fmt.Println("About to write to file")
	tarBall := tar.NewWriter(tarFile)
	defer tarBall.Close()

	info, err := os.Stat(source)
	if err != nil {
		return fileName, err
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	err = filepath.Walk(source,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			header, err := tar.FileInfoHeader(info, info.Name())
			if err != nil {
				return err
			}

			if baseDir != "" {
				header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
			}

			if err := tarBall.WriteHeader(header); err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(tarBall, file)
			return err
		})
	return fileName, err
}
