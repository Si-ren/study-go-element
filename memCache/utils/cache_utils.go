package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ParseStringToInt(string) int {
	return 0
	// return strconv.Itoa(int(int
}

func ParseIntToString(int) string {

	return ""
}

func ParseSize(size string) (uint64, error) {
	// 分割数字和单位
	reg := regexp.MustCompile(`(\d+)([A-Za-z]+)`)
	matches := reg.FindStringSubmatch(size)
	numStr := matches[1]
	unit := matches[2]
	fmt.Printf("number is %s, unit is %s\n", numStr, unit)
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, err
	}
	// 把单位都转化为大写
	unit = strings.ToUpper(unit)
	switch unit {
	case "KB":
		return uint64(num * 1024), nil
	case "MB":
		return uint64(num * 1024 * 1024), nil
	case "GB":
		return uint64(num * 1024 * 1024 * 1024), nil
	default:
		fmt.Println("unknown unit")
		return 0, errors.New("unknown unit")
	}
}
