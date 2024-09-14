package main

import (
	"fmt"
	"math"
	"strconv"
)

type HandleTag func(*[]ItemStat, int) (string, error)

var SupportedTags = map[string]HandleTag{
	TagA: handleTagA,
	TagL: handleTagL,
}

// function for handle -a tag, or without tag
// for example =
// ```
// > zee -a
// or
// > zee
// ```
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

// function for handle -l tag
// for example =
// ```
// > zee -l
// ```
func handleTagL(items *[]ItemStat, _ int) (string, error) {
	res := fmt.Sprint("total: ", len(*items), "\n\n")

	widthSize, unitDataSize := createMaxWidthSizeAndUnitDataWidth(items)

	for _, v := range *items {

		size, unitData := createSizeString(int(v.Size()))

		sizeString := fmt.Sprintf("%*s", widthSize, size)
		unitDataString := fmt.Sprintf("%-*s", unitDataSize, unitData)

		res += fmt.Sprint(v.FileInfo.Mode(), spc2, sizeString, spc, unitDataString, spc2, v.val, "\n")
	}

	return res, nil
}

func createMaxWidthSizeAndUnitDataWidth(items *[]ItemStat) (int, int) {
	var maxWidthSize, unitDataWidth int
	for _, v := range *items {

		s := strconv.Itoa(int(v.Size()))
		l := len(s)

		if l >= 7 || l >= 4 {
			unitDataWidth = 2
			if l >= 7 {
				r := len(strconv.Itoa(int(v.Size() / 1000000)))
				if maxWidthSize < r {
					maxWidthSize = r
				}
			} else if l >= 4 {
				r := len(strconv.Itoa(int(v.Size() / 1000)))
				if maxWidthSize < r {
					maxWidthSize = r
				}
			}
		} else {
			if unitDataWidth < l {
				unitDataWidth = 1
			}
			if maxWidthSize < l {
				maxWidthSize = l
			}
		}

	}

	return maxWidthSize, unitDataWidth
}

func createSizeString(size int) (string, string) {
	var s string
	var u string
	l := len(strconv.Itoa(size))

	if l >= 7 {
		s = strconv.Itoa(size / 1000000)
		u = "MB"
	} else if l >= 4 {
		s = strconv.Itoa(size / 1000)
		u = "KB"
	} else {
		s = strconv.Itoa(size)
		u = "B"
	}

	return s, u
}
