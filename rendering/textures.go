package rendering

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

const (
	CROSS_TEXTURE = "./resources/textures/cross.csv"
	ASD = "./resources/textures/asd.csv"
	a = "./resources/textures/cross.csv"
)

type Texture [][]uint8

func LoadTexture(path string) Texture {
	file, err := os.Open(path)
	if err != nil {
		panic(path + "not found")
	}
	defer file.Close()

	textureAsStrings, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}

	rows := len(textureAsStrings)
	cols := len(textureAsStrings[0])
	if rows != cols {
		panic(path + ": CSV should have same size row and column")
	}

	res := make([][]uint8, rows)
	for i, row := range textureAsStrings {
		r := make([]uint8, rows)
		for j, e := range row {
			n, err := strconv.Atoi(strings.TrimSpace(e))
			if err != nil {
				panic(err)
			}
			r[j] = uint8(n)
		}
		res[i] = r
	}

	return res
}
