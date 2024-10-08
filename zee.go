package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

const (
	// tag
	TagA = "-a"
	TagL = "-l"

	// space
	spc  = " "
	spc2 = "  "

	// content-type
	BIN = "application/octet-stream"
)

type ZeeConfig struct {
	termwidth int
	path      string
	handleTag HandleTag
}

type ItemStat struct {
	fs.FileInfo
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
					throwError(fmt.Errorf("undefined tag %v", v))
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

		strItem := createStringItem(info)
		ItemStats = append(ItemStats, ItemStat{
			val:      strItem,
			length:   len(strItem),
			FileInfo: info,
		})
	}

	return config.handleTag(&ItemStats, config.termwidth)
}

func createStringItem(info fs.FileInfo) string {
	var icon string

	if info.IsDir() {
		icon = Type["dir"]
	} else if v, ok := Type[strings.ToLower(strings.Split(info.Name(), ".")[len(strings.Split(info.Name(), "."))-1])]; ok {
		icon = v
	} else if info.Mode().IsRegular() && info.Mode()&0o111 != 0 {
		icon = Type["bin"]
	} else {
		icon = Type["text"]
	}

	return fmt.Sprint(icon, spc, info.Name())
}
