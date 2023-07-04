package utils

import (
	"errors"
	"strconv"
)

var ErrParam = errors.New("error parsing param")

func Uint64(param string) uint64 {
	val, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return 0
	}
	return uint64(val)
}

func Int64(param string) int64 {
	val, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return 0
	}
	return val
}

func Uint(param string) uint {
	val, err := strconv.ParseUint(param, 10, 32)
	if err != nil {
		return 0
	}
	return uint(val)
}
