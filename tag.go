package main

import "io/fs"

const (
	TagA = "-a"
	TagS = "-s"
)

var SupportedTags = map[string]func([]fs.DirEntry) error{
	TagA: handleTagA,
	TagS: handleTagS,
}

func handleTagA([]fs.DirEntry) error {
	return nil
}

func handleTagS([]fs.DirEntry) error {
	return nil
}
