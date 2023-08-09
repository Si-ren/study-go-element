package utils_test

import (
	"fmt"
	"memCache/utils"
	"testing"
)

func TestParseSize(t *testing.T) {
	// num, err := utils.ParseSize("100mb")
	// if err != nil {
	// 	t.Error(err)
	// }
	// assert.Equal(t, num, uint64(100*1024*1024), "error number parsing size")
	_, err := utils.ParseSize("100rb")

	fmt.Println(err.Error())
	t.Log(err)

}
