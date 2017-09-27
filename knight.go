package knight

import (
	"errors"
	"strconv"
)

func AvaliableMoves(point string) ([]string, error) {
	var result []string
	decodedpoint, err := DecodePoint(point)
	if err != nil {
		return nil, errors.New("can't decode point: " + err.Error())
	}

	for i := 0; i < 8; i++ {
		encodedpoint, err := EncodePoint(decodedpoint)
		if err != nil {
			return nil, errors.New("Can't encode point" + err.Error())
		}
		result[i] = encodedpoint
	}
	return result, nil
}

func DecodePoint(point string) ([2]int, error) {
	var result [2]int
	var err error
	if len(point) != 2 {
		return result, errors.New("len != 2")
	}
	switch point[0:1] {
	case "a":
		result[0] = 1
	case "b":
		result[0] = 2
	case "c":
		result[0] = 3
	case "d":
		result[0] = 4
	case "e":
		result[0] = 5
	case "f":
		result[0] = 6
	case "g":
		result[0] = 7
	case "h":
		result[0] = 8
	}
	result[1], err = strconv.Atoi(point[1:])
	if err != nil {
		return result, errors.New("can't Atoi" + err.Error())
	}
	return result, nil
}

func EncodePoint(point [2]int) (string, error) {
	var result string
	slice := result[0:1]
	switch point[0] {
	case 1:
		slice = "a"
	case 2:
		slice = "b"
	case 3:
		slice = "c"
	case 4:
		slice = "d"
	case 5:
		slice = "e"
	case 6:
		slice = "f"
	case 7:
		slice = "g"
	case 8:
		slice = "h"
	}
	result = slice
	result += strconv.Itoa(point[1])
	return result, nil
}
