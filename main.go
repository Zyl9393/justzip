package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/justzip/flags"
)

func main() {
	f := flags.Parse(os.Args[1:])
	basePath := filepath.Dir(f.CompressPath)
	if err := zipFileOrFolder(f.ArchivePath, f.CompressPath, basePath); err != nil {
		fmt.Println(err)
	}
}

func zipFileOrFolder(zipPath, filePath, basePath string) (err error) {
	var newZipFile *os.File
	newZipFile, err = os.Create(zipPath)
	if err != nil {
		return err
	}
	defer func() {
		newZipFile.Close()
		if err != nil {
			os.Remove(zipPath)
		}
	}()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	if err = addToZip(zipWriter, filePath, basePath); err != nil {
		return err
	}
	return nil
}

func addToZip(zipWriter *zip.Writer, filename, basePath string) error {
	stat, err := os.Stat(filename)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		infos, err := ioutil.ReadDir(filename)
		if err != nil {
			return err
		}
		for _, info := range infos {
			fname := filepath.Join(filename, info.Name())
			fmt.Printf("Adding %s to archive.\n", fname)
			err = addToZip(zipWriter, fname, basePath)
			if err != nil {
				return err
			}
		}
		return nil
	}

	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	header, err := zip.FileInfoHeader(stat)
	if err != nil {
		return err
	}
	header.Name, err = filepath.Rel(basePath, filename)
	if err != nil {
		return err
	}
	header.Name = strings.ReplaceAll(header.Name, `\`, `/`) // Linux can only deal with /, while Windows can deal with / and \.
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}
