// Copyright (c) 2023 Julian Müller (ChaoticByte)

package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type UserInterface interface {
	Run()
	AvailableFormats(formats []VideoFormat)
	Progress(percentage float32, rate float64, delaying bool, waiting bool, retries int)
	InfoMessage(msg string)
	Aborted()
	Help()
}

type Cli struct {}

func (cli *Cli) Run() {
	// cli arguments
	var help bool
	var listFormats bool
	var url string
	var formatName string
	var outputFile string
	var timestampStart string
	var timestampStop string
	var overwrite bool
	var continueDl bool
	// var outputFile string
	flag.BoolVar(&help, "h", false, "")
	flag.BoolVar(&help, "help", false, "")
	flag.BoolVar(&listFormats, "list-formats", false, "")
	flag.StringVar(&url, "url", "", "")
	flag.StringVar(&formatName, "format", "auto", "")
	flag.StringVar(&outputFile, "output", "", "")
	flag.StringVar(&timestampStart, "start", "", "")
	flag.StringVar(&timestampStop, "stop", "", "")
	flag.BoolVar(&overwrite, "overwrite", false, "")
	flag.BoolVar(&continueDl, "continue", false, "")
	flag.Parse()
	var startDuration time.Duration
	var stopDuration time.Duration
	var err error
	if timestampStart == "" {
		startDuration = -1
	} else {
		startDuration, err = time.ParseDuration(timestampStart)
		if err != nil {
			fmt.Printf("Couldn't parse start timestamp '%v'.\n%v\n", timestampStart, err)
			os.Exit(1)
		}
	}
	if timestampStop == "" {
		stopDuration = -1
	} else {
		stopDuration, err = time.ParseDuration(timestampStop)
		if err != nil {
			fmt.Printf("Couldn't parse stop timestamp '%v'.\n%v\n", timestampStop, err)
			os.Exit(1)
		}
	}
	// run actions
	if help {
		cli.Help()
		os.Exit(0)
	} else if url == "" {
		cli.Help()
		os.Exit(1)
	}
	video, err := ParseGtvVideoUrl(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if video.Category != "streams" {
		fmt.Println("Video category '" + video.Category + "' not supported.")
		os.Exit(1)
	}
	meta, err := GetStreamEpisodeMeta(video.Id)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if listFormats {
		cli.AvailableFormats(meta.Formats)
		os.Exit(0)
	}
	format, err := meta.GetFormat(formatName)
	if err != nil {
		fmt.Println(err)
		cli.AvailableFormats(meta.Formats)
		os.Exit(1)
	}
	fmt.Printf("Title: %v\nFormat: %v\n", meta.Title, format.Name)
	defer fmt.Print("\n")
	if err = DownloadStreamEpisode(meta, format, startDuration, stopDuration, outputFile, overwrite, continueDl, cli); err != nil {
		fmt.Print("\n")
		fmt.Println(err)
		os.Exit(1)
	}
}

func (cli *Cli) AvailableFormats(formats []VideoFormat) {
	fmt.Println("Available formats:")
	for _, f := range formats {
		fmt.Println(" - " + f.Name)
	}
}

func (cli *Cli) Progress(percentage float32, rate float64, delaying bool, waiting bool, retries int) {
	if retries > 0 {
		fmt.Printf("\nDownloaded %.2f%% at %.2f MB/s (retry %v) ...      \r", percentage * 100.0, rate / 1000000.0, retries)
	} else if waiting {
		fmt.Printf("Downloaded %.2f%% at %.2f MB/s ...                 \r", percentage * 100.0, rate / 1000000.0)
	} else if delaying {
		fmt.Printf("Downloaded %.2f%% at %.2f MB/s delaying ...        \r", percentage * 100.0, rate / 1000000.0)
	} else {
		fmt.Printf("Downloaded %.2f%% at %.2f MB/s                     \r", percentage * 100.0, rate / 1000000.0)
	}
}

func (cli *Cli) InfoMessage(msg string) {
	fmt.Println(msg)
}

func (cli *Cli) Aborted() {
	fmt.Print("\nAborted.                                                ")
}

func (cli *Cli) Help() {
	fmt.Println(`lurch-dl --url string       The url to the video
         [-h --help]        Show this help and exit
         [--list-formats]   List available formats and exit
         [--format string]  The desired video format (default: auto)
         [--output string]  The output file. Will be determined automatically
                            if omitted.
         [--start string]   Define a video timestamp to start at, e.g. 12m34s
         [--stop string]    Define a video timestamp to stop at, e.g. 1h23m45s
         [--continue]       Continue the download if possible
         [--overwrite]      Overwrite the output file if it already exists

Version: ` + Version)
}
