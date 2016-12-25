# id3inspect

This is a small command line tool for viewing [ID3](https://en.wikipedia.org/wiki/ID3) media information.

## Usage

Just give it a file name and it will print everything it can find about the file in the stdout

```sh
$ id3inspect 06.Some.People.mp3
Title:  Some People
Artist: Goldfrapp
Album:  Seventh Tree
Year:   2008
Track:  6
Length: 4m40s
Genre:  Electronic
Samplerate: 44100
Bitrate:    222 KB
```

## Installation

As any simple go program, you can just use `go get`:

```sh
go get github.com/ironsmile/id3inspect
```

and you are good to go.

## Why?

I haven't found any suitable way of viewing the meta information for a random MP3 files in OSX. In Finder, iTunes or otherwise. This tool displays everything I need. Users of Linux distros and Windows would probably find it useless as their GUI file managers show the same information easily.

## License

MIT
