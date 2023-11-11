# lurch-dl - ein Downloader für [gronkh.tv](https://gronkh.tv)

[![en](https://img.shields.io/badge/lang-en-red.svg)](./README.md)
[![de](https://img.shields.io/badge/lang-de-yellow.svg)](./README.de.md)
![GitHub top language](https://img.shields.io/github/languages/top/ChaoticByte/lurch-dl)
![GitHub all releases](https://img.shields.io/github/downloads/ChaoticByte/lurch-dl/total)

## Features

- Download von [Stream-Episoden](https://gronkh.tv/streams/)
- Mit Start- und Stop-Timestamps kann auch nur ein Teil des Videos heruntergeladen werden
- Downloads können fortgesetzt werden
- Bedienung über das Terminal

## Bekannte Fehler

- Downloads können mit maximal 10 Mbyte/s heruntergeladen werden, zudem wird Buffering simuliert. Das ist notwendig um IP-Bans durch API-Ratelimits zu verhindern.
- Bei Start- und Stop-Timestamps kann es zu Abweichungen von ± 8 Sekunden kommen
- Manche Video-Player könnten mit der heruntergeladenen Videodatei Probleme haben. Dies kann mit FFMPEG behoben werden, indem das Video in eine MKV-Datei umgeschrieben wird: `ffmpeg -i video.ts -acodec copy -vcodec copy video.mkv`

## Unterstützte Platformen

- Linux

## Download / Installation

Executables sind unter [Releases](https://github.com/ChaoticByte/lurch-dl/releases) zu finden. Hier kann einfach eine Executable heruntergeladen, und über das Terminal ausgeführt werden.

## Verwendung

```
lurch-dl --url string       Die URL zum Video
         [-h --help]        Zeigt diesen Hilfetext an
         [--list-formats]   Listet alle verfügbaren Formate
         [--format string]  Das gewünschte Videoformat (default: auto)
         [--output string]  Der Dateiname. Wird automatisch ermittelt
                            wenn dieser Parameter weggelassen wird.
         [--start string]   Start-Timestamp, z.B. 12m34s
         [--stop string]    Stop-Timestamp, z.B. 1h23m45s
         [--continue]       Fortsetzen des Downloads, wenn möglich
         [--overwrite]      Überschreiben der Datei, wenn diese bereits existiert
```

### Beispiele

Download eines Videos im besten verfügbaren Format:

```
./lurch-dl --url https://gronkh.tv/streams/777

Title: GTV0777, 2023-11-09 - DIESER STREAM IST ILLEGAL UND SOLLTE VERBOTEN WERDEN!! ⭐ ️ 247 auf @GronkhTV ⭐ ️ !comic !archiv !a
Format: 1080p60
Downloaded 0.43% at 10.00 MB/s
...
```

Fortsetzen eines Downloads:

```
./lurch-dl --url https://gronkh.tv/streams/777 --continue

Title: GTV0777, 2023-11-09 - DIESER STREAM IST ILLEGAL UND SOLLTE VERBOTEN WERDEN!! ⭐ ️ 247 auf @GronkhTV ⭐ ️ !comic !archiv !a
Format: 1080p60
Downloaded 0.68% at 10.00 MB/s
...
```

Angeben eines Start- und Stop-Timestamps:

```
./lurch-dl --url https://gronkh.tv/streams/777 --start 5h6m41s --stop 5h6m58s
...
```

Auflisten aller verfügbaren Formate:

```
./lurch-dl --url https://gronkh.tv/streams/777 --list-formats

Available formats:
 - 1080p60
 - 720p
 - 360p
```

Download im angegebenen Format:

```
./lurch-dl --url https://gronkh.tv/streams/777 --format 720p

Title: GTV0777, 2023-11-09 - DIESER STREAM IST ILLEGAL UND SOLLTE VERBOTEN WERDEN!! ⭐ ️ 247 auf @GronkhTV ⭐ ️ !comic !archiv !a
Format: 720p
Downloaded 0.32% at 10.00 MB/s
...
```

Angeben eines Dateinamens:

```
./lurch-dl --url https://gronkh.tv/streams/777 --output Stream777.ts
...
```
