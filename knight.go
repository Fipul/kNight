package knight

import (
	"errors"
	"strconv"
)

func AvailableMoves(point string) ([8]string, error) {
	var result [8]string
	var err error
	x := 0
	decodedPoint, err := DecodePoint(point)

	if err != nil {
		return result, errors.New("can't decode point: " + err.Error())
	}
	if decodedPoint[0] < 1 || decodedPoint[0] > 8 || decodedPoint[1] < 1 || decodedPoint[1] > 8 {
		return result, err
	}

	if decodedPoint[1]-1 > 0 && decodedPoint[0]-2 > 0 {
		result[x], err = EncodePoint([2]int{decodedPoint[0] - 2, decodedPoint[1] - 1})
		if err != nil {
			return result, errors.New("can't encode")
		}
		x++
	}
	if decodedPoint[1]+1 < 9 && decodedPoint[0]-2 > 0 {
		result[x], err = EncodePoint([2]int{decodedPoint[0] - 2, decodedPoint[1] + 1})
		if err != nil {
			return result, errors.New("can't encode")
		}
		x++
	}
	if decodedPoint[0]-1 > 0 && decodedPoint[1]-2 > 0 {
		result[x], err = EncodePoint([2]int{decodedPoint[0] - 1, decodedPoint[1] - 2})
		if err != nil {
			return result, errors.New("can't encode")
		}
		x++
	}
	if decodedPoint[0]-1 > 0 && decodedPoint[1]+2 < 9 {
		result[x], err = EncodePoint([2]int{decodedPoint[0] - 1, decodedPoint[1] + 2})
		if err != nil {
			return result, errors.New("can't encode")
		}
		x++
	}
	if decodedPoint[0]+1 < 9 && decodedPoint[1]-2 > 0 {
		result[x], err = EncodePoint([2]int{decodedPoint[0] + 1, decodedPoint[1] - 2})
		if err != nil {
			return result, errors.New("can't encode")
		}
		x++
	}
	if decodedPoint[0]+1 < 9 && decodedPoint[1]+2 < 9 {
		result[x], err = EncodePoint([2]int{decodedPoint[0] + 1, decodedPoint[1] + 2})
		if err != nil {
			return result, errors.New("can't encode")
		}
		x++
	}
	if decodedPoint[1]-1 > 0 && decodedPoint[0]+2 < 9 {
		result[x], err = EncodePoint([2]int{decodedPoint[0] + 2, decodedPoint[1] - 1})
		if err != nil {
			return result, errors.New("can't encode")
		}
		x++
	}
	if decodedPoint[1]+1 < 9 && decodedPoint[0]+2 < 9 {
		result[x], err = EncodePoint([2]int{decodedPoint[0] + 2, decodedPoint[1] + 1})
		if err != nil {
			return result, errors.New("can't encode")
		}
		x++
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
	default:
		return result, errors.New("wrong point[0]")
	}

	result[1], err = strconv.Atoi(point[1:])
	if err != nil {
		return result, errors.New("wrong point[1]: can't Atoi: " + err.Error())
	}
	return result, nil
}

func EncodePoint(point [2]int) (string, error) {
	var result string
	if point[0] < 1 || point[0] > 8 || point[1] < 1 || point[1] > 8 {
		return result, errors.New("wrong point")
	}
	switch point[0] {
	case 1:
		result = "a"
	case 2:
		result = "b"
	case 3:
		result = "c"
	case 4:
		result = "d"
	case 5:
		result = "e"
	case 6:
		result = "f"
	case 7:
		result = "g"
	case 8:
		result = "h"
	}
	result += strconv.Itoa(point[1])
	return result, nil
}
