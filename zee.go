package main

import (
	"os"
	"slices"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

var SupportedTag = []string{"-a", "-s"}

type ZeeConfig struct {
	termwidth int
	path, tag string
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
		tag:       SupportedTag[0],
	}

	if len(args) > 1 {
		for _, v := range args[1:] {
			if strings.Contains(v, "-") {
				if slices.Contains(SupportedTag, v) {
					config.tag = v
				}
			} else {
				config.path = v
			}
		}
	}
}
