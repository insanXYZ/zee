package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	TagA = "-a"
	TagL = "-l"
	spc  = " "
	spc2 = "  "
)

type HandleTag func(*[]ItemStat, int) (string, error)

var SupportedTags = map[string]HandleTag{
	TagA: handleTagA,
	TagL: handleTagL,
}

func handleTagA(items *[]ItemStat, width int) (string, error) {
	var res string

	rows, layout := createLayout(items, width)

	for i := 0; i < rows; i++ {
		for j, v := range layout {
			fi := j * rows
			li := fi + rows
			res += fmt.Sprintf("%-*s", v, (*items)[fi:li][i].val)
			if j == len(layout)-1 {
				res += "\n"
			}
		}
	}

	return res, nil
}

func createLayout(items *[]ItemStat, width int) (int, []int) {
	var rows int
	var layout []int
	for rows = 1; true; rows++ {
		layout = []int{}

		cond := int(math.Ceil(float64(len(*items)) / float64(rows)))
		for j := 0; j < cond; j++ {
			var max int

			fi := j * rows
			li := fi + rows

			if li > len(*items) {
				li = len(*items)
			}

			for _, v := range (*items)[fi:li] {
				if max < v.length {
					max = v.length
				}
			}

			layout = append(layout, max+len(spc2))
		}

		if sumSlices(layout) < width {
			break
		}
	}

	return rows, layout
}

func sumSlices[T int | float32 | float64](i []T) T {
	var res T

	for _, v := range i {
		res += v
	}

	return res
}

func handleTagL(items *[]ItemStat, _ int) (string, error) {
	var maxWidthSize int

	res := fmt.Sprint("total: ", len(*items), "\n\n")

	for _, v := range *items {
		l := len(strconv.Itoa(int(v.FileInfo.Size())))
		if maxWidthSize < l {
			maxWidthSize = l
		}
	}

	for _, v := range *items {

		size := fmt.Sprintf("%-*s", maxWidthSize, strconv.Itoa(int(v.FileInfo.Size())))

		res += fmt.Sprint(v.FileInfo.Mode(), spc2, size, spc2, v.val, "\n")
	}

	return res, nil
}
