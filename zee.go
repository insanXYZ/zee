package main

import (
	"errors"
	"fmt"
	"fs"
	"io/fs"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

type ZeeConfig struct {
	termwidth int
	path      string
	handleTag HandleTag
}

type ItemStat struct {
	val    string
	length int
}

var args = os.Args

func main() {
	w, _, err := terminal.GetSize(0)
	if err != nil {
		panic(err.Error())
	}

	config := &ZeeConfig{
		termwidth: w,
		path:      ".",
		handleTag: SupportedTags[TagA],
	}

	if len(args) > 1 {
		for _, v := range args[1:] {
			if strings.Contains(v, "-") {
				if f, ok := SupportedTags[v]; ok {
					config.handleTag = f
				} else {
					throwError(errors.New(fmt.Sprintf("undefined tag %v", v)))
				}
			} else {
				config.path = v
			}
		}
	}

	s, err := readAndParseDir(config)
	if err != nil {
		throwError(err)
	}
	fmt.Println(s)
}

func readAndParseDir(config *ZeeConfig) (string, error) {
	dir, err := os.ReadDir(config.path)
	if err != nil {
		return "", err
	}

	if len(dir) == 0 {
		return "", nil
	}

	var ItemStats []ItemStat

	for _, v := range dir {
		info, err := v.Info()
		if err != nil {
			return "", nil
		}

		s := createStringItem(info)
		ItemStats = append(ItemStats, ItemStat{
			val:    s,
			length: len(s),
		})
	}

	return config.handleTag(ItemStats, config.termwidth)
}

func createStringItem(s fs.FileInfo) string {
	return ""
}
