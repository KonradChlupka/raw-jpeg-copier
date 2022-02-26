# raw-jpeg-copier

Raw-JPEG-copier can be used to copy raw images that match the names of
equivalent JPEG images.

Typical use case:

1. Take 2000 photos in `JPEG + RAW` mode,
2. Use a method of your choice to select 40 good shots, copy the identified
   JPEGs to a new directory,
3. Run main.go to copy the matching 40 raw images to the same directory, rather
   than finding them manually.

Raw-JPEG-copier doesn't delete any of your photos, only creates new copies of
the raw photos. But it's in alpha so be careful and make backups!

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
        the directory to which the raw files should be copied
  -raw-ext string
        case-sensitive extension for the raw files, incl. the leading period (default ".RAF")
  -source-dir string
        source directory with the raw files
```

Example:

```
go run main.go -source-dir "/Users/konrad.chlupka/Pictures/Capture One Catalog.cocatalog/Originals/2022/02/22/1" -jpeg-dir "/Users/konrad.chlupka/Pictures/tennis" -output-dir "/Users/konrad.chlupka/Pictures/tennis"
```
