package main

import "io/fs"

const (
	TagA = "-a"
	TagS = "-s"
)

type HandleTag func([]ItemStat, int)(string,error)


var SupportedTags = map[string]HandleTag{
	TagA: handleTagA,
	TagS: handleTagS,
}

func handleTagA(items []ItemStat, width int) (string , error) {
	return res string

  for i , v := range items {

  }


}

func handleTagS(items []ItemStat, width int) error {
	return nil
}
