// Copyright (c) 2023 Julian Müller (ChaoticByte)

package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var availFormatsRegex = regexp.MustCompile(`NAME="(.+)"`)

func parseAvailFormatsFromM3u8(m3u8 string) []VideoFormat {
	foundFormats := []VideoFormat{}
	m3u8 = strings.ReplaceAll(m3u8, "\r", "")
	parts := strings.Split(m3u8, "#EXT-X-STREAM-INF")
	for _, p := range parts {
		p := strings.Trim(p, " \n")
		if strings.HasPrefix(p, ":") && strings.Contains(p, "RESOLUTION=") && strings.Contains(p, "FRAMERATE=") && strings.Contains(p, "NAME=") {
			format := VideoFormat{}
			plItem := strings.Split(p, "\n")
			if len(plItem) < 2 {
				continue
			}
			formatName := availFormatsRegex.FindStringSubmatch(plItem[0])
			if formatName == nil {
				continue // didn't find format
			}
			format.Name = formatName[1]
			format.Url = plItem[1]
			foundFormats = append(foundFormats, format)
		}
	}
	return foundFormats
}

var targetDurationRegex = regexp.MustCompile(`#EXT-X-TARGETDURATION:(.+)`)

func parseChunkListFromM3u8(m3u8 string, baseurl string) (VideoChunkList, error) {
	chunklist := VideoChunkList{BaseUrl: baseurl}
	m3u8 = strings.ReplaceAll(m3u8, "\r", "")
	parts := strings.Split(m3u8, "#EXTINF")
	for _, p := range parts {
		if strings.HasPrefix(p, "#EXTM3U") {
			lines := strings.Split(p, "\n")
			for _, l := range lines {
				if strings.HasPrefix(l, "#EXT-X-TARGETDURATION") {
					targetDuration := targetDurationRegex.FindStringSubmatch(l)
					if targetDuration == nil {
						continue
					}
					chunkDuration, err := strconv.ParseFloat(targetDuration[1], 64)
					if err != nil {
						return chunklist, fmt.Errorf("could not convert %v to float", targetDuration[1])
					}
					chunklist.ChunkDuration = chunkDuration
				}
			}
		} else if strings.HasPrefix(p, ":") {
			lines := strings.Split(p, "\n")
			if len(lines) > 1 {
				chunklist.Chunks = append(chunklist.Chunks, lines[1])
			}
		}
	}
	return chunklist, nil
}
