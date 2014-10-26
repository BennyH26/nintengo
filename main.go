package main

import (
	"fmt"
	"os"

	"flag"

	"github.com/nwidger/nintengo/http"
	"github.com/nwidger/nintengo/nes"
)

func main() {
	options := &nes.Options{}

	flag.BoolVar(&options.CPUDecode, "cpu-decode", false, "decode CPU instructions")
	flag.StringVar(&options.Recorder, "recorder", "", "recorder to use: none | jpeg | gif")
	flag.StringVar(&options.AudioRecorder, "audio-recorder", "", "recorder to use: none | wav")
	flag.StringVar(&options.CPUProfile, "cpu-profile", "", "write CPU profile to file")
	flag.StringVar(&options.MemProfile, "mem-profile", "", "write memory profile to file")
	flag.StringVar(&options.HTTPAddress, "http", "", "HTTP service address (e.g., ':6060')")
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Fprintf(os.Stderr, "usage: <rom-file>\n")
		return
	}

	nes, err := nes.NewNES(flag.Arg(0), options)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	if options.HTTPAddress != "" {
		neserv := http.NewNEServer(nes, options.HTTPAddress)
		go neserv.Run()
	}

	err = nes.Run()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}
