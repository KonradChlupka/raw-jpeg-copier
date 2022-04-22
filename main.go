package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// main moves RAW files which match the names of the JPEG files.
func main() {
	jpegExt := flag.String("jpeg-ext", ".JPG", "case-sensitive extension for the JPEG files, incl. the leading period")
	rawExt := flag.String("raw-ext", ".RAF", "case-sensitive extension for the raw files, incl. the leading period")

	sourceDirWithRawFiles := flag.String("source-dir", "", "source directory with the raw files")
	dirWithJPEGs := flag.String("jpeg-dir", "", "directory with the jpeg files which you want to find the pairing raw for")
	rawOutputDir := flag.String("output-dir", "", "the directory to which the raw files should be moved")

	flag.Parse()
	if *sourceDirWithRawFiles == "" || *dirWithJPEGs == "" {
		flag.Usage()
	}

	if *rawOutputDir == "" {
		rawOutputDir = dirWithJPEGs
	}

	if err := os.MkdirAll(*rawOutputDir, 0755); err != nil {
		log.Fatalln("Failed to create the output directory:", err)
	}

	if err := filepath.WalkDir(*dirWithJPEGs, func(path string, _ fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		fileInfo, err := os.Stat(path)
		if err != nil {
			log.Println("Failed to get file stats:", err)
			return err
		}

		if fileInfo.IsDir() {
			log.Println("Skipping, found a directory:", path)
			return nil
		}

		if filepath.Ext(path) != *jpegExt {
			log.Println("Skipping, unknown file extension in file:", path)
			return nil
		}

		_, jpegFilename := filepath.Split(path)
		rawFilename := fmt.Sprintf("%s%s", strings.TrimSuffix(jpegFilename, filepath.Ext(path)), *rawExt)

		rawFileInSourceDir := filepath.Join(*sourceDirWithRawFiles, rawFilename)
		rawFileInDestDir := filepath.Join(*rawOutputDir, rawFilename)

		if err = os.Rename(rawFileInSourceDir, rawFileInDestDir); err != nil {
			log.Println("Failed to move the file:", err)
			return err
		}
		log.Printf("Finished moving file from %s to %s", rawFileInSourceDir, rawFileInDestDir)

		return nil
	}); err != nil {
		log.Println(err)
	}
}
