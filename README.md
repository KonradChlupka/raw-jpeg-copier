# raw-jpeg-copier

Raw-JPEG-copier can be used to ~~copy~~ move raw images that match the names of
equivalent JPEG images.

Typical use case:

1. Take 2000 photos in `JPEG + RAW` mode,
2. Use a method of your choice to select 40 good shots, copy/move the identified
   JPEGs to a new directory,
3. Run main.go to move the matching 40 raw images to the same directory, rather
   than finding them manually.

Raw-JPEG-copier moves your raw photos to a new location. It's in alpha so be
careful and make backups!

The raw and JPEG file extensions can be selected so this will work with
Fujifilm, Sony, Canon, Nikon, or any other camera brand.

Usage:

```
main:
  -jpeg-dir string
        directory with the jpeg files which you want to find the pairing raw for
  -jpeg-ext string
        case-sensitive extension for the JPEG files, incl. the leading period (default ".JPG")
  -output-dir string
        the directory to which the raw files should be moved
  -raw-ext string
        case-sensitive extension for the raw files, incl. the leading period (default ".RAF")
  -source-dir string
        source directory with the raw files
```

Example:

```
go run main.go -source-dir "/Users/konrad.chlupka/Pictures/Capture One Catalog.cocatalog/Originals/2022/02/22/1" -jpeg-dir "/Users/konrad.chlupka/Pictures/tennis" -output-dir "/Users/konrad.chlupka/Pictures/tennis"
```
