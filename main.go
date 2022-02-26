package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// main copies RAW files which match the name of the JPEG files.
func main() {
	jpegExt := flag.String("jpeg-ext", ".JPG", "case-sensitive extension for the JPEG files, incl. the leading period")
	rawExt := flag.String("raw-ext", ".RAF", "case-sensitive extension for the raw files, incl. the leading period")

	sourceDirWithRawFiles := flag.String("source-dir", "", "source directory with the raw files")
	dirWithJPEGs := flag.String("jpeg-dir", "", "directory with the jpeg files which you want to find the pairing raw for")
	rawOutputDir := flag.String("output-dir", "", "the directory to which the raw files should be copied")

	flag.Parse()
	if *sourceDirWithRawFiles == "" || *dirWithJPEGs == "" || *rawOutputDir == "" {
		flag.Usage()
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
			return nil
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
		sourceFd, err := os.Open(rawFileInSourceDir)
		if err != nil {
			log.Println("Failed to open source raw file:", err)
			log.Println("The matching raw file might be missing")
			return nil
		}
		defer sourceFd.Close()

		rawFileInDestDir := filepath.Join(*rawOutputDir, rawFilename)
		destFd, err := os.Create(rawFileInDestDir)
		if err != nil {
			log.Println("Failed to open destination raw file:", err)
			return err
		}
		defer destFd.Close()

		bytesWritten, err := io.Copy(destFd, sourceFd)
		if err != nil {
			log.Println("Failed to copy file:", err)
			return nil
		}
		log.Printf("Finished copying file from %s to %s\n %d bytes written", rawFileInSourceDir, rawFileInDestDir, bytesWritten)

		return nil
	}); err != nil {
		log.Println(err)
	}
}
