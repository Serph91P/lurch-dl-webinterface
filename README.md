# lurch-dl - a downloader for [gronkh.tv](https://gronkh.tv)

[![en](https://img.shields.io/badge/lang-en-red.svg)](./README.md)
[![de](https://img.shields.io/badge/lang-de-yellow.svg)](./README.de.md)
![GitHub top language](https://img.shields.io/github/languages/top/ChaoticByte/lurch-dl)
![GitHub release (with filter)](https://img.shields.io/github/v/release/ChaoticByte/lurch-dl)


## Features

- Download [Stream-Episodes](https://gronkh.tv/streams/)
- Specify a start- and stop-timestamp to download only a portion of the video
- Download a specific chapter
- Continuable Downloads
- Commandline Interface

## Known Issues

- Downloads are capped to 10 Mbyte/s and buffering is simulated to pre-empt IP blocking due to API ratelimiting
- Start- and stop-timestamps are not very accurate (± 8 seconds)
- Some videoplayers may have problems with the resulting file. To fix this, you can use ffmpeg to rewrite the video into a MKV-File: `ffmpeg -i video.ts -acodec copy -vcodec copy video.mkv`
- Emojis and other Unicode characters don't get displayed properly in a Powershell Console

## Supported Platforms

- Linux (i386, amd64*, arm, arm64)
- Windows (32bit, 64bit*, arm64)

\* tested

## Download / Installation

Executables will appear under [Releases](https://github.com/ChaoticByte/lurch-dl/releases). Just download one and run it via the terminal.

## Usage

This is a commandline application. This means that you can only use it in a terminal (Powershell on Windows).

This are the commandline arguments:

```
lurch-dl --url string       The url to the video
         [-h --help]        Show this help and exit
         [--list-chapters]  List chapters and exit
         [--list-formats]   List available formats and exit
         [--chapter int]    The chapter you want to download
                            The calculated start and stop timestamps can be
                            overwritten by --start and --stop
                            default: -1 (disabled)
         [--format string]  The desired video format
                            default: auto
         [--output string]  The output file. Will be determined automatically
                            if omitted.
         [--start string]   Define a video timestamp to start at, e.g. 12m34s
         [--stop string]    Define a video timestamp to stop at, e.g. 1h23m45s
         [--continue]       Continue the download if possible
         [--overwrite]      Overwrite the output file if it already exists
```

### Windows

1. Open the folder that contains lurch-dl.exe
2. `Shift`+`Right-Click` in the folder -> click on `Open Powershell window here`.
3. Run the application as shown above and below

### Examples

Download a video in its best available format (Windows):

```
.\lurch-dl.exe --url https://gronkh.tv/streams/777

Title: GTV0777, 2023-11-09 - DIESER STREAM IST ILLEGAL UND SOLLTE VERBOTEN WERDEN!! ⭐ ️ 247 auf @GronkhTV ⭐ ️ !comic !archiv !a
Format: 1080p60
Downloaded 0.43% at 10.00 MB/s
...
```

Continue a download (Windows):

```
.\lurch-dl.exe --url https://gronkh.tv/streams/777 --continue

Title: GTV0777, 2023-11-09 - DIESER STREAM IST ILLEGAL UND SOLLTE VERBOTEN WERDEN!! ⭐ ️ 247 auf @GronkhTV ⭐ ️ !comic !archiv !a
Format: 1080p60
Downloaded 0.68% at 10.00 MB/s
...
```

List all chapters (Windows):

```
.\lurch-dl.exe --url https://gronkh.tv/streams/777 --list-chapters

GTV0777, 2023-11-09 - DIESER STREAM IST ILLEGAL UND SOLLTE VERBOTEN WERDEN!! ⭐ ️ 247 auf @GronkhTV ⭐ ️ !comic !archiv !a

Chapters:
  1         0s	Just Chatting
  2    2h53m7s	Alan Wake II
  3    9h35m0s	Just Chatting
```

Download a specific chapter (Windows):

```
.\lurch-dl.exe --url https://gronkh.tv/streams/777 --chapter 2

GTV0777, 2023-11-09 - DIESER STREAM IST ILLEGAL UND SOLLTE VERBOTEN WERDEN!! ⭐ ️ 247 auf @GronkhTV ⭐ ️ !comic !archiv !a
Format: 1080p60
Chapter: 2. Alan Wake II

Downloaded 3.22% at 10.00 MB/s
...
```

Specify a start- and stop-timestamp (Linux):

```
./lurch-dl --url https://gronkh.tv/streams/777 --start 5h6m41s --stop 5h6m58s
...
```

List all available formats for a video (Linux):

```
./lurch-dl --url https://gronkh.tv/streams/777 --list-formats

Available formats:
 - 1080p60
 - 720p
 - 360p
```

Download the video in a specific format (Linux):

```
./lurch-dl --url https://gronkh.tv/streams/777 --format 720p

Title: GTV0777, 2023-11-09 - DIESER STREAM IST ILLEGAL UND SOLLTE VERBOTEN WERDEN!! ⭐ ️ 247 auf @GronkhTV ⭐ ️ !comic !archiv !a
Format: 720p
Downloaded 0.32% at 10.00 MB/s
...
```

Specify a filename (Windows):

```
.\lurch-dl.exe --url https://gronkh.tv/streams/777 --output Stream777.ts
...
```
