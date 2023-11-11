// Copyright (c) 2023 Julian Müller (ChaoticByte)

package main

import (
	"strings"
	"unicode"
)

var FnInvalidRunes = []rune("/<>:\"\\|?*")

func sanitizeUnicodeFilename(filename string) string {
	filename = strings.Trim(strings.ToValidUTF8(filename, ""), " \u00A0\t\n\r.")
	var filenameBuilder strings.Builder
	for _, r := range filename {
		isInvalid := !unicode.IsPrint(r)
		if isInvalid { continue }
		for _, c := range FnInvalidRunes {
			if r == c {
				isInvalid = true
				break
			}
		}
		if !isInvalid {
			filenameBuilder.WriteRune(r)
		}
	}
	return filenameBuilder.String()
}
