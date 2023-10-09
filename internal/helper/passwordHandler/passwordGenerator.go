package passwordHandler

import (
	"crud-go-project/internal/constant"
	"math/rand"
	"strconv"
)

func RangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}

func FileStructureGenerator(low, hi int, userNames []string) []string {
	var data = make([]string, 0)
	for i := 0; i < constant.NUMBER_OF_USERS; i++ {
		data = append(data, userNames[RangeIn(1, constant.NUMBER_OF_USERNAMES)]+strconv.Itoa(i)+" "+strconv.Itoa(RangeIn(low, hi)))
	}
	return data
}
