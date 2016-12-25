// idv3inspect: tool for showing information about media files.
//
// It will just print the basic id3v1 and id3v2 tags for a file. This icludes
// Artist, Album, Title, Track Number, Year, Genre, Sample and Bitrate
//
// Run with --help or -h for more information.
//
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ironsmile/logger"
	taglib "github.com/wtolson/go-taglib"
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s MEDIA_FILE\n", os.Args[0])
	}
}

func main() {
	flag.Parse()

	logger.Default().Errorer = log.New(os.Stderr, "Error: ", 0)

	if len(flag.Args()) < 1 {
		logger.Fatalf("Missing argument: filename\nSee -h for more information\n")
	}

	fileName := flag.Arg(0)

	// Because I do not like the taglib's error message for "file not found"
	if _, err := os.Lstat(fileName); err != nil {
		logger.Fatalln(err)
	}

	// Open file and find tag in it
	tag, err := taglib.Read(fileName)
	if err != nil {
		logger.Fatal("Error while opening mp3 file: ", err)
	}
	defer tag.Close()

	fmt.Printf("Title:\t%s\n", tag.Title())
	fmt.Printf("Artist:\t%s\n", tag.Artist())
	fmt.Printf("Album:\t%s\n", tag.Album())
	fmt.Printf("Year:\t%d\n", tag.Year())
	fmt.Printf("Track:\t%d\n", tag.Track())
	fmt.Printf("Length:\t%s\n", tag.Length())
	fmt.Printf("Genre:\t%s\n", tag.Genre())
	fmt.Printf("Samplerate:\t%d\n", tag.Samplerate())
	fmt.Printf("Bitrate:\t%d KB\n", tag.Bitrate())
}
