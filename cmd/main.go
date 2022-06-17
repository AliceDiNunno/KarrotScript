package main

import (
	"KarrotScript/src"
	"github.com/davecgh/go-spew/spew"
	"log"
	"os"
)

func main() {
	var filename string

	if len(os.Args) == 2 {
		filename = os.Args[1]
	} else {
		log.Fatal("Please specify a .ksc file")
	}

	feeder := NewLocalFileFeeder()

	ksi := src.NewKarrotScript(feeder)

	err := ksi.Parser.Parse(filename)

	spew.Dump(err)
}
