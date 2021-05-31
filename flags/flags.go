package flags

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Flags struct {
	CompressPath string
	ArchivePath  string
}

func Parse(args []string) *Flags {
	f := &Flags{}
	fs := flag.NewFlagSet("app", flag.ExitOnError)
	fs.StringVar(&f.CompressPath, "i", "", "Path to file/folder to compress")
	fs.StringVar(&f.ArchivePath, "o", "", "Where to write the zip")
	fs.Usage = func() {
		fmt.Println("Usage: justzip -i path/to/file/or/folder -o path/to/archive.zip")
	}
	fs.Parse(args)
	f.CompressPath = filepath.FromSlash(f.CompressPath)
	f.ArchivePath = filepath.FromSlash(f.ArchivePath)
	if f.CompressPath == "" || f.ArchivePath == "" {
		fs.Usage()
		os.Exit(2)
	}
	return f
}
