package main

import (
	"fmt"
	"strconv"
)

const (
	TagA  = "-a"
	TagL  = "-l"
	space = " "
)

type HandleTag func([]ItemStat, int) (string, error)

var SupportedTags = map[string]HandleTag{
	TagA: handleTagA,
	TagL: handleTagL,
}

func handleTagA(items []ItemStat, width int) (string, error) {
	var res string
	return res, nil
}

func handleTagL(items []ItemStat, _ int) (string, error) {
	var res string
	var maxWidthSize int

	for _, v := range items {
		l := len(strconv.Itoa(int(v.FileInfo.Size())))
		if maxWidthSize < l {
			maxWidthSize = l
		}
	}

	for _, v := range items {

		size := fmt.Sprintf("%-*s", maxWidthSize, strconv.Itoa(int(v.FileInfo.Size())))

		res += fmt.Sprint(v.FileInfo.Mode(), space, size, space, v.val, "\n")
	}

	return res, nil
}
