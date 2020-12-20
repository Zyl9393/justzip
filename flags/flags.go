package flags

import (
	"flag"
	"path/filepath"
)

type Flags struct {
	CompressPath string
	ArchivePath  string
}

func Parse(args []string) *Flags {
	f := &Flags{}
	fs := flag.NewFlagSet("app", flag.PanicOnError)
	fs.StringVar(&f.CompressPath, "i", "", "Path to file/folder to compress")
	fs.StringVar(&f.ArchivePath, "o", "", "Where to write the zip")
	fs.Parse(args)
	f.CompressPath = filepath.FromSlash(f.CompressPath)
	f.ArchivePath = filepath.FromSlash(f.ArchivePath)
	if f.CompressPath == "" {
		panic("must specify path for -i")
	}
	if f.ArchivePath == "" {
		panic("must specify path for -o")
	}
	return f
}
