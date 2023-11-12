// Copyright (c) 2023 Julian Müller (ChaoticByte)

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"
)

const ApiBaseurlStreamEpisodeInfo = "https://api.gronkh.tv/v1/video/info?episode=%s"
const ApiBaseurlStreamEpisodePlInfo = "https://api.gronkh.tv/v1/video/playlist?episode=%s"

var ApiHeadersBase = http.Header{
	"User-Agent": {"Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/119.0"},
	"Accept-Language": {"de,en-US;q=0.7,en;q=0.3"},
	//"Accept-Encoding": {"gzip"},
	"Origin": {"https://gronkh.tv"},
    "Referer": {"https://gronkh.tv/"},
	"Connection": {"keep-alive"},
	"Sec-Fetch-Dest": {"empty"},
	"Sec-Fetch-Mode": {"cors"},
	"Sec-Fetch-Site": {"same-site"},
	"Pragma": {"no-cache"},
	"Cache-Control": {"no-cache"},
	"TE": {"trailers"},
}

var ApiHeadersMetaAdditional = http.Header{
	"Accept": {"application/json, text/plain, */*"},
}

var ApiHeadersVideoAdditional = http.Header{
	"Accept": {"*/*"},
}

type VideoChunkList struct {
	BaseUrl string
	Chunks []string
	ChunkDuration float64
}

func (cl *VideoChunkList) Cut(from time.Duration, to time.Duration) VideoChunkList {
	var newChunks []string
	var firstChunk = 0
	if from != -1 {
		firstChunk = int(from.Seconds() / cl.ChunkDuration)
	}
	if to != -1 {
		lastChunk := min(int(to.Seconds() / cl.ChunkDuration)+1, len(cl.Chunks))
		newChunks = cl.Chunks[firstChunk:lastChunk]
	} else {
		newChunks = cl.Chunks[firstChunk:]
	}
	return VideoChunkList{
		BaseUrl: cl.BaseUrl,
		Chunks: newChunks,
		ChunkDuration: cl.ChunkDuration,
	}
}

func GetVideoChunkList(video VideoFormat) (VideoChunkList, error) {
	baseUrl := video.Url[:strings.LastIndex(video.Url, "/")]
	data, err := httpGet(video.Url, []http.Header{ApiHeadersBase, ApiHeadersMetaAdditional}, nil, time.Second * 10)
	if err != nil {
		return VideoChunkList{}, err
	}
	chunklist, err := parseChunkListFromM3u8(string(data), baseUrl)
	return chunklist, err
}

type VideoFormat struct {
	Name string
	Url string
}

type FormatNotFoundError struct {
	FormatName string
}

func (err *FormatNotFoundError) Error() string {
	return "Format " + err.FormatName + " is not available."
}

type Chapter struct {
	Title string `json:"title"`
	Offset time.Duration `json:"offset"`
}

type StreamEpisodeMeta struct {
	Episode string
	Formats []VideoFormat
	Title string `json:"title"`
	ProposedFilename string
	PlaylistUrl string `json:"playlist_url"`
	Chapters []Chapter `json:"chapters"`
}

func (meta StreamEpisodeMeta) GetFormat(formatName string) (VideoFormat, error) {
	if formatName == "auto" {
		// at the moment, the best format is always the first
		return meta.Formats[0], nil
	} else {
		var format VideoFormat
		formatFound := false
		for _, f := range meta.Formats {
			if f.Name == formatName {
				format = f
				formatFound = true
			}
		}
		if formatFound {
			return format, nil
		} else {
			return format, &FormatNotFoundError{FormatName: formatName}
		}
	}
}

func GetStreamEpisodeMeta(episode string, chapterIdx int) (StreamEpisodeMeta, error) {
	meta := StreamEpisodeMeta{}
	meta.Episode = episode
	info_data, err := httpGet(
		fmt.Sprintf(ApiBaseurlStreamEpisodeInfo, episode),
		[]http.Header{ApiHeadersBase, ApiHeadersMetaAdditional},
		nil, time.Second * 10,
	)
	if err != nil {
		return meta, err
	}
	// Title
	json.Unmarshal(info_data, &meta)
	meta.Title = strings.ToValidUTF8(meta.Title, "")
	// sanitized proposedFilename
	if chapterIdx >= 0 && chapterIdx < len(meta.Chapters) {
		meta.ProposedFilename = fmt.Sprintf("GTV%04s - %v. %s.ts", episode, chapterIdx+1, meta.Chapters[chapterIdx].Title)
	} else {
		meta.ProposedFilename = sanitizeUnicodeFilename(meta.Title) + ".ts"
	}
	// Sort Chapters & correct offset
	for i := range meta.Chapters {
		meta.Chapters[i].Offset = meta.Chapters[i].Offset * time.Second
	}
	sort.Slice(meta.Chapters, func(i int, j int) bool {
		return meta.Chapters[i].Offset < meta.Chapters[j].Offset
	})
	// Formats
	playlist_url_data, err := httpGet(
		fmt.Sprintf(ApiBaseurlStreamEpisodePlInfo, episode),
		[]http.Header{ApiHeadersBase, ApiHeadersMetaAdditional},
		nil, time.Second * 10,
	)
	if err != nil {
		return meta, err
	}
	json.Unmarshal(playlist_url_data, &meta)
	playlist_data, err := httpGet(
		meta.PlaylistUrl,
		[]http.Header{ApiHeadersBase, ApiHeadersMetaAdditional},
		nil, time.Second * 10,
	)
	meta.Formats = parseAvailFormatsFromM3u8(string(playlist_data))
	return meta, err
}
