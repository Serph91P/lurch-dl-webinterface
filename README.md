# LurchDL - a downloader for [gronkh.tv](https://gronkh.tv)

[![en](https://img.shields.io/badge/lang-en-red.svg)](./README.md)
[![de](https://img.shields.io/badge/lang-de-yellow.svg)](./README.de.md)

## Features

- Download [Stream-Episodes](https://gronkh.tv/streams/)
- Specify a start- and stop-timestamp to download only a portion of the video
- Continuable Downloads
- Commandline Interface

## Known Issues

- Downloads are capped to 10 Mbyte/s and buffering is simulated to pre-empt IP blocking due to API ratelimiting
- Start- and stop-timestamps are not very accurate (± 8 seconds)
- Some videoplayers may have problems with the resulting file. To fix this, you can use ffmpeg to rewrite the video into a MKV-File: `ffmpeg -i video.ts -acodec copy -vcodec copy video.mkv`

## Supported Platforms

- Linux

## Download / Installation

Executables will appear under [Releases](https://github.com/ChaoticByte/lurch-dl/releases). Just download one and run it via the terminal.

## Usage

```
lurch-dl --url string       The url to the video
         [-h --help]        Show this help and exit
         [--list-formats]   List available formats and exit
         [--format string]  The desired video format (default: auto)
         [--output string]  The output file. Will be determined automatically
                            if omitted.
         [--start string]   Define a video timestamp to start at, e.g. 12m34s
         [--stop string]    Define a video timestamp to stop at, e.g. 1h23m45s
         [--continue]       Continue the download if possible
         [--overwrite]      Overwrite the output file if it already exists
```

### Examples

Download a video in its best available format:

```
./lurch-dl --url https://gronkh.tv/streams/777

Title: GTV0777, 2023-11-09 - DIESER STREAM IST ILLEGAL UND SOLLTE VERBOTEN WERDEN!! ⭐ ️ 247 auf @GronkhTV ⭐ ️ !comic !archiv !a
Format: 1080p60
Downloaded 0.43% at 10.00 MB/s
...
```

Continue a download:

```
./lurch-dl --url https://gronkh.tv/streams/777 --continue

Title: GTV0777, 2023-11-09 - DIESER STREAM IST ILLEGAL UND SOLLTE VERBOTEN WERDEN!! ⭐ ️ 247 auf @GronkhTV ⭐ ️ !comic !archiv !a
Format: 1080p60
Downloaded 0.68% at 10.00 MB/s
...
```

Specify a start- and stop-timestamp:

```
./lurch-dl --url https://gronkh.tv/streams/777 --start 5h6m41s --stop 5h6m58s
...
```

List all available formats for a video:

```
./lurch-dl --url https://gronkh.tv/streams/777 --list-formats

Available formats:
 - 1080p60
 - 720p
 - 360p
```

Download the video in a specific format:

```
./lurch-dl --url https://gronkh.tv/streams/777 --format 720p

Title: GTV0777, 2023-11-09 - DIESER STREAM IST ILLEGAL UND SOLLTE VERBOTEN WERDEN!! ⭐ ️ 247 auf @GronkhTV ⭐ ️ !comic !archiv !a
Format: 720p
Downloaded 0.32% at 10.00 MB/s
...
```

Specify a filename:

```
./lurch-dl --url https://gronkh.tv/streams/777 --output Stream777.ts
...
```
