package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		ret[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			ret[i][j] = uint8((i + j) / 2)
		}
	}
	return ret
}

func slices() {
	pic.Show(Pic)
}

// 実行前に最新版のpicを取得しておく
// go get golang.org/x/tour/pic@latest
// go mod tidy
